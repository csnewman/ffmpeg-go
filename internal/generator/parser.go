package main

import (
	"fmt"
	"log"
	"path"
	"slices"
	"strings"

	"github.com/go-clang/bootstrap/clang"
)

const AVLibPath = "/opt/homebrew/Cellar/ffmpeg/6.0_1/include/"

var files = []string{
	"libavcodec/codec_id.h",
	//"libavcodec/avcodec.h",
	"libavformat/avformat.h",
}

func Parse() *Module {
	p := &Parser{
		mod: &Module{
			functions: make(map[string]*Function),
			structs:   make(map[string]*Struct),
			enums:     make(map[string]*Enum),
			callbacks: make(map[string]*Function),
		},
	}

	for _, file := range files {
		filePath := path.Join(AVLibPath, file)

		fmt.Println(filePath)

		p.parseFile(fmt.Sprintf("[%v]", file), filePath)
	}

	return p.mod
}

type Parser struct {
	path string
	mod  *Module
}

func (p *Parser) parseFile(indent string, path string) {
	p.path = path

	idx := clang.NewIndex(0, 1)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit(
		path,
		[]string{
			"-fparse-all-comments",
		},
		nil,
		clang.TranslationUnit_IncludeBriefCommentsInCodeCompletion|clang.TranslationUnit_DetailedPreprocessingRecord,
	)
	defer tu.Dispose()

	diagnostics := tu.Diagnostics()
	for _, d := range diagnostics {
		fmt.Println("PROBLEM:", d.Spelling())
	}

	tu.
		TranslationUnitCursor().
		Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
			p.parseTopLevel(indent, cursor)

			return clang.ChildVisit_Continue
		})
}

func (p *Parser) parseTopLevel(indent string, c clang.Cursor) {
	loc, _, _, _ := c.Location().FileLocation()
	if loc.Name() != p.path {
		return
	}

	switch c.Kind() {
	case clang.Cursor_TypedefDecl:
		p.parseTypedef(indent, c)

	case clang.Cursor_VarDecl:
		log.Println("vardecl", "name", c.Spelling())
		log.Println(" ", c.Type().Spelling())

	case clang.Cursor_FunctionDecl:
		p.parseFunction(indent, c)

	case clang.Cursor_MacroDefinition:
		if c.IsMacroFunctionLike() {
			return
		}

		log.Println("macro def", "name", c.Spelling())

	case clang.Cursor_EnumDecl:
		p.parseEnum(indent, c)

	case clang.Cursor_StructDecl:
		p.parseStruct(indent, c)

	case clang.Cursor_MacroExpansion, clang.Cursor_InclusionDirective, clang.Cursor_UnionDecl:
		// ignore

	default:
		log.Panicln("Unexpected top level", "kind", c.Kind().String())
	}
}

func (p *Parser) parseTypedef(indent string, c clang.Cursor) {
	name := c.Spelling()
	indent = fmt.Sprintf("%v[%v]", indent, name)

	log.Println("typedef", "name", c.Spelling())

	var params []*Param

	c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		log.Println("  --- ", "kind", cursor.Kind().String(), "name", cursor.Spelling())

		if cursor.Kind() == clang.Cursor_ParmDecl {
			name := cursor.Spelling()
			if name == "" {
				log.Fatal("no param name")
			}

			ty := p.parseType(fmt.Sprintf("%v[%v]", indent, name), cursor.Type())

			params = append(params, &Param{
				Name: name,
				Type: ty,
			})
		}

		return clang.ChildVisit_Continue
	})

	log.Println("dk ", c.Definition().Kind())
	log.Println("ds ", c.Definition().Spelling())
	log.Println("s ", c.TypedefDeclUnderlyingType().Spelling())
	log.Println("k ", c.TypedefDeclUnderlyingType().Kind())
	log.Println("cs ", c.TypedefDeclUnderlyingType().CanonicalType().Spelling())
	log.Println("ck ", c.TypedefDeclUnderlyingType().CanonicalType().Kind())

	if len(params) > 0 {
		ut := c.TypedefDeclUnderlyingType()
		pt := ut.PointeeType()

		if pt.NumArgTypes() != int32(len(params)) {
			log.Panicln("arg mismatch")
		}

		result := p.parseType(fmt.Sprintf("%v[%v]", indent, name), pt.ResultType())

		if _, ok := p.mod.callbacks[name]; ok {
			log.Panicln("callback already exists")

			return
		}

		p.mod.callbacks[name] = &Function{
			Name:   name,
			Args:   params,
			Result: result,
		}
		p.mod.callbackOrder = append(p.mod.callbackOrder, name)

		return
	}

	dec := c.TypedefDeclUnderlyingType().Declaration()

	dec.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		log.Println("  Inner ", "kind", cursor.Kind().String(), "name", cursor.Spelling())

		return clang.ChildVisit_Continue
	})

	switch dec.Kind() {
	case clang.Cursor_StructDecl:
		p.parseStruct(indent, dec)

	default:
		log.Panicln("Unknown typedef", "kind", dec.Kind())
	}

}

func (p *Parser) parseFunction(indent string, c clang.Cursor) {
	name := c.Spelling()

	if _, ok := p.mod.functions[name]; ok {
		log.Panicln("Function", name, "already exists")
	}

	log.Println("Parsing function", name)
	indent = fmt.Sprintf("%v[%v]", indent, name)

	result := p.parseType(fmt.Sprintf("%v[return]", indent), c.ResultType())

	var args []*Param

	for i := 0; i < int(c.NumArguments()); i++ {
		arg := c.Argument(uint32(i))

		if arg.Kind() != clang.Cursor_ParmDecl {
			log.Panicln(indent, "Argument not of parmdecl type", arg.Kind())
		}

		name := arg.Spelling()
		if name == "" {
			log.Fatal("no param name")
		}

		ty := p.parseType(fmt.Sprintf("%v[%v]", indent, name), arg.Type())

		args = append(args, &Param{
			Name: name,
			Type: ty,
		})
	}

	p.mod.functions[name] = &Function{
		Name:   name,
		Args:   args,
		Result: result,
	}
	p.mod.functionOrder = append(p.mod.functionOrder, name)
}

func (p *Parser) parseType(indent string, t clang.Type) Type {

	switch t.Kind() {
	case clang.Type_Void:
		log.Println(indent, "Parsing type", t.Spelling(), "as void")
		return nil

	case clang.Type_Int, clang.Type_UInt, clang.Type_UChar, clang.Type_Char_S, clang.Type_Float, clang.Type_Double,
		clang.Type_Enum, clang.Type_Record:
		log.Println(indent, "Parsing type", t.Spelling(), "as ident")
		name := t.CanonicalType().Spelling()
		name = strings.TrimPrefix(name, "const ")
		name = strings.TrimPrefix(name, "struct ")
		name = strings.TrimPrefix(name, "enum ")

		if strings.HasPrefix(name, "unsigned ") {
			name = strings.TrimPrefix(name, "unsigned ")
			name = fmt.Sprintf("u%v", name)
		}

		if name == "" {
			log.Panicln(indent, "No name")
		}

		if strings.Contains(name, " ") {
			log.Panicln(indent, "name", name, "contains space")
		}

		return &IdentType{
			Name: name,
		}

	case clang.Type_Typedef:
		log.Println(indent, "Parsing type", t.Spelling(), "as typedef")

		name := t.DefName()
		if name == "" {
			log.Panicln(indent, "Unexpected def name", name)
		}

		return &IdentType{
			Name: t.DefName(),
		}

	case clang.Type_Pointer:
		log.Println(indent, "Parsing type", t.Spelling(), "as pointer")
		inner := p.parseType(indent, t.PointeeType())

		return &PointerType{
			Inner: inner,
		}

	case clang.Type_Elaborated:
		log.Println(indent, "Parsing type", t.Spelling(), "as elaborated")
		return p.parseType(indent, t.NamedType())

	case clang.Type_FunctionProto:
		log.Println(indent, "Parsing type", t.Spelling(), "as funcproto")

		var args []Type

		for i := int32(0); i < t.NumArgTypes(); i++ {
			arg := t.ArgType(uint32(i))
			parsed := p.parseType(indent, arg)
			args = append(args, parsed)
		}

		result := p.parseType(indent, t.ResultType())

		return &FuncType{
			Args:   args,
			Result: result,
		}

	case clang.Type_ConstantArray:
		log.Println(indent, "Parsing type", t.Spelling(), "as const array")
		inner := p.parseType(indent, t.ArrayElementType())

		return &ConstArray{
			Inner: inner,
			Size:  t.ArraySize(),
		}

	case clang.Type_IncompleteArray:
		log.Println(indent, "Parsing type", t.Spelling(), "as array")
		inner := p.parseType(indent, t.ArrayElementType())

		return &Array{
			Inner: inner,
		}

	default:
		log.Println(indent, "Parsing type", t.Spelling(), "as ???")
		log.Panicln(indent, "Unknown type", t.Kind().Spelling())
		return nil
	}
}

func (p *Parser) parseEnum(indent string, c clang.Cursor) {
	log.Println("enum", "name", c.Spelling())

	name := c.Spelling()

	if val, ok := p.mod.enums[name]; ok && len(val.Constants) > 0 {
		log.Println("already exists")

		return
	}

	enum := &Enum{
		Name: name,
	}

	c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() != clang.Cursor_EnumConstantDecl {
			log.Panicln("Unknown enum type", "kind", cursor.Kind().String())
		}

		enum.Constants = append(enum.Constants, cursor.Spelling())

		return clang.ChildVisit_Continue
	})

	p.mod.enumOrder = slices.DeleteFunc(p.mod.enumOrder, func(s string) bool {
		return s == name
	})

	p.mod.enums[name] = enum
	p.mod.enumOrder = append(p.mod.enumOrder, name)
}

func (p *Parser) parseStruct(indent string, c clang.Cursor) {
	log.Println("struct", "name", c.Spelling())

	name := c.Spelling()

	if val, ok := p.mod.structs[name]; ok && len(val.Fields) > 0 {
		log.Println("already exists")

		return
	}

	s := &Struct{
		Name: name,
	}

	c.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() == clang.Cursor_FieldDecl {
			name := cursor.Spelling()
			if name == "" {
				log.Fatal("no field name")
			}

			ty := p.parseType(fmt.Sprintf("%v[%v]", indent, name), cursor.Type())

			s.Fields = append(s.Fields, &Field{
				Name: name,
				Type: ty,
			})
		}

		return clang.ChildVisit_Continue
	})

	p.mod.structOrder = slices.DeleteFunc(p.mod.structOrder, func(s string) bool {
		return s == name
	})

	p.mod.structs[name] = s
	p.mod.structOrder = append(p.mod.structOrder, name)
}

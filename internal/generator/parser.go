package main

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/go-clang/bootstrap/clang"
)

const AVLibPath = "/opt/homebrew/Cellar/ffmpeg/6.0_1/include/"

var files = []string{
	"libavcodec/avcodec.h",
}

func Parse() *Module {
	p := &Parser{
		mod: &Module{
			functions: make(map[string]*Function),
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
		log.Println("typedef", "name", c.Spelling())

		log.Println(" ", c.TypedefDeclUnderlyingType().Spelling())
		log.Println(" ", c.TypedefDeclUnderlyingType().Kind())
		log.Println(" ", c.TypedefDeclUnderlyingType().CanonicalType().Spelling())
		log.Println(" ", c.TypedefDeclUnderlyingType().CanonicalType().Kind())

		dec := c.TypedefDeclUnderlyingType().CanonicalType().Declaration()

		dec.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
			log.Println("  Inner ", "kind", cursor.Kind().String(), "name", cursor.Spelling())

			return clang.ChildVisit_Continue
		})

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

	case clang.Cursor_MacroExpansion, clang.Cursor_InclusionDirective, clang.Cursor_EnumDecl, clang.Cursor_StructDecl,
		clang.Cursor_UnionDecl:

	default:
		log.Panicln("Unexpected top level", "kind", c.Kind().String())
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

	case clang.Type_Int, clang.Type_UInt, clang.Type_UChar, clang.Type_Char_S, clang.Type_Enum, clang.Type_Record:
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

	default:
		log.Println(indent, "Parsing type", t.Spelling(), "as ???")
		log.Panicln(indent, "Unknown type", t.Kind().Spelling())
		return nil
	}
}

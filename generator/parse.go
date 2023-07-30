package generator

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/csnewman/ffmpeg-go/generator/parser"
)

var printfMatcher = regexp.MustCompile(`av_printf_format\(\d+,\s+\d+\)`)

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lastStatic := 10
	blockDepth := 0

	content := ""

	for scanner.Scan() {
		lastStatic++

		line := scanner.Text()

		if blockDepth > 0 {
			if line == "{" {
				blockDepth++
			}

			if line == "}" {
				blockDepth--
			}

			continue
		}

		if strings.HasPrefix(line, "static ") {
			lastStatic = 0

			if strings.HasSuffix(line, "{") {
				blockDepth = 1
				continue
			}
		}

		if line == "{" && lastStatic <= 2 {
			blockDepth = 1
			continue
		}

		content += line + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	content = strings.ReplaceAll(content, "av_warn_unused_result", "")
	content = strings.ReplaceAll(content, "av_always_inline", "")
	content = printfMatcher.ReplaceAllString(content, "")

	return content, nil
}

func (f *File) processUnit(unit parser.IUnitContext) {
	for _, context := range unit.AllTopLevel() {
		f.processTopLevel(context)
	}
}

func (f *File) processTopLevel(topLevel parser.ITopLevelContext) {
	switch c := topLevel.(type) {
	case *parser.TopLevelIfndefContext:
		log.Println("TopLevelIfndefContext:enter", strings.TrimSpace(c.Ifndef().GetText()))

		// Assume not defined
		for _, context := range c.AllTopLevel() {
			f.processTopLevel(context)
		}

		log.Println("TopLevelIfndefContext:exit")

	case *parser.TopLevelIfdirContext:
		log.Println("TopLevelIfdirContext:enter", strings.TrimSpace(c.Ifdir().GetText()))
		// Assume not defined
		log.Println("TopLevelIfdirContext:exit")

	case *parser.TopLevelDefineContext:
		log.Println("TopLevelDefineContext", strings.TrimSpace(c.GetText()))

	case *parser.TopLevelIncludeContext:
		log.Println("TopLevelIncludeContext", strings.TrimSpace(c.GetText()))
		// Ignore

	case *parser.TopLevelEnumContext:
		f.processEnum(c)

	case *parser.TopLevelEnumTypedefContext:
		f.processEnumTypedef(c)

	case *parser.TopLevelStructTypeDefContext:
		f.processStructTypeDef(c)

	case *parser.TopLevelStructContext:
		f.processStruct(c)

	case *parser.TopLevelFunctionContext:
		f.processFunction(c)

	case *parser.TopLevelStaticFunctionContext:
		log.Println("TopLevelStaticFunctionContext", strings.TrimSpace(c.GetText()))

	case *parser.TopLevelTypeDefFuncContext:
		log.Println("TopLevelTypeDefFuncContext", strings.TrimSpace(c.GetText()))

	case *parser.TopLevelTypeDefSimpleContext:
		f.processTypeDef(c)

	default:
		log.Panicln("Unknown type", reflect.TypeOf(topLevel), topLevel.GetText())
	}
}

func (f *File) processTypeDef(c *parser.TopLevelTypeDefSimpleContext) {
	log.Println("TopLevelTypeDefSimpleContext", strings.TrimSpace(c.GetText()))
	name := c.GetItemname().GetText()
	tdef := processType(c.Type_())

	switch tdef.Namespace {
	case IdentNamespaceStruct:
		sct, exists := f.structs[name]

		if !exists {
			sct = &Struct{
				Name: name,
			}

			f.structOrder = append(f.structOrder, name)
			f.structs[name] = sct
		}
	}

}

func (f *File) processEnum(c *parser.TopLevelEnumContext) {
	if c.GetItemname() == nil {
		log.Println("Ignoring unnamed enum")

		return
	}

	name := c.GetItemname().GetText()
	log.Println("TopLevelEnumContext", name)

	enum := &Enum{
		Name:   name,
		Values: f.parseEnumBody(c.EnumBody()),
	}

	f.enumOrder = append(f.enumOrder, name)
	f.enums[name] = enum
}

func (f *File) processEnumTypedef(c *parser.TopLevelEnumTypedefContext) {
	name := c.GetItemname().GetText()
	log.Println("TopLevelEnumTypedefContext", name)

	enum := &Enum{
		Name:       name,
		Values:     f.parseEnumBody(c.EnumBody()),
		IsAnonName: true,
	}

	f.enumOrder = append(f.enumOrder, name)
	f.enums[name] = enum
}

func (f *File) parseEnumBody(body parser.IEnumBodyContext) []string {
	var values []string

	for body != nil {
		var item parser.IEnumItemContext

		switch b := body.(type) {
		case *parser.EnumBodyIfdirContext:
			body = b.EnumBody()
			continue

		case *parser.EnumBodyDefineContext:
			body = b.EnumBody()
			// TODO: Implement
			continue

		case *parser.EnumBodyItemContext:
			body = b.EnumBody()
			item = b.EnumItem()

		default:
			log.Panicln("Unknown type", reflect.TypeOf(body))
		}

		var valName string

		switch i := item.(type) {
		case *parser.EnumItemNoValueContext:
			valName = i.Identifier().GetText()
		case *parser.EnumItemConstantContext:
			valName = i.Identifier().GetText()
		case *parser.EnumItemMappedContext:
			valName = i.Identifier(0).GetText()
		default:
			log.Panicln("Unknown type", reflect.TypeOf(item))
		}

		log.Println("item", valName)
		values = append(values, valName)
	}

	return values
}

func (f *File) processStructTypeDef(c *parser.TopLevelStructTypeDefContext) {
	name := c.GetItemname().GetText()
	log.Println("TopLevelStructTypeDefContext", name)

	sct, exists := f.structs[name]

	if !exists {
		sct = &Struct{
			Name: name,
		}

		f.structOrder = append(f.structOrder, name)
		f.structs[name] = sct
	}

}

func (f *File) processStruct(c *parser.TopLevelStructContext) {
	name := c.GetItemname().GetText()
	log.Println("TopLevelStructContext", name)

	sct, exists := f.structs[name]

	if !exists {
		sct = &Struct{
			Name: name,
		}

		f.structOrder = append(f.structOrder, name)
		f.structs[name] = sct
	}
}

func (f *File) processFunction(c *parser.TopLevelFunctionContext) {
	name := c.GetItemname().GetText()
	log.Println("TopLevelFunctionContext", name)

	fn := &Func{
		Name:       name,
		ReturnType: processType(c.Type_()),
	}

	params := c.Params()
	switch p := params.(type) {
	case *parser.ParamsVoidContext:
		// Do nothing
	case *parser.ParamsValuesContext:
		for _, rp := range p.AllParam() {
			p := parseParam(rp)
			fn.Params = append(fn.Params, p)
		}
	default:
		log.Panicln("Unknown type", reflect.TypeOf(params))
	}

	f.functionOrder = append(f.functionOrder, name)
	f.functions[name] = fn
}

func parseParam(rp parser.IParamContext) *Param {
	switch r := rp.(type) {
	case *parser.ParamSimpleContext:
		return &Param{
			Name: r.Identifier().GetText(),
			Type: processType(r.Type_()),
		}
	case *parser.ParamArrayContext:
		return &Param{
			Name: r.Identifier().GetText(),
			Type: &Type{
				Namespace: IdentNamespaceArray,
				Inner:     processType(r.Type_()),
				// TODO: Extract size
			},
		}
	case *parser.ParamFunctionContext:
		return &Param{
			Name: r.Identifier().GetText(),
			Type: &Type{
				Namespace: IdentNamespaceFunction,
				// TODO: Extract def
			},
		}
	default:
		log.Panicln("Unknown type", reflect.TypeOf(rp), rp.GetText())
		return nil
	}
}

func processType(t parser.ITypeContext) *Type {
	switch it := t.(type) {
	case *parser.TypePassContext:
		return processInnerType(it.TypeInner())
	case *parser.TypePtrContext:
		return &Type{
			Namespace: IdentNamespacePtr,
			Name:      "",
			Inner:     processType(it.Type_()),
		}
	case *parser.TypePtrConstContext:
		return &Type{
			Namespace: IdentNamespacePtr,
			Name:      "",
			Inner:     processType(it.Type_()),
		}
	default:
		log.Panicln("Unknown type", reflect.TypeOf(t))
		return nil
	}
}

func processInnerType(t parser.ITypeInnerContext) *Type {
	switch it := t.(type) {
	case *parser.TypeInnerIdentContext:
		return &Type{
			Namespace: IdentNamespaceDefault,
			Name:      it.Identifier().GetText(),
			Inner:     nil,
		}
	case *parser.TypeInnerEnumContext:
		return &Type{
			Namespace: IdentNamespaceEnum,
			Name:      it.Identifier().GetText(),
			Inner:     nil,
		}
	case *parser.TypeInnerStructContext:
		return &Type{
			Namespace: IdentNamespaceStruct,
			Name:      it.Identifier().GetText(),
			Inner:     nil,
		}
	case *parser.TypeInnerVoidContext:
		return &Type{
			Namespace: IdentNamespaceVoid,
			Name:      "",
			Inner:     nil,
		}
	case *parser.TypeInnerConstContext:
		return &Type{
			Namespace: IdentNamespaceConst,
			Name:      "",
			Inner:     processInnerType(it.TypeInner()),
		}
	case *parser.TypeInnerUnsignedContext:
		ident := ""
		if it.Identifier() != nil {
			ident = it.Identifier().GetText()
		}

		switch ident {
		case "", "int":
			ident = "uint"
		case "char":
			ident = "uchar"
		default:
			log.Panicln("Unknown type ident", ident)
		}

		return &Type{
			Namespace: IdentNamespaceDefault,
			Name:      ident,
			Inner:     nil,
		}
	default:
		log.Panicln("Unknown type", reflect.TypeOf(t))

		return nil
	}
}

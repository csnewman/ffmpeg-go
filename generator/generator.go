package generator

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/csnewman/ffmpeg-go/generator/parser"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

const AVLibPath = "/usr/include/x86_64-linux-gnu/"

var CTypeMap = map[string]string{
	"uint8_t":  "byte",
	"char":     "byte",
	"uchar":    "uint8",
	"uint32_t": "uint32",
	"uint64_t": "uint64",
	"uint":     "uint",
	"int":      "int",
	"int64_t":  "int64",
	"size_t":   "uintptr",
	"double":   "float64",
}

type GenIn struct {
	Pkg        string
	HeaderName string
}

var Files = []GenIn{
	{
		Pkg:        "avutil",
		HeaderName: "avutil",
	},
	{
		Pkg:        "avutil",
		HeaderName: "frame",
	},
	{
		Pkg:        "avutil",
		HeaderName: "buffer",
	},
	{
		Pkg:        "avutil",
		HeaderName: "log",
	},
	{
		Pkg:        "avutil",
		HeaderName: "dict",
	},
	{
		Pkg:        "avutil",
		HeaderName: "samplefmt",
	},
	{
		Pkg:        "avutil",
		HeaderName: "rational",
	},
	{
		Pkg:        "avutil",
		HeaderName: "pixfmt",
	},
	{
		Pkg:        "avcodec",
		HeaderName: "packet",
	},
	{
		Pkg:        "avcodec",
		HeaderName: "codec",
	},
	{
		Pkg:        "avcodec",
		HeaderName: "codec_id",
	},
	{
		Pkg:        "avcodec",
		HeaderName: "codec_par",
	},
	{
		Pkg:        "avcodec",
		HeaderName: "avcodec",
	},
	{
		Pkg:        "avformat",
		HeaderName: "avformat",
	},
	{
		Pkg:        "avformat",
		HeaderName: "avio",
	},
}

func Testing() {
	goMapped := map[string]string{}

	for _, file := range Files {
		err := ProcessFile(
			AVLibPath+"lib"+file.Pkg+"/"+file.HeaderName+".h",
			file.Pkg,
			"lib"+file.Pkg,
			[]string{"lib" + file.Pkg + "/" + file.HeaderName + ".h"},
			[]string{
				"av_buffer_ref", "av_hex_dump", "av_pkt_dump2", "av_log", "avio_printf",
			},
			"av/"+file.Pkg+"_"+file.HeaderName+".gen.go",
			goMapped,
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("done")
}

func ProcessFile(
	path string,
	pkg string,
	lib string,
	imports []string,
	skipList []string,
	dst string,
	goMapped map[string]string,
) error {
	content, err := readFile(path)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dst+".temp", []byte(content), 0644); err != nil {
		return err
	}

	input := antlr.NewInputStream(content)

	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	u := p.Unit()

	f := &File{
		functions: map[string]*Func{},
		structs:   map[string]*Struct{},
		enums:     map[string]*Enum{},
	}
	f.processUnit(u)

	g := &Generator{
		pkg:      pkg,
		lib:      lib,
		imports:  imports,
		skipList: skipList,
		goMapped: goMapped,
	}
	g.process(f)

	if err := os.WriteFile(dst, []byte(g.output), 0644); err != nil {
		return err
	}

	return nil
}

type Generator struct {
	output         string
	indent         int
	pkg            string
	lib            string
	imports        []string
	skipList       []string
	neededPackages []string
	goMapped       map[string]string
}

func (g *Generator) process(f *File) {
	for _, name := range f.enumOrder {
		g.WriteLine("")
		g.processEnum(f.enums[name])
	}

	for _, name := range f.structOrder {
		g.WriteLine("")
		g.processStruct(f.structs[name])
	}

	for _, name := range f.functionOrder {
		g.WriteLine("")
		g.processFuncHandleError(f.functions[name])
	}

	body := g.output
	g.output = ""

	g.WriteLine("package av")
	g.WriteLine("")
	g.WriteLine("/*")
	g.WriteLine("#cgo pkg-config: " + g.lib)

	for _, imp := range g.imports {
		g.WriteLine("#include <" + imp + ">")
	}

	g.WriteLine("*/")
	g.WriteLine("import \"C\"")

	for _, np := range g.neededPackages {
		g.WriteLine("import \"" + np + "\"")
	}

	g.output += body

}

func (g *Generator) AddPkg(pkg string) {
	if slices.Contains(g.neededPackages, pkg) {
		return
	}

	g.neededPackages = append(g.neededPackages, pkg)
}

func (g *Generator) WriteLine(line string) {
	for i := 0; i < g.indent; i++ {
		g.output += "    "
	}

	g.output += line
	g.output += "\n"
}

func (g *Generator) Indent(by int) {
	g.indent += by
}

func CleanName(name string) string {
	return strings.TrimPrefix(name, "AV")
}

func (g *Generator) processEnum(e *Enum) {
	goName := CleanName(e.Name)
	g.goMapped[goName] = g.pkg

	g.WriteLine("// " + goName + " wraps " + e.Name + ".")

	if e.IsAnonName {
		g.WriteLine("type " + goName + " C." + e.Name)
	} else {
		g.WriteLine("type " + goName + " C.enum_" + e.Name)
	}

	g.WriteLine("")
	g.WriteLine("const (")
	g.Indent(1)

	for _, value := range e.Values {
		goValue := strings.TrimPrefix(value, "AV_")
		goValue = strings.TrimPrefix(goValue, "AV")
		goValue = strcase.ToCamel(goValue)

		g.WriteLine(goValue + " " + goName + " = C." + value)
	}

	g.Indent(-1)
	g.WriteLine(")")

	g.WriteLine("")
	g.WriteLine("func (s " + goName + ") String() string {")
	g.Indent(1)

	//g.WriteLine("switch s {")
	//g.Indent(1)

	//g.WriteLine("default:")
	//g.Indent(1)
	//g.WriteLine("return \"" + e.Name + "::UNKNOWN\"")
	//g.Indent(-1)

	//g.Indent(-1)
	//g.WriteLine("}")
	//g.WriteLine("")

	g.AddPkg("fmt")
	g.WriteLine("return fmt.Sprintf(\"" + e.Name + "(%d)\", s)")

	g.Indent(-1)
	g.WriteLine("}")

}

func (g *Generator) processStruct(s *Struct) {
	goName := CleanName(s.Name)
	g.goMapped[goName] = g.pkg

	g.WriteLine("// " + goName + " wraps " + s.Name + ".")
	g.WriteLine("type " + goName + " C.struct_" + s.Name)
}

func (g *Generator) processFuncHandleError(f *Func) {
	oldOutput := g.output
	oldIndent := g.indent
	oldNeeded := g.neededPackages

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)

			g.output = oldOutput
			g.indent = oldIndent
			g.neededPackages = oldNeeded

			g.WriteLine("// Skipping " + f.Name + " due to error.")

			errLines := strings.Split(strings.TrimSpace(fmt.Sprint(err)), "\n")
			for _, line := range errLines {
				g.WriteLine("// " + line)
			}
		}
	}()

	g.processFunc(f)
}

func (g *Generator) processFunc(f *Func) {
	if slices.Contains(g.skipList, f.Name) {
		g.WriteLine("// Skipping " + f.Name + " due to skip list.")
		return
	}

	goName := strings.TrimPrefix(f.Name, "av_")
	goName = strings.TrimPrefix(goName, "av")
	//goName = strings.TrimPrefix(goName, g.pkg+"_")

	params := ""
	hasThis := false

	var thisType *Type

	for i, param := range f.Params {
		if i == 0 && g.pkg == "avcodec" && param.Name == "pkt" && param.Type.Namespace == IdentNamespacePtr &&
			((param.Type.Inner.Namespace == IdentNamespaceDefault &&
				param.Type.Inner.Name == "AVPacket") ||
				(param.Type.Inner.Namespace == IdentNamespaceConst &&
					param.Type.Inner.Inner.Namespace == IdentNamespaceDefault &&
					param.Type.Inner.Inner.Name == "AVPacket")) {
			hasThis = true
			thisType = param.Type
			continue
		}

		if params != "" {
			params += ", "
		}

		newName := strcase.ToLowerCamel(param.Name)

		if newName == "type" {
			newName = "type_"
		}

		params += newName + " " + g.genType(param.Type)
	}

	thisParam := ""
	if hasThis {
		thisParam = "(self " + g.genType(thisType) + ") "
	}

	if hasThis {
		goName = strings.TrimPrefix(goName, "packet_")
		goName = strings.TrimSuffix(goName, "_packet")
	}

	goName = strcase.ToCamel(goName)

	g.WriteLine("// " + goName + " wraps " + f.Name + ".")

	if f.ReturnType.Namespace == IdentNamespaceVoid {
		g.WriteLine("func " + thisParam + goName + "(" + params + ") {")
	} else {
		ret := g.genType(f.ReturnType)
		g.WriteLine("func " + thisParam + goName + "(" + params + ") " + ret + " {")
	}

	g.Indent(1)

	args := ""
	var postProcessArgs []string

	for i, param := range f.Params {
		if i == 0 && g.pkg == "avcodec" && param.Name == "pkt" && param.Type.Namespace == IdentNamespacePtr &&
			((param.Type.Inner.Namespace == IdentNamespaceDefault &&
				param.Type.Inner.Name == "AVPacket") ||
				(param.Type.Inner.Namespace == IdentNamespaceConst &&
					param.Type.Inner.Inner.Namespace == IdentNamespaceDefault &&
					param.Type.Inner.Inner.Name == "AVPacket")) {
			hasThis = true
			thisType = param.Type

			args += "(*C.AVPacket)(self)"
			continue
		}

		if args != "" {
			args += ", "
		}

		newName := strcase.ToLowerCamel(param.Name)

		if newName == "type" {
			newName = "type_"
		}

		pt := param.Type

		switch pt.Namespace {
		case IdentNamespaceDefault:
			args += "C." + pt.Name + "(" + newName + ")"

		case IdentNamespaceEnum:
			args += "C.enum_" + pt.Name + "(" + newName + ")"

		case IdentNamespacePtr:
			inner := pt.Inner

			// Ignore const
			if inner.Namespace == IdentNamespaceConst {
				inner = inner.Inner
			}

			switch inner.Namespace {
			case IdentNamespaceVoid:
				args += newName

			case IdentNamespaceDefault:
				if inner.Name == "char" {
					g.WriteLine("var mapped" + newName + " *C." + inner.Name)
					g.WriteLine("if " + newName + " != nil {")
					g.Indent(1)
					g.WriteLine("mapped" + newName + " = C.CString(*" + newName + ")")
					g.WriteLine("// TODO: Fix string mem leak")
					g.Indent(-1)
					g.WriteLine("}")

					args += "mapped" + newName
					break
				}

				args += "(*C." + inner.Name + ")(" + newName + ")"

			case IdentNamespacePtr:
				ip := inner.Inner

				// Ignore const
				if ip.Namespace == IdentNamespaceConst {
					ip = ip.Inner
				}

				switch ip.Namespace {
				case IdentNamespaceDefault:
					g.WriteLine("var " + newName + "Outer **C." + ip.Name)
					g.WriteLine("var " + newName + "Inner *C." + ip.Name)
					g.WriteLine("if " + newName + " != nil {")
					g.Indent(1)
					g.WriteLine(newName + "Inner = (*C." + ip.Name + ")(*" + newName + ")")
					g.WriteLine(newName + "Outer = &" + newName + "Inner")
					g.Indent(-1)
					g.WriteLine("}")

					postProcessArgs = append(postProcessArgs,
						"if "+newName+" != nil {",
						"    *"+newName+" = ("+g.genType(inner)+")("+newName+"Inner)",
						"}",
					)

					args += newName + "Outer"
				default:
					log.Panicln("Unknown namespace", ip.Namespace)
				}

			default:
				log.Panicln("Unknown namespace", inner.Namespace)
			}

		default:
			log.Panicln("Unknown namespace", pt.Namespace)
		}
	}

	funcCall := "C." + f.Name + "(" + args + ")"

	if f.ReturnType.Namespace == IdentNamespaceVoid {
		g.WriteLine(funcCall)
	} else {
		g.WriteLine("res := " + funcCall)
	}

	for _, arg := range postProcessArgs {
		g.WriteLine(arg)
	}

	ret := f.ReturnType

	switch ret.Namespace {
	case IdentNamespaceVoid:

	case IdentNamespaceDefault:
		mapped := g.genType(ret)
		g.WriteLine("return " + mapped + "(res)")

	case IdentNamespaceEnum:
		mapped := g.genType(ret)
		g.WriteLine("return " + mapped + "(res)")

	case IdentNamespacePtr:
		inner := ret.Inner

		// Ignore const
		if inner.Namespace == IdentNamespaceConst {
			inner = inner.Inner
		}

		switch inner.Namespace {
		case IdentNamespaceDefault:
			if inner.Name == "uint8_t" {
				g.AddPkg("unsafe")
				g.WriteLine("return unsafe.Pointer(res)")
			} else if inner.Name == "char" {
				g.WriteLine("var mappedRes *string")
				g.WriteLine("if res != nil {")
				g.Indent(1)
				g.WriteLine("resStr := C.GoString(res)")
				g.WriteLine("mappedRes = &resStr")
				g.Indent(-1)
				g.WriteLine("}")

				g.WriteLine("return mappedRes")
			} else {
				mapped := g.genType(ret)

				g.WriteLine("return (" + mapped + ")(res)")
			}
		default:
			log.Panicln("Unknown namespace", inner.Namespace)
		}

	default:
		log.Panicln("Unknown namespace", ret.Namespace)
	}

	g.Indent(-1)
	g.WriteLine("}")
}

func (g *Generator) genType(t *Type) string {
	switch t.Namespace {
	case IdentNamespaceVoid:
		return ""
	case IdentNamespacePtr:
		inner := t.Inner

		// Ignore const
		if inner.Namespace == IdentNamespaceConst {
			inner = inner.Inner
		}

		switch inner.Namespace {
		case IdentNamespaceVoid:
			g.AddPkg("unsafe")
			return "unsafe.Pointer"

		case IdentNamespaceDefault:
			if inner.Name == "uint8_t" {
				g.AddPkg("unsafe")
				return "unsafe.Pointer"
			} else if inner.Name == "char" {
				return "*string"
			}
			//else if inner.Name == "size_t" {
			//	return "*uintptr"
			//}

			_, hasMap := CTypeMap[inner.Name]
			if hasMap {
				log.Panicln("Unsupported pointer to native type", inner.Name)
			}

			goName := CleanName(inner.Name)

			return "*" + goName

		case IdentNamespacePtr:
			return "*" + g.genType(inner)

		default:
			return "not_imp:*" + string(inner.Namespace)
		}

	case IdentNamespaceDefault:
		if t.Name == "va_list" {
			log.Panicln("va_list detected")
		}

		mapped, hasMap := CTypeMap[t.Name]
		if hasMap {
			return mapped
		}

		goName := CleanName(t.Name)

		return goName

	case IdentNamespaceEnum:
		goName := CleanName(t.Name)

		return goName

	default:
		log.Panicln("Unknown namespace", t.Namespace)
		return "not_imp" + string(t.Namespace)
	}

}

package generator

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

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
	g := &Generator{
		skipList: []string{
			"av_buffer_ref", "av_hex_dump", "av_pkt_dump2", "av_log", "avio_printf",
		},
	}

	for _, file := range Files {
		g.neededPackages = nil

		f, err := parseFile(path.Join(AVLibPath, "lib"+file.Pkg, file.HeaderName+".h"))
		if err != nil {
			log.Fatalln("Failed to parse file", err)
		}

		w := g.process(f, file.Pkg, path.Join("lib"+file.Pkg, file.HeaderName+".h"))

		dst := path.Join("av", file.Pkg+"_"+file.HeaderName+".gen.go")

		if err := os.WriteFile(dst, []byte(w.Output()), 0644); err != nil {
			log.Fatalln("Failed to write file", f)
		}
	}

	log.Println("Complete")
}

type Generator struct {
	w              *Writer
	skipList       []string
	neededPackages []string
}

func (g *Generator) process(f *File, pkg string, imp string) *Writer {
	body := NewWriter()

	for _, name := range f.enumOrder {
		g.w = NewWriter()

		g.processEnum(f.enums[name])

		body.WriteLine("")
		body.WriteWriter(g.w)
	}

	for _, name := range f.structOrder {
		g.w = NewWriter()

		g.processStruct(f.structs[name])

		body.WriteLine("")
		body.WriteWriter(g.w)
	}

	for _, name := range f.functionOrder {
		g.w = NewWriter()
		oldNeeded := g.neededPackages

		err := g.processFunc(f.functions[name])

		body.WriteLine("")

		if err != nil {
			g.neededPackages = oldNeeded

			body.WriteLine("// Skipping " + name + " due to error.")
			body.WriteLine("// Error: " + err.Error())
		} else {
			body.WriteWriter(g.w)
		}
	}

	w := NewWriter()
	w.WriteLine("package av")
	w.WriteLine("")
	w.WriteLine("/*")
	w.WriteLine("#cgo pkg-config: lib" + pkg)
	w.WriteLine("#include <" + imp + ">")
	w.WriteLine("*/")
	w.WriteLine("import \"C\"")

	for _, np := range g.neededPackages {
		w.WriteLine("import \"" + np + "\"")
	}

	w.WriteWriter(body)

	return w
}

func (g *Generator) AddPkg(pkg string) {
	if slices.Contains(g.neededPackages, pkg) {
		return
	}

	g.neededPackages = append(g.neededPackages, pkg)
}

func CleanName(name string) string {
	goName := name
	goName = strings.TrimPrefix(goName, "AV_")
	goName = strings.TrimPrefix(goName, "AV")
	goName = strings.TrimPrefix(goName, "av_")
	goName = strings.TrimPrefix(goName, "av")

	return goName
}

func (g *Generator) processEnum(e *Enum) {
	goName := CleanName(e.Name)

	g.w.WriteLine("// " + goName + " wraps " + e.Name + ".")

	if e.IsAnonName {
		g.w.WriteLine("type " + goName + " C." + e.Name)
	} else {
		g.w.WriteLine("type " + goName + " C.enum_" + e.Name)
	}

	g.w.WriteLine("")
	g.w.WriteLine("const (")
	g.w.Indent(1)

	for _, value := range e.Values {
		goValue := CleanName(value)
		goValue = strcase.ToCamel(goValue)

		g.w.WriteLine(goValue + " " + goName + " = C." + value)
	}

	g.w.Indent(-1)
	g.w.WriteLine(")")
	g.w.WriteLine("")

	g.w.WriteLine("func (s " + goName + ") String() string {")
	g.w.Indent(1)
	g.AddPkg("fmt")
	g.w.WriteLine("return fmt.Sprintf(\"" + e.Name + "(%d)\", s)")
	g.w.Indent(-1)
	g.w.WriteLine("}")
}

func (g *Generator) processStruct(s *Struct) {
	goName := CleanName(s.Name)

	g.w.WriteLine("// " + goName + " wraps " + s.Name + ".")
	g.w.WriteLine("type " + goName + " C.struct_" + s.Name)
}

func (g *Generator) processFunc(f *Func) error {
	if slices.Contains(g.skipList, f.Name) {
		g.w.WriteLine("// Skipping " + f.Name + " due to skip list.")

		return nil
	}

	goName := CleanName(f.Name)

	params := ""
	hasThis := false

	var thisType *Type

	for i, param := range f.Params {
		if i == 0 && param.Name == "pkt" && param.Type.Namespace == IdentNamespacePtr &&
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

		genType, err := g.genType(param.Type)
		if err != nil {
			return err
		}

		params += newName + " " + genType
	}

	thisParam := ""
	if hasThis {
		genType, err := g.genType(thisType)
		if err != nil {
			return err
		}

		thisParam = "(self " + genType + ") "
	}

	if hasThis {
		goName = strings.TrimPrefix(goName, "packet_")
		goName = strings.TrimSuffix(goName, "_packet")
	}

	goName = strcase.ToCamel(goName)

	g.w.WriteLine("// " + goName + " wraps " + f.Name + ".")

	if f.ReturnType.Namespace == IdentNamespaceVoid {
		g.w.WriteLine("func " + thisParam + goName + "(" + params + ") {")
	} else {
		ret, err := g.genType(f.ReturnType)
		if err != nil {
			return err
		}

		g.w.WriteLine("func " + thisParam + goName + "(" + params + ") " + ret + " {")
	}

	g.w.Indent(1)

	args := ""
	var postProcessArgs []string

	for i, param := range f.Params {
		if i == 0 && param.Name == "pkt" && param.Type.Namespace == IdentNamespacePtr &&
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
					g.w.WriteLine("var mapped" + newName + " *C." + inner.Name)
					g.w.WriteLine("if " + newName + " != nil {")
					g.w.Indent(1)
					g.w.WriteLine("mapped" + newName + " = C.CString(*" + newName + ")")
					g.w.WriteLine("// TODO: Fix string mem leak")
					g.w.Indent(-1)
					g.w.WriteLine("}")

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
					g.w.WriteLine("var " + newName + "Outer **C." + ip.Name)
					g.w.WriteLine("var " + newName + "Inner *C." + ip.Name)
					g.w.WriteLine("if " + newName + " != nil {")
					g.w.Indent(1)
					g.w.WriteLine(newName + "Inner = (*C." + ip.Name + ")(*" + newName + ")")
					g.w.WriteLine(newName + "Outer = &" + newName + "Inner")
					g.w.Indent(-1)
					g.w.WriteLine("}")

					innerType, err := g.genType(inner)
					if err != nil {
						return err
					}

					postProcessArgs = append(postProcessArgs,
						"if "+newName+" != nil {",
						"    *"+newName+" = ("+innerType+")("+newName+"Inner)",
						"}",
					)

					args += newName + "Outer"
				default:
					return fmt.Errorf("unknown namespace: %s", ip.Namespace)
				}

			default:
				return fmt.Errorf("unknown namespace: %s", inner.Namespace)
			}

		default:
			return fmt.Errorf("unknown namespace: %s", pt.Namespace)
		}
	}

	funcCall := "C." + f.Name + "(" + args + ")"

	if f.ReturnType.Namespace == IdentNamespaceVoid {
		g.w.WriteLine(funcCall)
	} else {
		g.w.WriteLine("res := " + funcCall)
	}

	for _, arg := range postProcessArgs {
		g.w.WriteLine(arg)
	}

	ret := f.ReturnType

	switch ret.Namespace {
	case IdentNamespaceVoid:

	case IdentNamespaceDefault:
		mapped, err := g.genType(ret)
		if err != nil {
			return err
		}

		g.w.WriteLine("return " + mapped + "(res)")

	case IdentNamespaceEnum:
		mapped, err := g.genType(ret)
		if err != nil {
			return err
		}

		g.w.WriteLine("return " + mapped + "(res)")

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
				g.w.WriteLine("return unsafe.Pointer(res)")
			} else if inner.Name == "char" {
				g.w.WriteLine("var mappedRes *string")
				g.w.WriteLine("if res != nil {")
				g.w.Indent(1)
				g.w.WriteLine("resStr := C.GoString(res)")
				g.w.WriteLine("mappedRes = &resStr")
				g.w.Indent(-1)
				g.w.WriteLine("}")

				g.w.WriteLine("return mappedRes")
			} else {
				mapped, err := g.genType(ret)
				if err != nil {
					return err
				}

				g.w.WriteLine("return (" + mapped + ")(res)")
			}
		default:
			return fmt.Errorf("unknown namespace: %s", inner.Namespace)
		}

	default:
		return fmt.Errorf("unknown namespace: %s", ret.Namespace)
	}

	g.w.Indent(-1)
	g.w.WriteLine("}")

	return nil
}

func (g *Generator) genType(t *Type) (string, error) {
	switch t.Namespace {
	case IdentNamespaceVoid:
		return "", nil
	case IdentNamespacePtr:
		inner := t.Inner

		// Ignore const
		if inner.Namespace == IdentNamespaceConst {
			inner = inner.Inner
		}

		switch inner.Namespace {
		case IdentNamespaceVoid:
			g.AddPkg("unsafe")
			return "unsafe.Pointer", nil

		case IdentNamespaceDefault:
			if inner.Name == "uint8_t" {
				g.AddPkg("unsafe")
				return "unsafe.Pointer", nil
			} else if inner.Name == "char" {
				return "*string", nil
			}

			_, hasMap := CTypeMap[inner.Name]
			if hasMap {
				return "", fmt.Errorf("unsupported pointer to native type: %s", inner.Name)
			}

			goName := CleanName(inner.Name)
			return "*" + goName, nil

		case IdentNamespacePtr:
			gen, err := g.genType(inner)
			if err != nil {
				return "", err
			}

			return "*" + gen, nil

		default:
			return "", fmt.Errorf("unknown namespace: %s", inner.Namespace)
		}

	case IdentNamespaceDefault:
		if t.Name == "va_list" {
			return "", fmt.Errorf("va_list detected")
		}

		mapped, hasMap := CTypeMap[t.Name]
		if hasMap {
			return mapped, nil
		}

		goName := CleanName(t.Name)

		return goName, nil

	case IdentNamespaceEnum:
		goName := CleanName(t.Name)

		return goName, nil

	default:
		return "", fmt.Errorf("unknown namespace: %s", t.Namespace)
	}
}

package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

var primTypes = map[string]string{
	"int64_t":  "int64",
	"uint64_t": "uint64",
	"size_t":   "uint64",
}

type Generator struct {
	input *Module
}

func Gen(i *Module) {
	g := &Generator{
		input: i,
	}

	g.generateEnums()
	g.generateStructs()
	g.generateFuncs()
}

func newFile() *jen.File {
	o := jen.NewFile("ffmpeg")

	for _, file := range files {
		o.CgoPreamble(fmt.Sprintf("#include <%v>", file))
	}

	return o
}

func (g *Generator) generateEnums() {
	i := g.input

	o := newFile()

	for _, enumName := range i.enumOrder {
		enum := i.enums[enumName]

		log.Println("Generating enum", enum.Name)
		o.Commentf("--- Enum %v ---", enum.Name)
		o.Line()

		goName := enumName

		o.Commentf("%v wraps %v.", goName, enum.Name)
		o.Type().Id(goName).Qual("C", enum.Name)

		var valDefs []jen.Code

		for _, constant := range enum.Constants {
			constName := strcase.ToCamel(constant)

			valDefs = append(
				valDefs,
				jen.Commentf("%v wraps %v.", constName, constant),
				jen.Id(constName).Id(goName).Op("=").Qual("C", constant),
			)
		}

		if len(valDefs) > 0 {
			o.Const().Defs(valDefs...)
		}
	}

	err := o.Save("enums.gen.go")
	if err != nil {
		log.Panicln(err)
	}
}

func (g *Generator) generateStructs() {
	i := g.input

	o := newFile()

	for _, stName := range i.structOrder {
		st := i.structs[stName]

		log.Println("Generating struct", st.Name)
		o.Commentf("--- Struct %v ---", st.Name)
		o.Line()

		goName := st.Name

		o.Commentf("%v wraps %v.", goName, st.Name)

		o.Type().Id(goName).Struct(
			jen.Id("ptr").Op("*").Qual("C", st.Name),
		)

		for _, field := range st.Fields {
			fName := strcase.ToCamel(field.Name)

			o.Func().
				Params(jen.Id("s").Op("*").Id(goName)).
				Id(fName).
				Params().
				Block()
		}
	}

	err := o.Save("structs.gen.go")
	if err != nil {
		log.Panicln(err)
	}
}

var (
	fileType = &PointerType{
		Inner: &IdentType{Name: "FILE"},
	}
	vaListType = &IdentType{Name: "va_list"}
)

func (g *Generator) generateFuncs() {
	i := g.input

	o := newFile()

outer:
	for _, fnName := range i.functionOrder {
		fn := i.functions[fnName]

		log.Println("Generating func", fn.Name)
		o.Commentf("--- Function %v ---", fn.Name)
		o.Line()

		// Check if function contains unsupported features
		for _, arg := range fn.Args {
			skip := false

			if typeEquals(arg.Type, fileType) || typeEquals(arg.Type, vaListType) {
				skip = true
			}

			switch v := arg.Type.(type) {
			case *PointerType:
				switch v.Inner.(type) {
				case *FuncType:
					skip = true
				}
			}

			if skip {
				o.Commentf("%v skipped due to arg: %v.", fn.Name, arg.Type)
				o.Line()

				continue outer
			}
		}

		goName := convFuncName(fn.Name)

		o.Commentf("%v wraps %v.", goName, fn.Name)

		var (
			params []jen.Code
			args   []jen.Code
			body   []jen.Code
		)

		for _, arg := range fn.Args {
			pName := convParamName(arg.Name)

			switch v := arg.Type.(type) {
			case *IdentType:
				var goType *jen.Statement

				if m, ok := primTypes[v.Name]; ok {
					goType = jen.Id(m)
				} else {
					goType = jen.Id(v.Name)
				}

				params = append(params, jen.Id(pName).Add(goType))

				args = append(args, jen.Qual("C", v.Name).Params(jen.Id(pName)))

			case *PointerType:
				//params = append(params, jen.Id(pName).Id("int"))

				switch iv := v.Inner.(type) {
				case nil:
					params = append(params, jen.Id(pName).Id("TODO"))

				case *IdentType:

					if iv.Name == "char" {

						params = append(params, jen.Id(pName).Id("string"))

						//retType = []jen.Code{
						//	jen.Id("string"),
						//}
					} else if iv.Name == "uchar" {
						params = append(params, jen.Id(pName).Qual("unsafe", "Pointer"))

					} else if iv.Name == "uint8_t" {
						//retType = []jen.Code{
						//	jen.Qual("unsafe", "Pointer"),
						//}
						params = append(params, jen.Id(pName).Id("TODO"))
					} else {

						if m, ok := primTypes[iv.Name]; ok {
							//goType = jen.Id(m)
							params = append(params, jen.Id(pName).Op("*").Id(m))
						} else {
							//goType = jen.Id(iv.Name)

							params = append(params, jen.Id(pName).Op("*").Id(iv.Name))
						}

						//retType = []jen.Code{
						//	jen.Op("*").Id(iv.Name),
						//}
					}

				case *PointerType:

					switch iiv := iv.Inner.(type) {
					case *IdentType:

						params = append(params, jen.Id(pName).Op("*").Id(iiv.Name))

					default:

						params = append(params, jen.Id(pName).Id("TODO"))

					}

				default:
					log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
				}

			case *Array:
				params = append(params, jen.Id(pName).Id("TODO"))

			default:
				log.Panicln("unexpected type", reflect.TypeOf(arg.Type))
			}
		}

		cc := jen.Qual("C", fn.Name).Params(args...)

		var retType []jen.Code

		switch v := fn.Result.(type) {
		case nil:
			// nothing
			body = append(body, cc)

		case *IdentType:
			var goType *jen.Statement

			if m, ok := primTypes[v.Name]; ok {
				goType = jen.Id(m)
			} else {
				goType = jen.Id(v.Name)
			}

			retType = []jen.Code{
				goType,
			}

			body = append(body,
				jen.Id("ret").Op(":=").Add(cc),
				jen.Return(goType.Clone().Params(jen.Id("ret"))),
			)

		case *PointerType:
			switch iv := v.Inner.(type) {
			case *IdentType:

				if iv.Name == "char" {
					retType = []jen.Code{
						jen.Id("string"),
					}
				} else if iv.Name == "uint8_t" {
					retType = []jen.Code{
						jen.Qual("unsafe", "Pointer"),
					}
				} else {
					retType = []jen.Code{
						jen.Op("*").Id(iv.Name),
					}
				}

			default:
				log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
			}

			body = append(body,
				jen.Id("ret").Op(":=").Add(cc),
				//jen.Return(goType.Clone().Params(jen.Id("ret"))),
			)

			body = append(body, jen.Return(jen.Id("ret")))

		default:
			log.Panicln("unexpected type", reflect.TypeOf(fn.Result))
		}

		o.Func().
			//Params(jen.Id("s").Op("*").Id(goName)).
			Id(goName).
			Params(params...).
			Add(retType...).
			Block(body...)
	}

	err := o.Save("functions.gen.go")
	if err != nil {
		log.Panicln(err)
	}
}

func convParamName(val string) string {
	val = strcase.ToLowerCamel(val)

	if val == "type" {
		val = fmt.Sprintf("_%v", val)
	}

	return val
}

func convFuncName(val string) string {
	hasAV := strings.HasPrefix(val, "av")

	if hasAV {
		val = strings.TrimPrefix(val, "av")
	}

	hasIO := strings.HasPrefix(val, "io")

	if hasIO {
		val = strings.TrimPrefix(val, "io")
	}

	val = strcase.ToCamel(val)

	if hasIO {
		val = fmt.Sprintf("IO%v", val)
	}

	if hasAV {
		val = fmt.Sprintf("AV%v", val)
	}

	return val
}

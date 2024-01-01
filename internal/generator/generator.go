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
	"char":     "uint8",
	"uint32_t": "uint32",
	"int64_t":  "int64",
	"uint64_t": "uint64",
	"size_t":   "uint64",
	"double":   "float64",
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

		cName := enum.CName()

		o.Type().Id(goName).Qual("C", cName)

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

		cName := st.CName()

		o.Type().Id(goName).Struct(
			jen.Id("ptr").Op("*").Qual("C", cName),
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
		if fn.Variadic {
			o.Commentf("%v skipped due to variadic arg.", fn.Name)
			o.Line()

			continue
		}

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

		goName := g.convFuncName(fn.Name)

		o.Commentf("%v wraps %v.", goName, fn.Name)

		var (
			params   []jen.Code
			args     []jen.Code
			body     []jen.Code
			postCall []jen.Code
		)

		for _, arg := range fn.Args {
			pName := convParamName(arg.Name)

			switch v := arg.Type.(type) {
			case *IdentType:
				var goType *jen.Statement

				cName := v.Name

				if m, ok := primTypes[v.Name]; ok {
					goType = jen.Id(m)
				} else if e, ok := g.input.enums[v.Name]; ok {
					cName = e.CName()

					goType = jen.Id(v.Name)
				} else if s, ok := g.input.structs[v.Name]; ok {
					//cName = e.CName()
					//
					//goType = jen.Id(v.Name)

					_ = s

					continue outer
				} else {
					goType = jen.Id(v.Name)
				}

				params = append(params, jen.Id(pName).Add(goType))

				args = append(args, jen.Qual("C", cName).Params(jen.Id(pName)))

			case *PointerType:
				//params = append(params, jen.Id(pName).Id("int"))

				switch iv := v.Inner.(type) {
				case nil:
					params = append(params, jen.Id(pName).Id("TODO"))

					continue outer

				case *IdentType:
					if iv.Name == "char" {
						params = append(params, jen.Id(pName).Op("*").Id("CStr"))
						convName := fmt.Sprintf("tmp%v", pName)

						body = append(
							body,
							jen.Var().Id(convName).Op("*").Qual("C", "char"),
							jen.If(jen.Id(pName).Op("!=").Id("nil")).Block(
								jen.Id(convName).Op("=").Id(pName).Dot("ptr"),
							),
						)

						args = append(args, jen.Id(convName))
					} else if iv.Name == "uchar" {
						params = append(params, jen.Id(pName).Qual("unsafe", "Pointer"))
						continue outer

					} else if iv.Name == "uint8_t" {
						//retType = []jen.Code{
						//	jen.Qual("unsafe", "Pointer"),
						//}
						params = append(params, jen.Id(pName).Id("TODO"))
						continue outer
					} else {

						if m, ok := primTypes[iv.Name]; ok {
							params = append(params, jen.Id(pName).Op("*").Id(m))
							continue outer

						} else if s, ok := g.input.structs[iv.Name]; ok {
							params = append(params, jen.Id(pName).Op("*").Id(iv.Name))

							convName := fmt.Sprintf("tmp%v", pName)

							body = append(
								body,
								jen.Var().Id(convName).Op("*").Qual("C", s.CName()),
								jen.If(jen.Id(pName).Op("!=").Id("nil")).Block(
									jen.Id(convName).Op("=").Id(pName).Dot("ptr"),
								),
							)

							args = append(args, jen.Id(convName))
						} else {
							//goType = jen.Id(iv.Name)

							params = append(params, jen.Id(pName).Op("*").Id(iv.Name))
							continue outer
						}

						//retType = []jen.Code{
						//	jen.Op("*").Id(iv.Name),
						//}
					}

				case *PointerType:

					switch iiv := iv.Inner.(type) {
					case *IdentType:

						if iiv.Name == "uint8_t" {
							params = append(params, jen.Id(pName).Id("TODO"))
							continue outer
						} else if iiv.Name == "char" {
							params = append(params, jen.Id(pName).Id("TODO"))
							continue outer
						} else {

							if m, ok := primTypes[iiv.Name]; ok {
								params = append(params, jen.Id(pName).Op("**").Id(m))
								continue outer
							} else if s, ok := g.input.structs[iiv.Name]; ok {
								params = append(params, jen.Id(pName).Op("**").Id(iiv.Name))

								ptrName := fmt.Sprintf("ptr%v", pName)
								tmpName := fmt.Sprintf("tmp%v", pName)
								oldName := fmt.Sprintf("oldTmp%v", pName)
								innerName := fmt.Sprintf("inner%v", pName)

								body = append(
									body,
									jen.Var().Id(ptrName).Op("**").Qual("C", s.CName()),
									jen.Var().Id(tmpName).Op("*").Qual("C", s.CName()),
									jen.Var().Id(oldName).Op("*").Qual("C", s.CName()),
									jen.If(jen.Id(pName).Op("!=").Id("nil")).Block(
										jen.Id(innerName).Op(":=").Op("*").Id(pName),
										jen.If(jen.Id(innerName).Op("!=").Id("nil")).Block(
											jen.Id(tmpName).Op("=").Id(innerName).Dot("ptr"),
											jen.Id(oldName).Op("=").Id(tmpName),
										),
										jen.Id(ptrName).Op("=").Op("&").Id(tmpName),
									),
								)

								postCall = append(
									postCall,
									jen.If(jen.Id(tmpName).Op("!=").Id(oldName).Op("&&").Id(pName).Op("!=").Id("nil")).Block(

										jen.If(jen.Id(tmpName).Op("!=").Id("nil")).Block(
											jen.Op("*").Id(pName).Op("=").Op("&").Id(iiv.Name).Values(jen.Dict{
												jen.Id("ptr"): jen.Id(tmpName),
											}),
										).Else().Block(
											jen.Op("*").Id(pName).Op("=").Id("nil"),
										),
									),
								)

								args = append(args, jen.Id(ptrName))
							} else {
								params = append(params, jen.Id(pName).Op("**").Id(iiv.Name))
								continue outer
							}
						}

						//params = append(params, jen.Id(pName).Op("**").Id(iiv.Name))

					default:

						params = append(params, jen.Id(pName).Id("TODO"))
						continue outer

					}

				default:
					log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
				}

			case *Array:
				params = append(params, jen.Id(pName).Id("TODO"))
				continue outer

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
			body = append(body, postCall...)

		case *IdentType:
			var goType *jen.Statement

			if m, ok := primTypes[v.Name]; ok {
				goType = jen.Id(m)
			} else if s, ok := g.input.structs[v.Name]; ok {
				//cName = e.CName()
				//
				//goType = jen.Id(v.Name)

				_ = s

				continue outer
			} else {
				goType = jen.Id(v.Name)
			}

			retType = []jen.Code{
				goType,
			}

			body = append(body, jen.Id("ret").Op(":=").Add(cc))
			body = append(body, postCall...)
			body = append(body, jen.Return(goType.Clone().Params(jen.Id("ret"))))

		case *PointerType:
			body = append(body,
				jen.Id("ret").Op(":=").Add(cc),
			)
			body = append(body, postCall...)

			switch iv := v.Inner.(type) {
			case nil:
				retType = []jen.Code{
					jen.Qual("unsafe", "Pointer"),
				}
				body = append(body, jen.Return(jen.Id("ret")))
				continue outer

			case *IdentType:

				if iv.Name == "char" {
					retType = []jen.Code{
						jen.Op("*").Id("CStr"),
					}
					body = append(body, jen.Return(jen.Id("wrapCStr").Params(jen.Id("ret"))))
				} else if iv.Name == "uint8_t" {
					retType = []jen.Code{
						jen.Qual("unsafe", "Pointer"),
					}
					body = append(body, jen.Return(jen.Id("ret")))

					continue outer
				} else if _, ok := g.input.structs[iv.Name]; ok {
					retType = []jen.Code{
						jen.Op("*").Id(iv.Name),
					}

					body = append(
						body,
						jen.Var().Id("retMapped").Op("*").Id(iv.Name),
						jen.If(jen.Id("ret").Op("!=").Id("nil")).Block(
							jen.Id("retMapped").Op("=").Op("&").Id(iv.Name).Values(jen.Dict{
								jen.Id("ptr"): jen.Id("ret"),
							}),
						),
						jen.Return(jen.Id("retMapped")),
					)
				} else {
					retType = []jen.Code{
						jen.Op("*").Id(iv.Name),
					}
					body = append(body, jen.Return(jen.Id("ret")))
				}

			default:
				log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
			}

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

func (g *Generator) convFuncName(val string) string {
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

	// Temporary hack
	if _, ok := g.input.structs[val]; ok {
		val = fmt.Sprintf("%v_", val)
	}

	return val
}

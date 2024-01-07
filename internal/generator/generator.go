package main

import "C"
import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

var primTypes = map[string]string{
	"int":      "int",
	"uint":     "uint",
	"char":     "uint8",
	"uchar":    "uint8",
	"ulong":    "uint32",
	"uint8_t":  "uint8",
	"uint16_t": "uint16",
	"uint32_t": "uint32",
	"int64_t":  "int64",
	"uint64_t": "uint64",
	"size_t":   "uint64",
	"float":    "float32",
	"double":   "float64",
}

type Generator struct {
	input *Module
}

func Gen(i *Module) {
	g := &Generator{
		input: i,
	}

	g.generateConstants()
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

func (g *Generator) generateConstants() {
	i := g.input

	o := newFile()

	for _, constName := range i.constantOrder {
		constant := i.constants[constName]

		log.Println("Generating constant", constant.Name)

		goName := g.convCamel(constant.Name)

		if strings.HasPrefix(constName, "AVERROR_") {
			goName = fmt.Sprintf("%vConst", goName)
		}

		o.Commentf("%v wraps %v.", goName, constant.Name)

		o.Const().Id(goName).Op("=").Qual("C", constName)
	}

	err := o.Save("constants.gen.go")
	if err != nil {
		log.Panicln(err)
	}
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
			constName := g.convCamel(constant)

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

		if st.ByValue {
			o.Type().Id(goName).Struct(
				jen.Id("value").Qual("C", cName),
			)
		} else {
			o.Type().Id(goName).Struct(
				jen.Id("ptr").Op("*").Qual("C", cName),
			)

			o.Func().
				Params(jen.Id("s").Op("*").Id(goName)).
				Id("RawPtr").
				Params().
				Qual("unsafe", "Pointer").
				Block(jen.Return(jen.Qual("unsafe", "Pointer").Params(jen.Id("s").Dot("ptr"))))
		}

	fieldLoop:
		for _, field := range st.Fields {
			fName := strcase.ToCamel(field.Name)

			cName := field.Name
			if cName == "type" || cName == "range" {
				cName = fmt.Sprintf("_%v", cName)
			}

			var (
				getBody    []jen.Code
				getRetType []jen.Code
				getParams  []jen.Code
				setBody    []jen.Code
				setParams  []jen.Code
				tgt        *jen.Statement
			)

			genType := field.Type

			if field.BitWidth != -1 {
				o.Commentf("%v skipped due to bitfield", fName)
				o.Line()

				continue fieldLoop
			} else if iv, ok := field.Type.(*ConstArray); ok {

				fName = fmt.Sprintf("%vEntry", fName)

				getParams = append(getParams, jen.Id("i").Id("uint"))
				setParams = append(setParams, jen.Id("i").Id("uint"))

				if st.ByValue {
					getBody = append(
						getBody,
						jen.Id("value").Op(":=").Id("s").Dot("value").Dot(cName).Index(jen.Id("i")),
					)
					tgt = jen.Id("s").Dot("value").Dot(cName).Index(jen.Id("i"))
				} else {
					getBody = append(
						getBody,
						jen.Id("value").Op(":=").Id("s").Dot("ptr").Dot(cName).Index(jen.Id("i")),
					)
					tgt = jen.Id("s").Dot("ptr").Dot(cName).Index(jen.Id("i"))
				}

				genType = iv.Inner
			} else {
				if st.ByValue {
					getBody = append(
						getBody,
						jen.Id("value").Op(":=").Id("s").Dot("value").Dot(cName),
					)
					tgt = jen.Id("s").Dot("value").Dot(cName)
				} else {
					getBody = append(
						getBody,
						jen.Id("value").Op(":=").Id("s").Dot("ptr").Dot(cName),
					)
					tgt = jen.Id("s").Dot("ptr").Dot(cName)
				}
			}

			switch v := genType.(type) {

			case *IdentType:

				if m, ok := primTypes[v.Name]; ok {
					getRetType = []jen.Code{jen.Id(m)}
					setParams = append(setParams, jen.Id("value").Id(m))

					getBody = append(getBody, jen.Return(jen.Id(m).Params(jen.Id("value"))))

					if v.IsAnonEnum {
						setBody = append(
							setBody,
							tgt.Op("=").Id("value"),
						)
					} else {
						setBody = append(
							setBody,
							tgt.Op("=").Params(jen.Qual("C", v.Name)).Params(jen.Id("value")),
						)
					}
				} else if s, ok := g.input.structs[v.Name]; ok {
					if !s.ByValue {
						o.Commentf("%v skipped due to ident byvalue", fName)
						o.Line()

						continue fieldLoop
					}

					getRetType = []jen.Code{
						jen.Op("*").Id(s.Name),
					}
					setParams = append(setParams, jen.Id("value").Op("*").Id(s.Name))

					getBody = append(
						getBody,
						jen.Return(jen.Op("&").Id(s.Name).Values(jen.Dict{
							jen.Id("value"): jen.Id("value"),
						})),
					)

					setBody = append(
						setBody,
						tgt.Op("=").Id("value").Dot("value"),
					)
				} else if _, ok := g.input.callbacks[v.Name]; ok {
					o.Commentf("%v skipped due to ident callback", fName)
					o.Line()

					continue fieldLoop
				} else if e, ok := g.input.enums[v.Name]; ok {
					getRetType = []jen.Code{jen.Id(v.Name)}
					setParams = append(setParams, jen.Id("value").Id(v.Name))

					getBody = append(getBody, jen.Return(jen.Id(v.Name).Params(jen.Id("value"))))

					setBody = append(
						setBody,
						tgt.Op("=").Params(jen.Qual("C", e.CName())).Params(jen.Id("value")),
					)
				} else {
					log.Panicln("unexpected")
				}

			case *PointerType:
				switch iv := v.Inner.(type) {
				case nil:
					getRetType = []jen.Code{
						jen.Qual("unsafe", "Pointer"),
					}
					setParams = append(setParams, jen.Id("value").Qual("unsafe", "Pointer"))
					getBody = append(getBody, jen.Return(jen.Id("value")))
					setBody = append(setBody, tgt.Op("=").Id("value"))

				case *IdentType:

					if iv.Name == "URLContext" || iv.Name == "AVFilterCommand" {
						o.Commentf("%v skipped due to ptr to ignored type", fName)
						o.Line()

						continue fieldLoop
					} else if iv.Name == "char" {
						getRetType = []jen.Code{
							jen.Op("*").Id("CStr"),
						}
						setParams = append(setParams, jen.Id("value").Op("*").Id("CStr"))
						getBody = append(getBody, jen.Return(jen.Id("wrapCStr").Params(jen.Id("value"))))
						setBody = append(setBody, tgt.Op("=").Id("value").Dot("ptr"))

					} else if iv.Name == "uint8_t" {
						getRetType = []jen.Code{
							jen.Qual("unsafe", "Pointer"),
						}
						getBody = append(getBody, jen.Return(jen.Qual("unsafe", "Pointer").Params(jen.Id("value"))))

						setParams = append(setParams, jen.Id("value").Qual("unsafe", "Pointer"))
						setBody = append(setBody, tgt.Op("=").Params(jen.Op("*").Qual("C", iv.Name)).Params(jen.Id("value")))
					} else if _, ok := primTypes[iv.Name]; ok {
						o.Commentf("%v skipped due to prim ptr", fName)
						o.Line()

						continue fieldLoop
					} else if _, ok := g.input.enums[iv.Name]; ok {
						o.Commentf("%v skipped due to enum ptr", fName)
						o.Line()

						continue fieldLoop
					} else if _, ok := g.input.callbacks[iv.Name]; ok {
						o.Commentf("%v skipped due to callback ptr", fName)
						o.Line()

						continue fieldLoop
					} else if ist, ok := g.input.structs[iv.Name]; ok {
						if ist.ByValue {
							o.Commentf("%v skipped due to struct value ptr", fName)
							o.Line()

							continue fieldLoop
						}

						getRetType = []jen.Code{
							jen.Op("*").Id(iv.Name),
						}
						setParams = append(setParams, jen.Id("value").Op("*").Id(iv.Name))

						getBody = append(
							getBody,
							jen.Var().Id("valueMapped").Op("*").Id(iv.Name),
							jen.If(jen.Id("value").Op("!=").Id("nil")).Block(
								jen.Id("valueMapped").Op("=").Op("&").Id(iv.Name).Values(jen.Dict{
									jen.Id("ptr"): jen.Id("value"),
								}),
							),
							jen.Return(jen.Id("valueMapped")),
						)

						setBody = append(
							setBody,
							jen.If(jen.Id("value").Op("!=").Id("nil")).Block(
								tgt.Clone().Op("=").Id("value").Dot("ptr"),
							).Else().Block(
								tgt.Clone().Op("=").Id("nil"),
							),
						)
					} else {
						log.Panicln("unexpected")
					}

				case *FuncType:
					o.Commentf("%v skipped due to func ptr", fName)
					o.Line()

					continue fieldLoop

				case *PointerType:

					switch iiv := iv.Inner.(type) {
					case *IdentType:
						if st, ok := g.input.structs[iiv.Name]; ok {

							getRetType = []jen.Code{
								jen.Op("*").Id(iiv.Name),
							}

							cName := st.CName()
							getParams = append(getParams, jen.Id("i").Id("uint"))
							fName = fmt.Sprintf("%vEntry", fName)

							getBody = append(getBody,
								jen.Id("ptr").Op(":=").Id("arrayGet").
									Types(jen.Op("*").Qual("C", cName)).
									Params(
										jen.Id("value"),
										jen.Id("uintptr").Params(jen.Id("i")),
									),
								jen.Var().Id("valueMapped").Op("*").Id(iiv.Name),
								jen.If(jen.Id("ptr").Op("!=").Id("nil")).Block(
									jen.Id("valueMapped").Op("=").Op("&").Id(iiv.Name).Values(jen.Dict{
										jen.Id("ptr"): jen.Id("ptr"),
									}),
								),
								jen.Return(jen.Id("valueMapped")),
							)
						} else {
							o.Commentf("%v skipped due to unknown ptr ptr", fName)
							o.Line()

							continue fieldLoop
						}

					default:
						log.Panicln("unexpected type", reflect.TypeOf(iv.Inner))
					}

				default:
					log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
				}

			case *ConstArray:
				o.Commentf("%v skipped due to const array", fName)
				o.Line()

				continue fieldLoop

			case *UnionType:
				o.Commentf("%v skipped due to union type", fName)
				o.Line()

				continue fieldLoop

			default:
				log.Panicln("unexpected type", reflect.TypeOf(field.Type))
			}

			o.Func().
				Params(jen.Id("s").Op("*").Id(goName)).
				Id(fName).
				Params(getParams...).
				Add(getRetType...).
				Block(getBody...)

			o.Line()

			if len(setBody) > 0 {
				o.Func().
					Params(jen.Id("s").Op("*").Id(goName)).
					Id(fmt.Sprintf("Set%v", fName)).
					Params(setParams...).
					Block(setBody...)

				o.Line()
			}

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

		if typeEquals(fn.Result, fileType) {
			o.Commentf("%v skipped due to return.", fn.Name)
			o.Line()

			continue outer
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
				o.Commentf("%v skipped due to %v.", fn.Name, arg.Name)
				o.Line()

				continue outer
			}
		}

		goName := g.convCamel(fn.Name)

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
				if m, ok := primTypes[v.Name]; ok {
					params = append(params, jen.Id(pName).Id(m))
					args = append(args, jen.Qual("C", v.Name).Params(jen.Id(pName)))
				} else if e, ok := g.input.enums[v.Name]; ok {
					params = append(params, jen.Id(pName).Id(v.Name))
					args = append(args, jen.Qual("C", e.CName()).Params(jen.Id(pName)))
				} else if s, ok := g.input.structs[v.Name]; ok {
					if s.ByValue {
						params = append(params, jen.Id(pName).Op("*").Id(s.Name))
						args = append(args, jen.Id(pName).Dot("value"))
					} else {
						o.Commentf("%v skipped due to %v", fn.Name, pName)
						o.Line()

						continue outer
					}
				} else {
					params = append(params, jen.Id(pName).Id(v.Name))
					args = append(args, jen.Qual("C", v.Name).Params(jen.Id(pName)))
				}

			case *PointerType:
				switch iv := v.Inner.(type) {
				case nil:
					params = append(params, jen.Id(pName).Qual("unsafe", "Pointer"))
					args = append(args, jen.Id(pName))

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
					} else if iv.Name == "uint8_t" || iv.Name == "uchar" {
						params = append(params, jen.Id(pName).Qual("unsafe", "Pointer"))
						args = append(args, jen.Params(jen.Op("*").Qual("C", iv.Name)).Params(jen.Id(pName)))
					} else {

						if m, ok := primTypes[iv.Name]; ok {
							params = append(params, jen.Id(pName).Op("*").Id(m))

							o.Commentf("%v skipped due to %v", fn.Name, pName)
							o.Line()
							continue outer

						} else if s, ok := g.input.structs[iv.Name]; ok {
							if s.ByValue {
								o.Commentf("%v skipped due to %v", fn.Name, pName)
								o.Line()
								continue outer
							}

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

							o.Commentf("%v skipped due to %v", fn.Name, pName)
							o.Line()
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

							o.Commentf("%v skipped due to %v", fn.Name, pName)
							o.Line()
							continue outer
						} else if iiv.Name == "char" {
							params = append(params, jen.Id(pName).Id("TODO"))

							o.Commentf("%v skipped due to %v", fn.Name, pName)
							o.Line()
							continue outer
						} else {

							if m, ok := primTypes[iiv.Name]; ok {
								params = append(params, jen.Id(pName).Op("**").Id(m))

								o.Commentf("%v skipped due to %v", fn.Name, pName)
								o.Line()
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

								o.Commentf("%v skipped due to %v", fn.Name, pName)
								o.Line()
								continue outer
							}
						}

						//params = append(params, jen.Id(pName).Op("**").Id(iiv.Name))

					default:
						params = append(params, jen.Id(pName).Id("TODO"))

						o.Commentf("%v skipped due to %v", fn.Name, pName)
						o.Line()
						continue outer

					}

				default:
					log.Panicln("unexpected type", reflect.TypeOf(v.Inner))
				}

			case *Array:
				params = append(params, jen.Id(pName).Id("TODO"))

				o.Commentf("%v skipped due to %v", fn.Name, pName)
				o.Line()
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

			body = append(body, jen.Id("ret").Op(":=").Add(cc))
			body = append(body, postCall...)

			if v.Name == "int" {
				retType = []jen.Code{jen.Params(jen.Id("int"), jen.Id("error"))}
				body = append(
					body,
					jen.Return(
						jen.Id("int").Params(jen.Id("ret")).Op(",").
							Id("WrapErr").Params(jen.Id("int").Params(jen.Id("ret"))),
					),
				)
			} else if m, ok := primTypes[v.Name]; ok {
				retType = []jen.Code{jen.Id(m)}
				body = append(body, jen.Return(jen.Id(m).Params(jen.Id("ret"))))
			} else if s, ok := g.input.structs[v.Name]; ok {
				if s.ByValue {
					retType = []jen.Code{jen.Op("*").Id(v.Name)}
					body = append(
						body,
						jen.Return(jen.Op("&").Id(v.Name).Values(jen.Dict{
							jen.Id("value"): jen.Id("ret"),
						})),
					)
				} else {
					o.Commentf("%v skipped due to return", fn.Name)
					o.Line()
					continue outer
				}
			} else {
				retType = []jen.Code{jen.Id(v.Name)}
				body = append(body, jen.Return(jen.Id(v.Name).Params(jen.Id("ret"))))
			}

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
					body = append(
						body,
						jen.Return(jen.Qual("unsafe", "Pointer").Params(jen.Id("ret"))),
					)
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

func (g *Generator) convCamel(val string) string {
	hasAV := strings.HasPrefix(val, "av") || strings.HasPrefix(val, "AV")

	if hasAV {
		val = strings.TrimPrefix(val, "av")
		val = strings.TrimPrefix(val, "AV")
	}

	hasIO := strings.HasPrefix(val, "io") || strings.HasPrefix(val, "IO")

	if hasIO {
		val = strings.TrimPrefix(val, "io")
		val = strings.TrimPrefix(val, "IO")
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

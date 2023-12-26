package main

import (
	"fmt"
	"log"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

type Generator struct {
	input *Module
}

func Gen(i *Module) {
	g := &Generator{
		input: i,
	}

	g.generateEnums()
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

		log.Println("Generating", enum.Name)
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

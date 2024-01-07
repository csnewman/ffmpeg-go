package main

import (
	"log"
)

func main() {
	log.Println("Bindings generator")

	m := Parse()

	m.structs["AVRational"].ByValue = true
	m.enums["AVOptionType"].Comment = ""

	Gen(m)
}

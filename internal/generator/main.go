package main

import (
	"log"
)

func main() {
	log.Println("Bindings generator")

	m := Parse()

	m.structs["AVRational"].ByValue = true

	Gen(m)
}

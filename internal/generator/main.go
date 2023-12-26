package main

import (
	"log"
)

func main() {
	log.Println("Bindings generator")

	m := Parse()

	Gen(m)
}

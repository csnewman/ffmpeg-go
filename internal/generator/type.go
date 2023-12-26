package main

import "fmt"

type Type interface {
	typeMarker()

	String() string
}

type IdentType struct {
	Name string
}

func (t *IdentType) typeMarker()    {}
func (t *IdentType) String() string { return t.Name }

type PointerType struct {
	Inner Type
}

func (t *PointerType) typeMarker()    {}
func (t *PointerType) String() string { return fmt.Sprintf("*%v", t.Inner) }

type FuncType struct {
	Args   []Type
	Result Type
}

func (t *FuncType) typeMarker()    {}
func (t *FuncType) String() string { return "func" }

type ConstArray struct {
	Inner Type
	Size  int64
}

func (t *ConstArray) typeMarker()    {}
func (t *ConstArray) String() string { return fmt.Sprintf("%v[%v]", t.Inner, t.Size) }

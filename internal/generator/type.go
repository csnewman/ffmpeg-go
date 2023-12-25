package main

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
func (t *PointerType) String() string { return "*" + t.Inner.String() }

type FuncType struct {
	Args   []Type
	Result Type
}

func (t *FuncType) typeMarker()    {}
func (t *FuncType) String() string { return "func" }

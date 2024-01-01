package main

type Module struct {
	functionOrder []string
	functions     map[string]*Function

	structOrder []string
	structs     map[string]*Struct

	enumOrder []string
	enums     map[string]*Enum

	callbackOrder []string
	callbacks     map[string]*Function
}

type Function struct {
	Name     string
	Args     []*Param
	Result   Type
	Variadic bool
}

type Param struct {
	Name string
	Type Type
}

type Struct struct {
	Name     string
	Typedefd bool
	Fields   []*Field
}

type Field struct {
	Name string
	Type Type
}

type Enum struct {
	Name      string
	Typedefd  bool
	Constants []string
}

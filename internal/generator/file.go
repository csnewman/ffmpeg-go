package main

type Module struct {
	functionOrder []string
	functions     map[string]*Function
}

type Function struct {
	Name   string
	Args   []*Param
	Result Type
}

type Param struct {
	Name string
	Type Type
}

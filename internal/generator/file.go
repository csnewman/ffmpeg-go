package main

import "fmt"

type Module struct {
	functionOrder []string
	functions     map[string]*Function

	structOrder []string
	structs     map[string]*Struct

	enumOrder []string
	enums     map[string]*Enum

	callbackOrder []string
	callbacks     map[string]*Function

	constantOrder []string
	constants     map[string]*Constant
}

type Function struct {
	Name     string
	Args     []*Param
	Result   Type
	Variadic bool
	Ptr      bool
}

type Param struct {
	Name string
	Type Type
}

type Struct struct {
	Name     string
	Typedefd bool
	ByValue  bool
	Fields   []*Field
}

func (s *Struct) CName() string {
	if s.Typedefd {
		return s.Name
	}

	return fmt.Sprintf("struct_%v", s.Name)
}

type Field struct {
	Name     string
	Type     Type
	BitWidth int32
}

type Enum struct {
	Name      string
	Typedefd  bool
	Constants []string
}

func (s *Enum) CName() string {
	if s.Typedefd {
		return s.Name
	}

	return fmt.Sprintf("enum_%v", s.Name)
}

type Constant struct {
	Name string
}

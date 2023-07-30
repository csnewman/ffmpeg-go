package generator

type File struct {
	functionOrder []string
	functions     map[string]*Func
	structOrder   []string
	structs       map[string]*Struct
	enumOrder     []string
	enums         map[string]*Enum
}

type Struct struct {
	Name string
}

type Func struct {
	Name       string
	Params     []*Param
	ReturnType *Type
}

type Enum struct {
	Name       string
	Values     []string
	IsAnonName bool
}

type Param struct {
	Name string
	Type *Type
}

type IdentNamespace string

const (
	IdentNamespaceDefault  = "default"
	IdentNamespaceArray    = "array"
	IdentNamespaceFunction = "function"
	IdentNamespaceEnum     = "enum"
	IdentNamespaceStruct   = "struct"
	IdentNamespaceVoid     = "void"
	IdentNamespaceConst    = "const"
	IdentNamespacePtr      = "ptr"
)

type Type struct {
	Namespace IdentNamespace
	Name      string
	Inner     *Type
}

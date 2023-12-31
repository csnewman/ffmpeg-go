package main

import (
	"fmt"
	"github.com/go-clang/bootstrap/clang"
	"log"
	"strings"
)

type Type interface {
	typeMarker()

	String() string

	Equals(other Type) bool
}

func typeEquals(a Type, b Type) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Equals(b)
}

type IdentType struct {
	Name string
}

func (t *IdentType) typeMarker()    {}
func (t *IdentType) String() string { return t.Name }

func (t *IdentType) Equals(other Type) bool {
	ot, ok := other.(*IdentType)
	if !ok {
		return false
	}

	return t.Name == ot.Name
}

type PointerType struct {
	Inner Type
}

func (t *PointerType) typeMarker()    {}
func (t *PointerType) String() string { return fmt.Sprintf("*%v", t.Inner) }

func (t *PointerType) Equals(other Type) bool {
	op, ok := other.(*PointerType)
	if !ok {
		return false
	}

	return typeEquals(t.Inner, op.Inner)
}

type FuncType struct {
	Args   []Type
	Result Type
}

func (t *FuncType) typeMarker()    {}
func (t *FuncType) String() string { return "func" }

func (t *FuncType) Equals(other Type) bool {
	of, ok := other.(*FuncType)
	if !ok {
		return false
	}

	if !typeEquals(t.Result, of.Result) {
		return false
	}

	if len(t.Args) != len(of.Args) {
		return false
	}

	for i := 0; i < len(t.Args); i++ {
		if !typeEquals(t.Args[i], of.Args[i]) {
			return false
		}
	}

	return true
}

type ConstArray struct {
	Inner Type
	Size  int64
}

func (t *ConstArray) typeMarker()    {}
func (t *ConstArray) String() string { return fmt.Sprintf("%v[%v]", t.Inner, t.Size) }

func (t *ConstArray) Equals(other Type) bool {
	ot, ok := other.(*ConstArray)
	if !ok {
		return false
	}

	return t.Size == ot.Size && typeEquals(t.Inner, ot.Inner)
}

type Array struct {
	Inner Type
}

func (t *Array) typeMarker()    {}
func (t *Array) String() string { return fmt.Sprintf("%v[?]", t.Inner) }

func (t *Array) Equals(other Type) bool {
	ot, ok := other.(*Array)
	if !ok {
		return false
	}

	return typeEquals(t.Inner, ot.Inner)
}

func (p *Parser) parseType(indent string, t clang.Type, tokens *TokenCollection) Type {
	switch t.Kind() {
	case clang.Type_Void:
		log.Println(indent, "Parsing type", t.Spelling(), "as void")
		return nil

	case clang.Type_Int, clang.Type_UInt, clang.Type_Long, clang.Type_ULong, clang.Type_UChar, clang.Type_Char_S,
		clang.Type_Float, clang.Type_Double, clang.Type_Enum, clang.Type_Record:
		log.Println(indent, "Parsing type", t.Spelling(), "as ident")
		name := t.CanonicalType().Spelling()
		name = strings.TrimPrefix(name, "const ")
		name = strings.TrimPrefix(name, "struct ")
		name = strings.TrimPrefix(name, "enum ")

		if strings.HasPrefix(name, "unsigned ") {
			name = strings.TrimPrefix(name, "unsigned ")
			name = fmt.Sprintf("u%v", name)
		}

		if name == "" {
			log.Panicln(indent, "No name")
		}

		if strings.Contains(name, " ") {
			log.Panicln(indent, "name", name, "contains space")
		}

		if tokens != nil {
			tokens.TrimPrefix(clang.Token_Keyword, strptr("const"))
			tokens.TrimPrefix(clang.Token_Keyword, strptr("struct"))
			tokens.TrimPrefix(clang.Token_Keyword, strptr("enum"))
			wasUnsigned := tokens.TrimPrefix(clang.Token_Keyword, strptr("unsigned"))

			if len(tokens.tokens) != 1 {
				log.Panicln(indent, "Tokens != 1", tokens)
			}

			tok := tokens.PopEnd()

			if tok.kind != clang.Token_Identifier && tok.kind != clang.Token_Keyword {
				log.Panicln(indent, "Unexpected final token", tok)
			}

			tokenName := tok.text

			if wasUnsigned {
				tokenName = fmt.Sprintf("u%v", tokenName)
			}

			if tokenName != name {
				if tok.kind == clang.Token_Identifier {
					log.Println(indent, "Token name mismatch", tokenName, name, "- replacing")

					return &IdentType{
						Name: tokenName,
					}
				}

				log.Panicln(indent, "Token name mismatch", tokenName, name)
			}
		}

		return &IdentType{
			Name: name,
		}

	case clang.Type_Typedef:
		log.Println(indent, "Parsing type", t.Spelling(), "as typedef")

		name := t.DefName()
		if name == "" {
			log.Panicln(indent, "Unexpected def name", name)
		}

		return &IdentType{
			Name: t.DefName(),
		}

	case clang.Type_Pointer:
		log.Println(indent, "Parsing type", t.Spelling(), "as pointer")

		if tokens != nil {
			tok := tokens.PopEnd()

			if !tok.Is(clang.Token_Punctuation, strptr("*")) {
				log.Println(indent, "Punctuation mismatch, ignoring tokens", tokens)
				tokens = nil
			}
		}

		inner := p.parseType(indent, t.PointeeType(), tokens)

		return &PointerType{
			Inner: inner,
		}

	case clang.Type_Elaborated:
		log.Println(indent, "Parsing type", t.Spelling(), "as elaborated")
		return p.parseType(indent, t.NamedType(), nil)

	case clang.Type_FunctionProto:
		log.Println(indent, "Parsing type", t.Spelling(), "as funcproto")

		var args []Type

		for i := int32(0); i < t.NumArgTypes(); i++ {
			arg := t.ArgType(uint32(i))
			parsed := p.parseType(indent, arg, nil)
			args = append(args, parsed)
		}

		result := p.parseType(indent, t.ResultType(), nil)

		return &FuncType{
			Args:   args,
			Result: result,
		}

	case clang.Type_ConstantArray:
		log.Println(indent, "Parsing type", t.Spelling(), "as const array")
		inner := p.parseType(indent, t.ArrayElementType(), nil)

		return &ConstArray{
			Inner: inner,
			Size:  t.ArraySize(),
		}

	case clang.Type_IncompleteArray:
		log.Println(indent, "Parsing type", t.Spelling(), "as array")
		inner := p.parseType(indent, t.ArrayElementType(), nil)

		return &Array{
			Inner: inner,
		}

	default:
		log.Println(indent, "Parsing type", t.Spelling(), "as ???")
		log.Panicln(indent, "Unknown type", t.Kind().Spelling())
		return nil
	}
}

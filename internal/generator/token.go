package main

import (
	"fmt"
	"github.com/go-clang/bootstrap/clang"
	"strings"
)

func strptr(val string) *string {
	return &val
}

type Token struct {
	kind clang.TokenKind
	text string
}

func (t *Token) String() string {
	switch t.kind {
	case clang.Token_Punctuation:
		return fmt.Sprintf("<pun %v>", t.text)
	case clang.Token_Keyword:
		return fmt.Sprintf("<key %v>", t.text)
	case clang.Token_Identifier:
		return fmt.Sprintf("<idn %v>", t.text)
	case clang.Token_Literal:
		return fmt.Sprintf("<lit %v>", t.text)
	case clang.Token_Comment:
		return "<cmt>"
	default:
		return "<???>"
	}
}

func (t *Token) Is(kind clang.TokenKind, text *string) bool {
	if t == nil {
		return false
	}

	if kind != t.kind {
		return false
	}

	if text == nil {
		return true
	}

	return *text == t.text
}

type TokenCollection struct {
	tokens []*Token
}

func (c *TokenCollection) String() string {
	var sb strings.Builder

	for _, t := range c.tokens {
		sb.WriteString(t.String())
	}

	return sb.String()
}

func (c *TokenCollection) PeekEnd() *Token {
	if c == nil {
		return nil
	}

	l := len(c.tokens)

	if l == 0 {
		return nil
	}

	return c.tokens[l-1]
}

func (c *TokenCollection) PopEnd() *Token {
	if c == nil {
		return nil
	}

	l := len(c.tokens)

	if l == 0 {
		return nil
	}

	t := c.tokens[l-1]

	c.tokens = c.tokens[0 : l-1]

	return t
}

func (c *TokenCollection) PeekStart() *Token {
	if c == nil {
		return nil
	}

	l := len(c.tokens)

	if l == 0 {
		return nil
	}

	return c.tokens[0]
}

func (c *TokenCollection) PopStart() *Token {
	if c == nil {
		return nil
	}

	l := len(c.tokens)

	if l == 0 {
		return nil
	}

	t := c.tokens[0]

	c.tokens = c.tokens[1:]

	return t
}

func (c *TokenCollection) TrimPrefix(kind clang.TokenKind, text *string) bool {
	seen := false

	for {
		t := c.PeekStart()
		if t.Is(kind, text) {
			c.PopStart()
			seen = true

			continue
		}

		return seen
	}
}

func (c *TokenCollection) LeftCut(kind clang.TokenKind, text *string) bool {
	for i, token := range c.tokens {
		if token.Is(kind, text) {
			c.tokens = c.tokens[:i+1]

			return true
		}

	}

	return false
}

func (p *Parser) parseTokens(toks []clang.Token) *TokenCollection {
	var parsed []*Token

	for _, token := range toks {
		parsed = append(parsed, &Token{
			kind: token.Kind(),
			text: p.tu.TokenSpelling(token),
		})
	}

	p.tu.DisposeTokens(toks)

	return &TokenCollection{
		tokens: parsed,
	}
}

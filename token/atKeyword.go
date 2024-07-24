package token

import "fmt"

var _ Token = (*AtKeyword)(nil)

type AtKeyword struct {
	literal string
	subkind Subkind
}

func NewAtKeyword(literal string, subkind Subkind) *AtKeyword {
	return &AtKeyword{
		literal: literal,
		subkind: subkind,
	}
}

// Kind implements Token.
func (a *AtKeyword) Kind() Kind {
	return AT_KEYWORD
}

// Literal implements Token.
func (a *AtKeyword) Literal() string {
	return a.literal
}

// String implements Token.
func (a *AtKeyword) String() string {
	return fmt.Sprintf(kind_subkind_literal_template, a.Kind(), a.Subkind(), a.Literal())
}

// Subkind implements Token.
func (a *AtKeyword) Subkind() Subkind {
	return a.subkind
}

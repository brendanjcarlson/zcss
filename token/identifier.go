package token

import "fmt"

var _ Token = (*Identifier)(nil)

type Identifier struct {
	literal string
}

func NewIdentifier(literal string) *Identifier {
	return &Identifier{
		literal: literal,
	}
}

// Kind implements Token.
func (i *Identifier) Kind() Kind {
	return IDENTIFIER
}

// String implements Token.
func (i *Identifier) String() string {
	return fmt.Sprintf(kind_literal_template, i.Kind(), i.Literal())
}

// Literal implements Token.
func (i *Identifier) Literal() string {
	return i.literal
}

// Subkind implements Token.
func (i *Identifier) Subkind() Subkind {
	return SUBKIND_NONE
}

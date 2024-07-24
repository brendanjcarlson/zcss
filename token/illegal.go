package token

import "fmt"

var _ Token = (*Illegal)(nil)

type Illegal struct {
	literal string
}

func NewIllegal(literal string) *Illegal {
	return &Illegal{
		literal: literal,
	}
}

// Kind implements Token.
func (i *Illegal) Kind() Kind {
	return ILLEGAL
}

// Literal implements Token.
func (i *Illegal) Literal() string {
	return i.literal
}

// String implements Token.
func (i *Illegal) String() string {
	return fmt.Sprintf(kind_literal_template, i.Kind(), i.Literal())
}

// Subkind implements Token.
func (i *Illegal) Subkind() Subkind {
	return SUBKIND_NONE
}

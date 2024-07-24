package token

import "fmt"

var _ Token = (*Property)(nil)

type Property struct {
	literal string
}

func NewProperty(literal string) *Property {
	return &Property{
		literal: literal,
	}
}

// Kind implements Token.
func (p *Property) Kind() Kind {
	return PROPERTY
}

// Literal implements Token.
func (p *Property) Literal() string {
	return p.literal
}

// String implements Token.
func (p *Property) String() string {
	return fmt.Sprintf(kind_literal_template, p.Kind(), p.Literal())
}

// Subkind implements Token.
func (p *Property) Subkind() Subkind {
	return SUBKIND_NONE
}

package token

import "fmt"

var _ Token = (*Selector)(nil)

type Selector struct {
	literal string
	subkind Subkind
}

func NewSelector(literal string, subkind Subkind) *Selector {
	return &Selector{
		literal: literal,
		subkind: subkind,
	}
}

// Kind implements Token.
func (s *Selector) Kind() Kind {
	return SELECTOR
}

// Literal implements Token.
func (s *Selector) Literal() string {
	return s.literal
}

// String implements Token.
func (s *Selector) String() string {
	return fmt.Sprintf(kind_subkind_literal_template, s.Kind(), s.Subkind(), s.Literal())
}

// Subkind implements Token.
func (s *Selector) Subkind() Subkind {
	return s.subkind
}

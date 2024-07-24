package token

import "fmt"

var _ Token = (*Value)(nil)

type Value struct {
	literal string
	subkind Subkind
}

func NewValue(literal string, subkind Subkind) *Value {
	return &Value{
		literal: literal,
		subkind: subkind,
	}
}

// Kind implements Token.
func (v *Value) Kind() Kind {
	return VALUE
}

// Literal implements Token.
func (v *Value) Literal() string {
	return v.literal
}

// String implements Token.
func (v *Value) String() string {
	return fmt.Sprintf(kind_subkind_literal_template, v.Kind(), v.Subkind(), v.Literal())
}

// Subkind implements Token.
func (v *Value) Subkind() Subkind {
	return v.subkind
}

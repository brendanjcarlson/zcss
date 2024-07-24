package token

import "fmt"

var _ Token = (*Function)(nil)

type Function struct {
	literal string
	subkind Subkind
}

func NewFunction(literal string, subkind Subkind) *Function {
	return &Function{
		literal: literal,
		subkind: subkind,
	}
}

// Kind implements Token.
func (f *Function) Kind() Kind {
	return FUNCTION
}

// String implements Token.
func (f *Function) String() string {
	return fmt.Sprintf("[ KIND: %s | SUBKIND: %s | VALUE: %q ]", f.Kind(), f.Subkind(), f.Literal())
}

// Subkind implements Token.
func (f *Function) Subkind() Subkind {
	return f.subkind
}

// Value implements Token.
func (f *Function) Literal() string {
	return f.literal
}

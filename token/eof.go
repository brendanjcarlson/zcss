package token

import "fmt"

var _ Token = (*EOF)(nil)

type EOF struct{}

func NewEOF() *EOF {
	return &EOF{}
}

// Kind implements Token.
func (e *EOF) Kind() Kind {
	return END_OF_FILE
}

// String implements Token.
func (e *EOF) String() string {
	return fmt.Sprintf("[ KIND: %s ]", e.Kind())
}

// Subkind implements Token.
func (e *EOF) Subkind() Subkind {
	return SUBKIND_NONE
}

// Literal implements Token.
func (e *EOF) Literal() string {
	return ""
}

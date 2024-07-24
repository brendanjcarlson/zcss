package token

import "fmt"

var _ Token = (*CustomProperty)(nil)

type CustomProperty struct {
	literal string
}

func NewCustomProperty(literal string) *CustomProperty {
	return &CustomProperty{
		literal: literal,
	}
}

// Kind implements Token.
func (c *CustomProperty) Kind() Kind {
	return CUSTOM_PROPERTY
}

// Literal implements Token.
func (c *CustomProperty) Literal() string {
	return c.literal
}

// String implements Token.
func (c *CustomProperty) String() string {
	return fmt.Sprintf(kind_literal_template, c.Kind(), c.Literal())
}

// Subkind implements Token.
func (c *CustomProperty) Subkind() Subkind {
	return SUBKIND_NONE
}

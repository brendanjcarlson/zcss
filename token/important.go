package token

import "fmt"

var _ Token = (*Important)(nil)

type Important struct{}

func NewImportant() *Important {
	return &Important{}
}

// Kind implements Token.
func (i *Important) Kind() Kind {
	return IMPORTANT
}

// Literal implements Token.
func (i *Important) Literal() string {
	return "!important"
}

// String implements Token.
func (i *Important) String() string {
	return fmt.Sprintf(kind_literal_template, i.Kind(), i.Literal())
}

// Subkind implements Token.
func (i *Important) Subkind() Subkind {
	return SUBKIND_NONE
}

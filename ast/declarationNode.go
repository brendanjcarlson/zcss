package ast

import (
	"strings"

	"github.com/brendanjcarlson/zcss/token"
)

var _ Declaration = (*DeclarationNode)(nil)

type DeclarationNode struct {
	Property token.Token
	Value    []token.Token
}

// CSS implements Rule.
func (d *DeclarationNode) CSS(minified bool) string {
	out := new(strings.Builder)
	out.WriteString(d.Property.Literal())
	out.WriteString(":")
	if !minified {
		out.WriteString(" ")
	}
	for i, value := range d.Value {
		if i > 0 && i < len(d.Value) {
			out.WriteString(" ")
		}
		if value.Subkind() == token.FLOAT_LITERAL_VALUE {
			out.WriteString(strings.TrimLeft(value.Literal(), "0"))
		} else if value.Subkind() == token.DIMENSION_VALUE && strings.Contains(value.Literal(), ".") {
			out.WriteString(strings.TrimLeft(value.Literal(), "0"))
		} else {
			out.WriteString(value.Literal())
		}
	}
	return out.String()
}

// Literal implements Declaration.
func (d *DeclarationNode) Literal() string {
	out := new(strings.Builder)
	out.WriteString("(")
	out.WriteString(d.Property.Literal())
	out.WriteString(")")
	for _, value := range d.Value {
		out.WriteString("(")
		out.WriteString(value.Literal())
		out.WriteString(")")
	}
	return out.String()
}

// declaration implements Declaration.
func (d *DeclarationNode) declaration() {}

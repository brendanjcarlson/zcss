package ast

import "strings"

var _ Node = (*StyleSheetNode)(nil)

type StyleSheetNode struct {
	Rules []Rule
}

// CSS implements Node.
func (s *StyleSheetNode) CSS(minified bool) string {
	out := new(strings.Builder)
	for _, rule := range s.Rules {
		out.WriteString(rule.CSS(minified))
	}
	if !minified {
		out.WriteString("\n")
	}
	return out.String()
}

// Literal implements Node.
func (s *StyleSheetNode) Literal() string {
	out := new(strings.Builder)
	for _, rule := range s.Rules {
		out.WriteString(rule.Literal())
	}
	out.WriteString("\n")
	return out.String()
}

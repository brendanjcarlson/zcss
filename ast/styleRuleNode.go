package ast

import (
	"strings"

	"github.com/brendanjcarlson/zcss/token"
)

var _ Rule = (*StyleRuleNode)(nil)

type StyleRuleNode struct {
	Parent       Rule
	Selectors    []token.Token
	Declarations []Declaration
	Children     []Rule
}

// CSS implements Rule.
func (s *StyleRuleNode) CSS(minified bool) string {
	indentLevel := 0
	for parent := s.Parent; parent != nil; parent = parent.parent() {
		indentLevel++
	}

	out := new(strings.Builder)

	if !minified {
		out.WriteString("\n")
		for range indentLevel {
			out.WriteString("\t")
		}
	}
	for i, selector := range s.Selectors {
		if i > 0 {
			if !minified {
				out.WriteString(", ")
			} else {
				out.WriteString(",")
			}
		}
		out.WriteString(selector.Literal())
	}
	if !minified {
		out.WriteString(" ")
	}
	out.WriteString("{")
	if !minified {
		out.WriteString("\n")
	}

	for i, declaration := range s.Declarations {
		if !minified {
			for range indentLevel + 1 {
				out.WriteString("\t")
			}
		}
		out.WriteString(declaration.CSS(minified))
		if !minified {
			out.WriteString(";\n")
		} else if i < len(s.Declarations)-1 || len(s.Children) > 0 {
			out.WriteString(";")
		}
	}

	if len(s.Children) == 0 {
		if !minified {
			for range indentLevel {
				out.WriteString("\t")
			}
		}
		out.WriteString("}")
		if !minified {
			out.WriteString("\n")
		}
	} else {
		var j int
		for i, child := range s.Children {
			j = i
			out.WriteString(child.CSS(minified))
		}
		if !minified {
			out.WriteString("\n")
			for range indentLevel - j {
				out.WriteString("\t")
			}
		}
		out.WriteString("}")
	}

	return out.String()
}

// Literal implements Rule.
func (s *StyleRuleNode) Literal() string {
	indentLevel := 0
	for parent := s.Parent; parent != nil; parent = parent.parent() {
		indentLevel++
	}

	out := new(strings.Builder)

	out.WriteString("\n")
	for range indentLevel {
		out.WriteString("\t")
	}
	for i, selector := range s.Selectors {
		out.WriteString("(")
		if i > 0 {
			out.WriteString(" ")
		}
		out.WriteString(selector.Literal())
		out.WriteString(")")
	}
	out.WriteString("\n")

	for _, declaration := range s.Declarations {
		for range indentLevel + 1 {
			out.WriteString("\t")
		}
		out.WriteString("(")
		out.WriteString(declaration.Literal())
		out.WriteString(")\n")
	}

	if len(s.Children) == 0 {
		for range indentLevel {
			out.WriteString("\t")
		}
		out.WriteString("\n")
	} else {
		var j int
		for i, child := range s.Children {
			j = i
			out.WriteString(child.Literal())
		}
		out.WriteString("\n")
		for range indentLevel - j {
			out.WriteString("\t")
		}
	}

	return out.String()
}

// rule implements Rule.
func (s *StyleRuleNode) rule() {}

// parent implements Rule.
func (s *StyleRuleNode) parent() Rule {
	return s.Parent
}

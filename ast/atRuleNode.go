package ast

import "github.com/brendanjcarlson/zcss/token"

var _ Rule = (*AtRuleNode)(nil)

type AtRuleNode struct {
	Parent    Rule
	Token     token.Token
	Condition []token.Token
	Children  []Rule
}

// CSS implements Rule.
func (a *AtRuleNode) CSS(minified bool) string {
	return ""
}

// Literal implements Rule.
func (a *AtRuleNode) Literal() string {
	return ""
}

// rule implements Rule.
func (a *AtRuleNode) rule() {}

// parent implements Rule.
func (a *AtRuleNode) parent() Rule {
	return a.Parent
}

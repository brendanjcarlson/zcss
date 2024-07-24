package parser

import (
	"fmt"

	"github.com/brendanjcarlson/zcss/ast"
	"github.com/brendanjcarlson/zcss/lexer"
	"github.com/brendanjcarlson/zcss/token"
)

type Parser struct {
	tokens    []token.Token
	position  int
	cursor    int
	currToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	tokens := []token.Token{}

	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Kind() == token.END_OF_FILE {
			break
		}
	}

	p := &Parser{
		tokens: tokens,
	}

	p.consume()

	return p
}

func (p *Parser) ParseStyleSheet() *ast.StyleSheetNode {
	s := &ast.StyleSheetNode{}

	for p.currToken.Kind() != token.END_OF_FILE {
		switch p.currToken.Kind() {
		case token.AT_KEYWORD:
			s.Rules = append(s.Rules, p.consumeAtRule())
		case token.SELECTOR:
			s.Rules = append(s.Rules, p.consumeStyleRule(nil))
		default:
			panic(fmt.Sprintf("unexpected token: %s", p.currToken))
		}
	}

	return s
}

func (p *Parser) consume() {
	if p.cursor >= len(p.tokens) {
		p.currToken = token.NewEOF()
	} else {
		p.currToken = p.tokens[p.cursor]
	}
	p.position = p.cursor
	p.cursor += 1
}

func (p *Parser) consumeAtRule() *ast.AtRuleNode {
	r := &ast.AtRuleNode{}

	if p.currToken.Subkind() == token.MEDIA_AT_KEYWORD {
		r.Token = p.currToken
		p.consume()
		return p.consumeMediaAtRule(r)
	} else {
		return r
	}
}

func (p *Parser) consumeMediaAtRule(r *ast.AtRuleNode) *ast.AtRuleNode {
	return r
}

func (p *Parser) consumeStyleRule(parent ast.Rule) *ast.StyleRuleNode {
	r := &ast.StyleRuleNode{}
	if parent != nil {
		r.Parent = parent
	}

	r.Selectors = append(r.Selectors, p.currToken)
	p.consume()

	if p.currToken.Kind() != token.OPEN_CURLY_BRACKET {
		panic(fmt.Sprintf("expected open curly bracket, got %s", p.currToken))
	} else {
		p.consume()
	}

	for p.currToken.Kind() == token.PROPERTY {
		r.Declarations = append(r.Declarations, p.consumeDeclaration())
	}

	for p.currToken.Kind() == token.SELECTOR {
		r.Children = append(r.Children, p.consumeStyleRule(r))
	}

	if p.currToken.Kind() == token.CLOSE_CURLY_BRACKET {
		p.consume()
	} else {
		panic(fmt.Sprintf("expected close curly bracket, got %s", p.currToken))
	}

	return r
}

func (p *Parser) consumeDeclaration() *ast.DeclarationNode {
	d := &ast.DeclarationNode{}

	if p.currToken.Kind() == token.PROPERTY {
		d.Property = p.currToken
		p.consume()
	} else {
		panic(fmt.Sprintf("expected property, got %s", p.currToken))
	}

	if p.currToken.Kind() == token.COLON {
		p.consume()
	} else {
		panic(fmt.Sprintf("expected colon, got %s", p.currToken))
	}

	for p.currToken.Kind() == token.VALUE {
		d.Value = append(d.Value, p.currToken)
		p.consume()
	}

	if p.currToken.Kind() == token.SEMICOLON {
		p.consume()
	} else {
		panic(fmt.Sprintf("expected semicolon, got %s", p.currToken))
	}

	return d
}

func (p *Parser) peek() token.Token {
	if p.cursor >= len(p.tokens) {
		return token.NewEOF()
	} else {
		return p.tokens[p.cursor]
	}
}

func (p *Parser) peekN(offset int) token.Token {
	if p.cursor+offset >= len(p.tokens) {
		return token.NewEOF()
	} else {
		return p.tokens[p.cursor+offset]
	}
}

package token

import "fmt"

var _ Token = (*Colon)(nil)

type Colon struct{}

func NewColon() *Colon {
	return &Colon{}
}

// Kind implements Token.
func (c *Colon) Kind() Kind {
	return COLON
}

// Literal implements Token.
func (c *Colon) Literal() string {
	return ":"
}

// String implements Token.
func (c *Colon) String() string {
	return fmt.Sprintf(kind_literal_template, c.Kind(), c.Literal())
}

// Subkind implements Token.
func (c *Colon) Subkind() Subkind {
	return SUBKIND_NONE
}

var _ Token = (*Semicolon)(nil)

type Semicolon struct{}

func NewSemicolon() *Semicolon {
	return &Semicolon{}
}

// Kind implements Token.
func (s *Semicolon) Kind() Kind {
	return SEMICOLON
}

// Literal implements Token.
func (s *Semicolon) Literal() string {
	return ";"
}

// String implements Token.
func (s *Semicolon) String() string {
	return fmt.Sprintf(kind_literal_template, s.Kind(), s.Literal())
}

// Subkind implements Token.
func (s *Semicolon) Subkind() Subkind {
	return SUBKIND_NONE
}

var _ Token = (*Comma)(nil)

type Comma struct{}

func NewComma() *Comma {
	return &Comma{}
}

// Kind implements Token.
func (c *Comma) Kind() Kind {
	return COMMA
}

// Literal implements Token.
func (c *Comma) Literal() string {
	return ","
}

// String implements Token.
func (c *Comma) String() string {
	return fmt.Sprintf(kind_literal_template, c.Kind(), c.Literal())
}

// Subkind implements Token.
func (c *Comma) Subkind() Subkind {
	return SUBKIND_NONE
}

var _ Token = (*OpenCurlyBracket)(nil)

type OpenCurlyBracket struct{}

func NewOpenCurlyBracket() *OpenCurlyBracket {
	return &OpenCurlyBracket{}
}

// Kind implements Token.
func (o *OpenCurlyBracket) Kind() Kind {
	return OPEN_CURLY_BRACKET
}

// Literal implements Token.
func (o *OpenCurlyBracket) Literal() string {
	return "{"
}

// String implements Token.
func (o *OpenCurlyBracket) String() string {
	return fmt.Sprintf(kind_literal_template, o.Kind(), o.Literal())
}

// Subkind implements Token.
func (o *OpenCurlyBracket) Subkind() Subkind {
	return SUBKIND_NONE
}

var _ Token = (*CloseCurlyBracket)(nil)

type CloseCurlyBracket struct{}

func NewCloseCurlyBracket() *CloseCurlyBracket {
	return &CloseCurlyBracket{}
}

// Kind implements Token.
func (o *CloseCurlyBracket) Kind() Kind {
	return CLOSE_CURLY_BRACKET
}

// Literal implements Token.
func (o *CloseCurlyBracket) Literal() string {
	return "}"
}

// String implements Token.
func (o *CloseCurlyBracket) String() string {
	return fmt.Sprintf(kind_literal_template, o.Kind(), o.Literal())
}

// Subkind implements Token.
func (o *CloseCurlyBracket) Subkind() Subkind {
	return SUBKIND_NONE
}

var _ Token = (*OpenParenthesis)(nil)

type OpenParenthesis struct{}

func NewOpenParenthesis() *OpenParenthesis {
	return &OpenParenthesis{}
}

// Kind implements Token.
func (o *OpenParenthesis) Kind() Kind {
	return OPEN_PARENTHESIS
}

// Literal implements Token.
func (o *OpenParenthesis) Literal() string {
	return "("
}

// String implements Token.
func (o *OpenParenthesis) String() string {
	return fmt.Sprintf(kind_literal_template, o.Kind(), o.Literal())
}

// Subkind implements Token.
func (o *OpenParenthesis) Subkind() Subkind {
	return SUBKIND_NONE
}

var _ Token = (*CloseParenthesis)(nil)

type CloseParenthesis struct{}

func NewCloseParenthesis() *CloseParenthesis {
	return &CloseParenthesis{}
}

// Kind implements Token.
func (o *CloseParenthesis) Kind() Kind {
	return CLOSE_PARENTHESIS
}

// Literal implements Token.
func (o *CloseParenthesis) Literal() string {
	return ")"
}

// String implements Token.
func (o *CloseParenthesis) String() string {
	return fmt.Sprintf(kind_literal_template, o.Kind(), o.Literal())
}

// Subkind implements Token.
func (o *CloseParenthesis) Subkind() Subkind {
	return SUBKIND_NONE
}

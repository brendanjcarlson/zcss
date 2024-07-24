package lexer

import (
	"unicode"

	"github.com/brendanjcarlson/zcss/token"
)

type mode int

const (
	mode_top_level mode = iota
	mode_style_block
	mode_media_or_supports_condition
	mode_media_or_supports_block
	mode_layer_block
	mode_font_face_block
	mode_url
	mode_import
	mode_value
	mode_color
)

func (m mode) String() string {
	switch m {
	case mode_top_level:
		return "mode_top_level"
	case mode_style_block:
		return "mode_style_block"
	case mode_media_or_supports_condition:
		return "mode_media_or_supports_condition"
	case mode_media_or_supports_block:
		return "mode_media_or_supports_block"
	case mode_layer_block:
		return "mode_layer_block"
	case mode_font_face_block:
		return "mode_font_face_block"
	case mode_url:
		return "mode_url"
	case mode_import:
		return "mode_import"
	case mode_value:
		return "mode_value"
	case mode_color:
		return "mode_color"
	default:
		return "mode_unknown"
	}
}

type Lexer struct {
	input     string
	position  int
	cursor    int
	char      rune
	modeStack []mode
}

func New(input string) *Lexer {
	l := &Lexer{
		input:     input,
		position:  0,
		cursor:    0,
		char:      0,
		modeStack: make([]mode, 0),
	}
	l.consume()
	return l
}

func (l *Lexer) NextToken() token.Token {
	if l.char == 0 {
		return token.NewEOF()
	}

	switch l.currMode() {
	case mode_top_level:
		l.skipWhitespace()

		switch l.char {

		case 0:
			defer l.consume()
			return token.NewEOF()

		case '@':
			l.consume()
			literal := l.consumeAtKeyword()
			if subkind, ok := token.IsAtKeyword(literal); ok {
				if subkind == token.MEDIA_AT_KEYWORD || subkind == token.SUPPORTS_AT_KEYWORD {
					l.pushMode(mode_media_or_supports_condition)
				}
				return token.NewAtKeyword(literal, subkind)
			} else {
				return token.NewIllegal(literal)
			}

		case '{':
			defer l.consume()
			l.pushMode(mode_style_block)
			return token.NewOpenCurlyBracket()

		case '}':
			defer l.consume()
			l.popMode()
			return token.NewCloseCurlyBracket()

		default:
			if token.IsCssSelectorStart(l.char) {
				literal := l.consumeSelector()
				if subkind, ok := token.IsCssSelector(literal); ok {
					return token.NewSelector(literal, subkind)
				} else {
					defer l.consume()
					return token.NewIllegal(string(l.char))
				}
			} else {
				defer l.consume()
				return token.NewIllegal(string(l.char))
			}
		}

	case mode_style_block:
		l.skipWhitespace()

		switch {

		case l.char == 0:
			defer l.consume()
			return token.NewEOF()

		case l.char == '{':
			defer l.consume()
			return token.NewOpenCurlyBracket()

		case l.char == '}':
			defer l.consume()
			l.popMode()
			return token.NewCloseCurlyBracket()

		case l.char == ';':
			defer l.consume()
			return token.NewSemicolon()

		case l.char == ':':
			if !unicode.IsSpace(l.peek()) {
				literal := l.consumeSelector()
				if subkind, ok := token.IsCssSelector(literal); ok {
					l.pushMode(mode_style_block)
					return token.NewSelector(literal, subkind)
				} else {
					return token.NewIllegal(literal)
				}
			} else {
				defer l.consume()
				return token.NewColon()
			}

		case token.IsCssSelectorStart(l.char), l.char == '&':
			literal := l.consumeIdentifier()
			if _, ok := token.IsCssProperty(literal); ok {
				l.pushMode(mode_value)
				return token.NewProperty(literal)
			} else if subkind, ok := token.IsCssSelector(literal); ok {
				l.pushMode(mode_style_block)
				return token.NewSelector(literal, subkind)
			} else {
				return token.NewIllegal(literal)
			}

		default:
			defer l.consume()
			return token.NewIllegal(string(l.char))
		}

	case mode_media_or_supports_condition:
		l.skipWhitespace()

		switch {
		case l.char == 0:
			defer l.consume()
			return token.NewEOF()

		case l.char == '(':
			defer l.consume()
			return token.NewOpenParenthesis()

		case l.char == ')':
			defer l.consume()
			return token.NewCloseParenthesis()

		case l.char == ':':
			if !unicode.IsSpace(l.peek()) {
				literal := l.consumeSelector()
				if subkind, ok := token.IsCssSelector(literal); ok {
					l.pushMode(mode_style_block)
					return token.NewSelector(literal, subkind)
				} else {
					return token.NewIllegal(literal)
				}
			} else {
				defer l.consume()
				return token.NewColon()
			}

		case l.char == '{':
			l.popMode()
			l.pushMode(mode_media_or_supports_block)
			defer l.consume()
			return token.NewOpenCurlyBracket()

		default:
			if unicode.IsLetter(l.char) {
				literal := l.consumeIdentifier()
				return token.NewIdentifier(literal)
			} else if unicode.IsDigit(l.char) || l.char == '.' {
				subkind, literal := l.consumeNumeric()
				return token.NewValue(literal, subkind)
			} else {
				defer l.consume()
				return token.NewIllegal(string(l.char))
			}
		}

	case mode_media_or_supports_block:
		l.skipWhitespace()

		switch l.char {

		case 0:
			defer l.consume()
			return token.NewEOF()

		case '{':
			defer l.consume()
			return token.NewOpenCurlyBracket()

		case '}':
			defer l.consume()
			l.popMode()
			return token.NewCloseCurlyBracket()

		case '(':
			defer l.consume()
			return token.NewOpenParenthesis()

		case ')':
			defer l.consume()
			return token.NewCloseParenthesis()

		case ':':
			if !unicode.IsSpace(l.peek()) {
				literal := l.consumeSelector()
				if subkind, ok := token.IsCssSelector(literal); ok {
					l.pushMode(mode_style_block)
					return token.NewSelector(literal, subkind)
				} else {
					return token.NewIllegal(literal)
				}
			} else {
				defer l.consume()
				return token.NewColon()
			}

		default:
			literal := l.consumeSelector()
			if subkind, ok := token.IsCssSelector(literal); ok {
				l.pushMode(mode_style_block)
				return token.NewSelector(literal, subkind)
			} else {
				return token.NewIllegal(literal)
			}
		}

	case mode_value:
		l.skipWhitespace()

		switch l.char {

		case 0:
			defer l.consume()
			return token.NewEOF()

		case ',':
			defer l.consume()
			return token.NewComma()

		case ':':
			defer l.consume()
			return token.NewColon()

		case ';':
			defer l.consume()
			l.popMode()
			return token.NewSemicolon()

		case '(':
			defer l.consume()
			return token.NewOpenParenthesis()

		case ')':
			defer l.consume()
			return token.NewCloseParenthesis()

		case '}':
			defer l.consume()
			l.popMode()
			return token.NewCloseCurlyBracket()

		default:
			if unicode.IsLetter(l.char) {
				literal := l.consumeValue()
				if subkind, ok := token.IsCssValue(literal); ok {
					return token.NewValue(literal, subkind)
				} else {
					return token.NewIllegal(literal)
				}
			} else if unicode.IsDigit(l.char) || l.char == '.' {
				subkind, literal := l.consumeNumeric()
				return token.NewValue(literal, subkind)
			} else {
				return token.NewIllegal(string(l.char))
			}
		}

	default:
		defer l.consume()
		return token.NewIllegal(string(l.char))
	}
}

func (l *Lexer) currMode() (m mode) {
	if len(l.modeStack) > 0 {
		m = l.modeStack[len(l.modeStack)-1]
	} else {
		m = mode_top_level
	}
	// if len(l.modeStack) > 0 {
	// 	for _, n := range l.modeStack {
	// 		fmt.Println(n)
	// 	}
	// } else {
	// 	fmt.Println(mode_top_level)
	// }
	return
}

func (l *Lexer) pushMode(mode mode) {
	l.modeStack = append(l.modeStack, mode)
}

func (l *Lexer) popMode() {
	if len(l.modeStack) > 0 {
		l.modeStack = l.modeStack[:len(l.modeStack)-1]
	}
}

func (l *Lexer) consume() {
	if l.cursor >= len(l.input) {
		l.char = 0
	} else {
		l.char = rune(l.input[l.cursor])
	}
	l.position = l.cursor
	l.cursor += 1
}

func (l *Lexer) consumeIdentifier() string {
	position := l.position
	if l.char == '&' {
		l.consume()
	}
	if l.char == ':' {
		l.consume()
		if l.char == ':' {
			l.consume()
		}
	}
	for unicode.IsLetter(l.char) || l.char == '-' || l.char == '_' {
		l.consume()
	}
	return l.input[position:l.position]
}

func (l *Lexer) consumeSelector() string {
	position := l.position
	for {
		if l.peek() == ',' || l.peek() == '{' {
			break
		}
		l.consume()
	}
	return l.input[position:l.position]
}

func (l *Lexer) consumeAtKeyword() string {
	position := l.position
	for unicode.IsLetter(l.char) || l.char == '-' {
		l.consume()
	}
	return l.input[position:l.position]
}

func (l *Lexer) consumeValue() string {
	position := l.position
	for {
		if unicode.IsSpace(l.char) || l.char == ';' {
			break
		}
		l.consume()
	}
	return l.input[position:l.position]
}

func (l *Lexer) consumeNumeric() (kind token.Subkind, literal string) {
	position := l.position
	seenDecimal := false
	kind = token.INTEGER_LITERAL_VALUE
loop:
	for {
		switch {
		case unicode.IsSpace(l.char),
			l.char == ',',
			l.char == ')',
			l.char == ';',
			l.char == '-',
			l.char == '+',
			l.char == '*',
			l.char == '/':
			break loop
		case l.char == '%':
			kind = token.PERCENTAGE_VALUE
			break loop
		case l.char == '.':
			kind = token.FLOAT_LITERAL_VALUE
			if !seenDecimal {
				seenDecimal = true
			} else {
				break loop
			}
		case unicode.IsLetter(l.char):
			kind = token.DIMENSION_VALUE
			for unicode.IsLetter(l.char) {
				l.consume()
			}
			break loop
		}
		l.consume()
	}
	return kind, l.input[position:l.position]
}

func (l *Lexer) consumeNumberLiteral() []rune {
	position := l.position
	seenDecimal := false
	for unicode.IsDigit(l.char) || l.char == '.' {
		if l.char == '.' {
			if seenDecimal {
				break
			} else {
				seenDecimal = true
			}
		}
		l.consume()
	}
	return []rune(l.input[position:l.position])
}

func (l *Lexer) peek() rune {
	return l.peekN(1)
}

func (l *Lexer) peekN(n int) rune {
	if l.cursor-1+n >= len(l.input) {
		return 0
	}
	return rune(l.input[l.cursor-1+n])
}

func (l *Lexer) peekString(n int) string {
	if l.cursor-1+n >= len(l.input) {
		return ""
	} else {
		return string(l.input[l.cursor : l.cursor-1+n])
	}
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.char) {
		l.consume()
	}
}

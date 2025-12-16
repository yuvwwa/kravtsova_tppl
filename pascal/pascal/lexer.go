package pascal

import (
	"fmt"
	"strings"
	"unicode"
)

type Lexer struct {
	text        string
	pos         int
	currentChar rune
}

func NewLexer(text string) *Lexer {
	l := &Lexer{
		text: text,
		pos:  0,
	}
	if len(text) > 0 {
		l.currentChar = rune(text[0])
	} else {
		l.currentChar = 0
	}
	return l
}

func (l *Lexer) advance() {
	l.pos++
	if l.pos >= len(l.text) {
		l.currentChar = 0
	} else {
		l.currentChar = rune(l.text[l.pos])
	}
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar != 0 && unicode.IsSpace(l.currentChar) {
		l.advance()
	}
}

func (l *Lexer) integer() string {
	var result strings.Builder
	for l.currentChar != 0 && unicode.IsDigit(l.currentChar) {
		result.WriteRune(l.currentChar)
		l.advance()
	}
	return result.String()
}

func (l *Lexer) id() string {
	var result strings.Builder
	for l.currentChar != 0 && (unicode.IsLetter(l.currentChar) || unicode.IsDigit(l.currentChar) || l.currentChar == '_') {
		result.WriteRune(l.currentChar)
		l.advance()
	}
	return result.String()
}

func (l *Lexer) peek() rune {
	peekPos := l.pos + 1
	if peekPos >= len(l.text) {
		return 0
	}
	return rune(l.text[peekPos])
}

func (l *Lexer) NextToken() (Token, error) {
	for l.currentChar != 0 {
		if unicode.IsSpace(l.currentChar) {
			l.skipWhitespace()
			continue
		}

		if unicode.IsDigit(l.currentChar) {
			return NewToken(INTEGER, l.integer()), nil
		}

		if unicode.IsLetter(l.currentChar) || l.currentChar == '_' {
			idStr := l.id()
			idUpper := strings.ToUpper(idStr)
			if idUpper == "BEGIN" {
				return NewToken(BEGIN, idStr), nil
			}
			if idUpper == "END" {
				return NewToken(END, idStr), nil
			}
			return NewToken(ID, idStr), nil
		}

		if l.currentChar == ':' && l.peek() == '=' {
			l.advance()
			l.advance()
			return NewToken(ASSIGN, ":="), nil
		}

		switch l.currentChar {
		case '+':
			l.advance()
			return NewToken(PLUS, "+"), nil
		case '-':
			l.advance()
			return NewToken(MINUS, "-"), nil
		case '*':
			l.advance()
			return NewToken(MUL, "*"), nil
		case '/':
			l.advance()
			return NewToken(DIV, "/"), nil
		case '(':
			l.advance()
			return NewToken(LPAREN, "("), nil
		case ')':
			l.advance()
			return NewToken(RPAREN, ")"), nil
		case ';':
			l.advance()
			return NewToken(SEMI, ";"), nil
		case '.':
			l.advance()
			return NewToken(DOT, "."), nil
		default:
			return Token{}, fmt.Errorf("invalid character: %c", l.currentChar)
		}
	}

	return NewToken(EOF, ""), nil
}
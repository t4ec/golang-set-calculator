package lexer

import (
	"solution/token"
)

type Lexer struct {
	text  []rune
	char  rune
	start int
	end   int
}

func NewLexer(text string) Lexer {
	return Lexer{text: []rune(text)}
}

func (l *Lexer) readChar() {
	if l.end >= len(l.text) {
		l.char = 0
	} else {
		l.char = l.text[l.end]
		l.start = l.end
		l.end += 1
	}
}

func (l *Lexer) skipWhitespaces() {
	for {
		switch l.char {
		case ' ', '\n', '\t', '\r':
			l.readChar()
		default:
			return
		}
	}
}

func (l *Lexer) NextToken() token.Token {
	l.readChar()
	l.skipWhitespaces()
	switch l.char {
	case '[':
		return token.Token{Type: token.LBRACKET, Value: string(l.char)}
	case ']':
		return token.Token{Type: token.RBRACKET, Value: string(l.char)}
	default:
		if l.char == rune(0) {
			return token.Token{Type: token.EOF, Value: ""}
		} else {
			return token.Token{Type: token.WORD, Value: l.readWord()}
		}
	}
}

func (l *Lexer) readWord() string {
	chars := []rune{}
	for {
		switch l.char {
		case ' ', '\n', '\t', '[', ']':
			return string(chars)
		default:
			chars = append(chars, l.char)
			l.readChar()
		}
	}
}

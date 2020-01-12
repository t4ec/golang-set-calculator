package lexer

import (
	"solution/token"
	"testing"
)

type testpair struct {
	value  string
	tokens []token.Token
}

var tests = []testpair{
	{
		"[ SUM [ DIF a.txt b.txt c.txt ] ]",
		[]token.Token{
			token.Token{Value: "[", Type: token.LBRACKET},
			token.Token{Value: "SUM", Type: token.WORD},
			token.Token{Value: "[", Type: token.LBRACKET},
			token.Token{Value: "DIF", Type: token.WORD},
			token.Token{Value: "a.txt", Type: token.WORD},
			token.Token{Value: "b.txt", Type: token.WORD},
			token.Token{Value: "c.txt", Type: token.WORD},
			token.Token{Value: "]", Type: token.RBRACKET},
			token.Token{Value: "]", Type: token.RBRACKET},
		},
	},
}

func TestLexer(t *testing.T) {
	for _, test := range tests {
		l := NewLexer(test.value)
		var tokens []token.Token

		equal := true
		for _, expectedToken := range test.tokens {
			tok := l.NextToken()
			if tok != expectedToken {
				equal = false
			}
			tokens = append(tokens, tok)
		}
		if !equal {
			t.Errorf("For string \"%s\"\nexpected tokens:\t\"%v\",\ngot:\t\t\t\"%v\"", test.value, test.tokens, tokens)
		}
	}
}

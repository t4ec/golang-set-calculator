package parser

import (
	"fmt"

	"solution/lexer"
	"solution/token"
)

type Parser struct {
	Lexer lexer.Lexer
}

func NewParser(text string) Parser {
	lexer := lexer.NewLexer(text)
	return Parser{Lexer: lexer}
}

func (p *Parser) Parse() {
	for {
		t := p.Lexer.NextToken()
		switch t.Type {
		case token.LBRACKET:
			p.parseExpression()
		}
	}
}

func (p *Parser) parseExpression() {
	t := p.Lexer.NextToken()
	if t.Type == token.WORD {
		fmt.Printf("Operator: %v", t.Value)
	}
}

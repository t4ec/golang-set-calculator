package parser

import (
	"fmt"

	"solution/ast"
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

func (p *Parser) Parse() (ast.ASTNode, error) {
	astNode, err := p.parseOperator()
	return astNode, err
}

func (p *Parser) parseOperator() (ast.ASTNode, error) {
	t := p.Lexer.NextToken()
	// fmt.Printf("parseOperator, token value: \"%v\", token type: \"%v\"\n", t.Value, t.Type)
	switch t.Type {
	case token.LBRACKET:
		astNode, err := p.parseOperator()
		if err != nil {
			return ast.ASTOperatorNode{}, fmt.Errorf("failed parsing Operator: %v", err)
		}
		return astNode, nil
	case token.WORD:
		switch t.Value {
		case ast.DIFFERENCE, ast.INTERSECTION, ast.UNION:
			childred, err := p.parseOperands()
			if err != nil {
				return ast.ASTOperatorNode{}, fmt.Errorf("invalid operands for operator %v:\n%v", t.Value, err)
			}
			return ast.ASTOperatorNode{Operation: t.Value, Children: childred}, nil
		default:
			return ast.ASTOperatorNode{}, fmt.Errorf("invalid operator: %v", t.Value)
		}
	}
	return ast.ASTOperatorNode{}, fmt.Errorf("unexpected formatting for token.Value = \"%v\", token.Type = \"%v\"", t.Value, t.Type)
}

func (p *Parser) parseOperands() ([]ast.ASTNode, error) {
	var operands []ast.ASTNode
	for {
		t := p.Lexer.NextToken()
		// fmt.Printf("parseOperands, token: \"%v\", token type: \"%v\"\n", t.Value, t.Type)
		switch t.Type {
		case token.LBRACKET:
			astNode, err := p.parseOperator()
			if err != nil {
				return []ast.ASTNode{}, fmt.Errorf("invalid operator formatting:\n%v", t.Value)
			}
			operands = append(operands, astNode)
		case token.WORD:
			astNode := ast.ASTValueNode{FileName: t.Value}
			operands = append(operands, astNode)
		case token.RBRACKET:
			return operands, nil
		default:
			return []ast.ASTNode{}, fmt.Errorf("invalid operands formatting:\n%v", t.Value)
		}
	}
}

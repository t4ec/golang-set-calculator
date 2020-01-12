package main

import (
	"fmt"
	"solution/parser"
)

const testString = "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]"

func main() {
	p := parser.NewParser(testString)
	ast, err := p.Parse()
	if err != nil {
		fmt.Printf("failed parsing expression:\n%v", err)
		return
	}
	ast.Evaluate()
	// fmt.Printf("AST ROOT: %v", ast.Evaluate())
}

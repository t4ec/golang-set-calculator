package main

import (
	"fmt"
	"os"
	"strings"

	"solution/parser"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("please provide expression parameter, example: [ SUM a.txt b.txt ]"))
	}
	expressionString := os.Args[1:]
	expression := strings.Join(expressionString, " ")

	p := parser.NewParser(expression)
	ast, err := p.Parse()
	if err != nil {
		fmt.Printf("failed parsing expression:\n%v", err)
		return
	}
	result, err := ast.Evaluate()
	if err != nil {
		fmt.Printf("failed evaluating AST:\n%v", err)
		return
	}
	for _, element := range result {
		fmt.Println(element)
	}
}

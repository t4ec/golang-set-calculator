package main

import (
	"fmt"
	"solution/parser"
	// "strings"
)

const testString = "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]"

type Expression struct {
	name     string
	operands Operands
}

type Operands []Expression

func parseExpression(expression string) {
	buffer := ""
	words := []string{}
	var currentExpression *Expression = nil

	for _, char := range expression {
		switch char {
		case ' ':
			if buffer != "" {
				words = append(words, buffer)
				buffer = ""
			}
			// continue
		case '[':
			if buffer != "" {
				if currentExpression == nil {
					*currentExpression = Expression{name: buffer, operands: []Expression{}}
				}
				words = append(words, buffer)
				buffer = ""
			}
			// fmt.Printf("%c", char)
		case ']':
			if buffer != "" {
				words = append(words, buffer)
				buffer = ""
			}
			// fmt.Printf("%c", char)
		default:
			buffer = buffer + string(char)
			// fmt.Printf("%c", char)
		}
	}
	fmt.Printf("%s", words[0:2])
}

func main() {
	parser.NewParser(testString)
	parser.Parse()
}

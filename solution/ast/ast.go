package ast

import (
	"fmt"
)

const (
	// Operation types
	UNION        = "SUM"
	DIFFERENCE   = "DIF"
	INTERSECTION = "INT"
)

type ASTNode interface {
	Evaluate() []int
}

type ASTOperatorNode struct {
	Operation string
	Children  []ASTNode
}

func (n ASTOperatorNode) Evaluate() []int {
	fmt.Printf("Evaluate operation %v\n", n.Operation)
	switch n.Operation {
	case UNION: // TODO
		result := []int{}
		for _, child := range n.Children {
			result = append(result, child.Evaluate()...)
		}
		return result
	case INTERSECTION: // TODO
		result := []int{}
		for _, child := range n.Children {
			result = append(result, child.Evaluate()...)
		}
		return result
	case DIFFERENCE: // TODO
		result := []int{}
		for _, child := range n.Children {
			result = append(result, child.Evaluate()...)
		}
		return result
	default: // TODO error handling
		return []int{}
	}
}

type ASTValueNode struct {
	Filepath string
	Value    []int
}

func (n ASTValueNode) Evaluate() []int {
	fmt.Printf("Evaluate value %v\n", n.Value)
	return n.Value
}

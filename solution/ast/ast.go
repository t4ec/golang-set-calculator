package ast

import (
	"solution/math"
)

// Operation types
const (
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
	switch n.Operation {
	case UNION: // TODO
		childrenSets := [][]int{}
		for _, child := range n.Children {
			childrenSets = append(childrenSets, child.Evaluate())
		}
		return math.Union(childrenSets)
	case INTERSECTION: // TODO
		childrenSets := [][]int{}
		for _, child := range n.Children {
			childrenSets = append(childrenSets, child.Evaluate())
		}
		return math.Intersection(childrenSets)
	case DIFFERENCE: // TODO
		childrenSets := [][]int{}
		for _, child := range n.Children {
			childrenSets = append(childrenSets, child.Evaluate())
		}
		return math.Difference(childrenSets)
	default: // TODO error handling
		return []int{}
	}
}

type ASTValueNode struct {
	Filepath string
	Value    []int
}

func (n ASTValueNode) Evaluate() []int {
	return n.Value
}

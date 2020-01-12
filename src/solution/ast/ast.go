package ast

import (
	"fmt"
	"solution/math"
)

// Operation types
const (
	UNION        = "SUM"
	DIFFERENCE   = "DIF"
	INTERSECTION = "INT"
)

type ASTNode interface {
	Evaluate() ([]int, error)
}

type ASTOperatorNode struct {
	Operation string
	Children  []ASTNode
}

func (n ASTOperatorNode) Evaluate() ([]int, error) {
	switch n.Operation {
	case UNION:
		childrenSets := [][]int{}
		for _, child := range n.Children {
			childValue, err := child.Evaluate()
			if err != nil {
				return []int{}, fmt.Errorf("could not evaluate operation %v:\n%v", n.Operation, err)
			}
			childrenSets = append(childrenSets, childValue)
		}
		return math.Union(childrenSets), nil
	case INTERSECTION:
		childrenSets := [][]int{}
		for _, child := range n.Children {
			childValue, err := child.Evaluate()
			if err != nil {
				return []int{}, fmt.Errorf("could not evaluate operation %v:\n%v", n.Operation, err)
			}
			childrenSets = append(childrenSets, childValue)
		}
		return math.Intersection(childrenSets), nil
	case DIFFERENCE:
		childrenSets := [][]int{}
		for _, child := range n.Children {
			childValue, err := child.Evaluate()
			if err != nil {
				return []int{}, fmt.Errorf("could not evaluate operation %v:\n%v", n.Operation, err)
			}
			childrenSets = append(childrenSets, childValue)
		}
		return math.Difference(childrenSets), nil
	default:
		return []int{}, fmt.Errorf("invalid operation: %v", n.Operation)
	}
}

type ASTValueNode struct {
	Filepath string
	Value    []int
}

func (n ASTValueNode) Evaluate() ([]int, error) {
	return n.Value, nil
}

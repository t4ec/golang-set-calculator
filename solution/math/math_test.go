package math

import (
	"testing"
)

type testpair struct {
	sets   [][]int
	result []int
}

var unionTests = []testpair{
	{
		sets:   [][]int{[]int{1, 2, 3}, []int{2, 3, 4}, []int{6}},
		result: []int{1, 2, 3, 4, 6},
	},
	{
		sets:   [][]int{[]int{}, []int{3, 4}, []int{1, 2, 3}},
		result: []int{1, 2, 3, 4},
	},
	{
		sets:   [][]int{[]int{6}},
		result: []int{6},
	},
	{
		sets:   [][]int{[]int{6}, []int{1, 2, 3, 4, 6, 7, 8, 9}},
		result: []int{1, 2, 3, 4, 6, 7, 8, 9},
	},
	{
		sets:   [][]int{[]int{6}, []int{1}, []int{10}},
		result: []int{1, 6, 10},
	},
}

var intersectionTests = []testpair{
	{
		sets:   [][]int{[]int{1, 2, 3, 4}, []int{2, 3, 4}},
		result: []int{2, 3, 4},
	},
	{
		sets:   [][]int{[]int{1, 2, 3, 4}, []int{2, 3, 4}, []int{3, 4, 5}},
		result: []int{3, 4},
	},
	{
		sets:   [][]int{[]int{1, 2, 3}, []int{2, 3, 4}, []int{6}},
		result: []int{},
	},
	{
		sets:   [][]int{[]int{}, []int{1, 2}},
		result: []int{},
	},
	{
		sets:   [][]int{[]int{6}},
		result: []int{6},
	},
	{
		sets:   [][]int{[]int{3, 4, 7, 10}, []int{2, 4, 5, 10}},
		result: []int{4, 10},
	},
}

var differenceTests = []testpair{
	{
		sets:   [][]int{[]int{1, 2, 3, 4}, []int{2, 3}},
		result: []int{1, 4},
	},
	{
		sets:   [][]int{[]int{1, 2, 3, 4}, []int{2, 3, 4}, []int{3, 4, 5}},
		result: []int{1},
	},
	{
		sets:   [][]int{[]int{1, 2, 3}, []int{1, 2, 3}},
		result: []int{},
	},
	{
		sets:   [][]int{[]int{}, []int{1, 2}},
		result: []int{},
	},
	{
		sets:   [][]int{[]int{6}},
		result: []int{6},
	},
}

func isEqualArrays(lhs []int, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, value := range lhs {
		if value != rhs[i] {
			return false
		}
	}
	return true
}

func TestUnion(t *testing.T) {
	for _, test := range unionTests {
		result := Union(test.sets)
		if !isEqualArrays(result, test.result) {
			t.Errorf("For %v,\nexpected: %v\ngot: %v", test.sets, test.result, result)
		}
	}
}

func TestIntersection(t *testing.T) {
	for _, test := range intersectionTests {
		result := Intersection(test.sets)
		if !isEqualArrays(result, test.result) {
			t.Errorf("For %v,\nexpected: %v\ngot: %v", test.sets, test.result, result)
		}
	}
}

func TestDifference(t *testing.T) {
	for _, test := range differenceTests {
		result := Difference(test.sets)
		if !isEqualArrays(result, test.result) {
			t.Errorf("For %v,\nexpected: %v\ngot: %v", test.sets, test.result, result)
		}
	}
}

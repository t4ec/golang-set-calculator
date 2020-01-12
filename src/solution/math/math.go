package math

import (
// "fmt"
)

func Union(sets [][]int) []int {
	if len(sets) == 1 {
		return sets[0]
	}

	result := []int{}
	for _, set := range sets {
		if len(result) == 0 {
			result = set
		}
		if len(set) == 0 {
			continue
		}

		// Keep ordered sets order so we don't have to sort them later
		lhs, rhs := result, set
		if lhs[0] > rhs[0] {
			lhs, rhs = set, result
		}
		lhsBuffer := lhs

		i := 0
		j := 0
		for {
			// If left-hand-side pointer is out of array we just append the rest of rhs
			// since they are sorted and arranged to each other. And we unite them.
			if i >= len(lhs) && j < len(rhs) {
				lhsBuffer = append(lhsBuffer, rhs[j:]...)
				break
			} else if j >= len(rhs) && i < len(lhs) {
				break
			} else if i >= len(lhs) && j >= len(rhs) {
				break
			}

			if lhs[i] < rhs[j] {
				i++
				continue
			} else if lhs[i] == rhs[j] {
				i++
				j++
				continue
			} else {
				lhsBuffer = append(lhsBuffer, rhs[j])
				i++
				j++
				continue
			}
		}
		result = lhsBuffer
	}
	return result
}

func Intersection(sets [][]int) []int {
	result := []int{}
	for setIndex, set := range sets {
		// intersection with empty set
		if len(set) == 0 {
			return []int{}
		}

		// initialize result with first set
		if setIndex == 0 {
			result = set
			continue
		}

		// Place ordered sets in order so we don't have to sort them later
		lhs, rhs := result, set
		if lhs[0] > rhs[0] {
			lhs, rhs = rhs, lhs
		}
		resultBuffer := []int{}
		for _, lhsVal := range lhs {
			if lhsVal < rhs[0] {
				continue
			}
			for _, rhsVal := range rhs {
				if lhsVal == rhsVal {
					resultBuffer = append(resultBuffer, rhsVal)
				}
			}
		}
		result = resultBuffer
	}
	return result
}

func Difference(sets [][]int) []int {
	result := []int{}
	for setIndex, set := range sets {
		// intersection with empty set
		if len(set) == 0 {
			return []int{}
		}

		// initialize result with first set
		if setIndex == 0 {
			result = set
			continue
		}

		// Place ordered sets in order so we don't have to sort them later
		lhs, rhs := result, set
		if lhs[0] > rhs[0] {
			lhs, rhs = rhs, lhs
		}
		resultBuffer := []int{}
		for _, lhsVal := range lhs {
			if lhsVal < rhs[0] {
				resultBuffer = append(resultBuffer, lhsVal)
			}
			if lhsVal > rhs[len(rhs)-1] {
				resultBuffer = append(resultBuffer, lhsVal)
			}
			for _, rhsVal := range rhs {
				if lhsVal == rhsVal {
					continue
				}
			}
		}
		result = resultBuffer
	}
	return result
}

// Copyright 2020, Dmytrii Sinko, All rights reserved.

package math

// Order sets: first element of 1st set should be <= then first element of 2nd set.
// This allows us to keep order after we do the math without additional sorting.
func alignOrderedSets(lhs []int, rhs []int) ([]int, []int) {
	if lhs[0] > rhs[0] {
		return rhs, lhs
	}
	return lhs, rhs
}

func Union(sets [][]int) []int {
	if len(sets) == 1 {
		return sets[0]
	}

	result := []int{}
	for _, set := range sets {
		if len(result) == 0 {
			result = set
			continue
		}
		if len(set) == 0 {
			continue
		}

		lhs, rhs := alignOrderedSets(result, set)

		i := 0
		j := 0
		buffer := []int{}
		for {
			if i >= len(lhs) || j >= len(rhs) {
				break
			}

			if lhs[i] < rhs[j] {
				buffer = append(buffer, lhs[i])
				i++
			} else if lhs[i] == rhs[j] {
				buffer = append(buffer, lhs[i])
				i++
				j++
			} else {
				buffer = append(buffer, rhs[j])
				j++
			}
		}

		for {
			if i >= len(lhs) {
				break
			}
			buffer = append(buffer, lhs[i])
			i++
		}

		for {
			if j >= len(rhs) {
				break
			}
			buffer = append(buffer, rhs[j])
			j++
		}
		result = buffer
	}
	return result
}

func Intersection(sets [][]int) []int {
	if len(sets) == 1 {
		return sets[0]
	}

	result := []int{}
	for setIn, set := range sets {
		if len(set) == 0 {
			return []int{}
		}

		if setIn == 0 {
			result = set
			continue
		}

		lhs, rhs := alignOrderedSets(result, set)

		i := 0
		j := 0
		buffer := []int{}
		for {
			if i >= len(lhs) || j >= len(rhs) {
				break
			}

			if lhs[i] < rhs[j] {
				i++
			} else if lhs[i] == rhs[j] {
				buffer = append(buffer, lhs[i])
				i++
				j++
			} else {
				j++
			}
		}
		result = buffer
	}
	return result
}

func Difference(sets [][]int) []int {
	if len(sets) == 1 {
		return sets[0]
	}

	result := []int{}
	for setIn, set := range sets {
		if setIn == 0 {
			result = set
			if len(result) == 0 {
				return []int{}
			}
			continue
		}
		if len(set) == 0 {
			continue
		}

		lhs, rhs := result, set

		i := 0
		j := 0
		buffer := []int{}
		for {
			if i >= len(lhs) || j >= len(rhs) {
				break
			}

			if lhs[i] < rhs[j] {
				buffer = append(buffer, lhs[i])
				i++
			} else if lhs[i] == rhs[j] {
				i++
				j++
			} else {
				j++
			}
		}

		for {
			if i >= len(lhs) {
				break
			}
			buffer = append(buffer, lhs[i])
			i++
		}

		result = buffer
	}
	return result
}

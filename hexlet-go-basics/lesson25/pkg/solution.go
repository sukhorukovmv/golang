package pkg

// BEGIN (write your solution here)
func MergeNumberLists(numberLists ...[]int) []int {
	res := make([]int, 0, len(numberLists))
	for _, list := range numberLists {
		res = append(res, list...)
	}
	return res
}

// END

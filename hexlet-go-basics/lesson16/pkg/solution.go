package pkg

// BEGIN (write your solution here)

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	}
	if maxLen > len(src) {
		maxLen = len(src)
	}
	res := make([]int, maxLen)
	copy(res, src)
	return res
}

// END

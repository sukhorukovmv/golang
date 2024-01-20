package pkg

// BEGIN (write your solution here)

func sum(arr []int) int {
	s := 0
	for _, a := range arr {
		s += a
	}
	return s
}

func MaxSum(sl1 []int, sl2 []int) []int {
	if sum(sl1) >= sum(sl2) {
		return sl1
	}
	return sl2
}

// END

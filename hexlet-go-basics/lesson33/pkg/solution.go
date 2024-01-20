package pkg

import (
	"time"
)

// BEGIN (write your solution here)

func sumParallel(arr []int, res *int) {
    *res = sum(arr)
}

func sum(arr []int) int {
	s := 0
	for _, a := range arr {
		s += a
	}
	return s
}

func MaxSum(sl1 []int, sl2 []int) []int {
    s1, s2 := 0, 0
    go sumParallel(sl1, &s1)
    go sumParallel(sl2, &s2)
    time.Sleep(100 * time.Millisecond)

	if s1 >= s2 {
		return sl1
	}
	return sl2
}

// END

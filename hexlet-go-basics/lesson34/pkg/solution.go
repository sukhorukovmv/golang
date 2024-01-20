package pkg

// BEGIN (write your solution here)

func SumWorker(numsCh chan []int, sumCh chan int) {
    for nums := range numsCh {
        sumCh <- sum(nums)
    }
}

func sum(arr []int) int {
	s := 0
	for _, a := range arr {
		s += a
	}
	return s
}

// END

package pkg

// BEGIN (write your solution here)

func Remove(nums []int, i int) []int {
	if i >= len(nums) || i < 0 {
		return nums
	}
	return append(nums[:i], nums[i+1:]...)
}

// END

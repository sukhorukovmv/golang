package pkg

// BEGIN (write your solution here)
func UniqueUserIDs(userIDs []int64) []int64 {
	uniq_elemets := make(map[int64]int, len(userIDs))
	res := make([]int64, 0, len(userIDs))
	for i, id := range userIDs {
		_, key_exist := uniq_elemets[id]
		if !key_exist {
			uniq_elemets[id] = i
			res = append(res, id)
		}
	}

	return res
}

// END

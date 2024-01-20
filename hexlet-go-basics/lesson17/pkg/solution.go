package pkg

import (
	"sort"
)

// BEGIN (write your solution here)

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	sort.Slice(userIDs, func(i, j int) bool {
		if userIDs[i] == userIDs[j] {
			userIDs[i] = -1
		}
		return userIDs[i] < userIDs[j]
	})

	equalIds := 0
	for _, id := range userIDs {
		if id == -1 {
			equalIds++
		}
	}
	return userIDs[equalIds:]
}

func UniqueSortedUserIDs2(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}

// END

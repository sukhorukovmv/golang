package pkg

type Person struct {
	Age uint8
}

// BEGIN (write your solution here)
type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	popAges := make(map[uint8]int, 0)
	for _, p := range pl {
		popAges[p.Age]++
	}
	return popAges
}

// END

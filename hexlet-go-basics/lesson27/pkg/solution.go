package pkg

type Counter struct {
	Value int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// BEGIN (write your solution here)
func (c *Counter) Inc(delta int) {
	if delta == 0 {
		c.Value += 1
	}
	c.Value += delta
	println(c.Value)
}

func (c *Counter) Dec(delta int) {
	if delta == 0 {
		c.Value = Max(0, c.Value-1)
	}
	c.Value = Max(0, c.Value-delta)
}

func (c *Counter) Inc2(delta int) {
	if delta == 0 {
		delta = 1
	}
	c.Value = Max(c.Value+delta, 0)
}

func (c *Counter) Dec2(delta int) {
	if delta == 0 {
		delta = 1
	}
	c.Inc2(-delta)
}

// END

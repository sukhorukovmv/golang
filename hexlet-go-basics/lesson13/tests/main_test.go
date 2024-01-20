package test

import (
	"testing"

	"package/pkg"

	"github.com/stretchr/testify/assert"
)

func TestSafeWrite(t *testing.T) {
	a := assert.New(t)
	a.Equal([5]int{}, pkg.SafeWrite([5]int{}, -1, 10))
	a.Equal([5]int{}, pkg.SafeWrite([5]int{}, -10, 10))
	a.Equal([5]int{}, pkg.SafeWrite([5]int{}, 10, 10))
	a.Equal([5]int{}, pkg.SafeWrite([5]int{}, 20, 10))
	a.Equal([5]int{1, 2, 10, 4, 5}, pkg.SafeWrite([5]int{1, 2, 3, 4, 5}, 2, 10))
	a.Equal([5]int{1, 2, 3, 4, 22}, pkg.SafeWrite([5]int{1, 2, 3, 4, 5}, 4, 22))
}

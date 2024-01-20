package solution

import (
	"github.com/stretchr/testify/assert"
	"package/pkg"
	"sort"
	"testing"
)

func TestRemove(t *testing.T) {
	a := assert.New(t)
	equalSlices(a, []int{1, 2}, pkg.Remove([]int{1, 2}, -1))
	equalSlices(a, []int{1, 2}, pkg.Remove([]int{1, 2}, -5))
	equalSlices(a, []int{2, 3}, pkg.Remove([]int{1, 2, 3}, 0))
	equalSlices(a, []int{1, 2}, pkg.Remove([]int{1, 2, 3}, 2))
	equalSlices(a, []int{1, 3}, pkg.Remove([]int{1, 2, 3}, 1))
	equalSlices(a, []int{1, 2, 3}, pkg.Remove([]int{1, 2, 3}, 3))
	equalSlices(a, []int{1, 2, 3}, pkg.Remove([]int{1, 2, 3}, 5))
}

func equalSlices(a *assert.Assertions, expected, actual []int) {
	sort.Ints(expected)
	sort.Ints(actual)
	a.Equal(expected, actual)
}

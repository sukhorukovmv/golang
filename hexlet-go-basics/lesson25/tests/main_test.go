package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"package/pkg"
)

func TestMergeNumberLists(t *testing.T) {
	a := assert.New(t)
	a.Equal(
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		pkg.MergeNumberLists([]int{1, 2, 3}, []int{4}, []int{5, 6, 7, 8}, []int{9, 10}),
	)
	a.Equal([]int{}, pkg.MergeNumberLists([]int{}))
	a.Equal([]int{}, pkg.MergeNumberLists(nil, nil))
	a.Equal([]int{10, 20, 50, 60}, pkg.MergeNumberLists([]int{10, 20}, nil, []int{50, 60}))
}

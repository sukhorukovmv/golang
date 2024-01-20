package solution

import (
	"package/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{}, pkg.IntsCopy([]int{}, 0))
	a.Equal([]int{}, pkg.IntsCopy([]int{1, 2}, 0))
	a.Equal([]int{}, pkg.IntsCopy([]int{1, 2}, -1))
	a.Equal([]int{}, pkg.IntsCopy([]int{1, 2}, -5))
	a.Equal([]int{1, 2}, pkg.IntsCopy([]int{1, 2, 3}, 2))
	a.Equal([]int{1, 2, 3, 4}, pkg.IntsCopy([]int{1, 2, 3, 4}, 4))
	a.Equal([]int{1, 2, 3, 4, 5}, pkg.IntsCopy([]int{1, 2, 3, 4, 5}, 10))
}

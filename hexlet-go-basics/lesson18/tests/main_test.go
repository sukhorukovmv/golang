package solution

import (
	"github.com/stretchr/testify/assert"
	"package/pkg"
	"testing"
)

func TestUniqueUserIDs(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int64{}, pkg.UniqueUserIDs([]int64{}))
	a.Equal([]int64{10}, pkg.UniqueUserIDs([]int64{10}))
	a.Equal([]int64{55}, pkg.UniqueUserIDs([]int64{55, 55}))
	a.Equal([]int64{55, 33, 22}, pkg.UniqueUserIDs([]int64{55, 55, 33, 22}))
	a.Equal([]int64{55, 2, 88, 33, 103}, pkg.UniqueUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
}

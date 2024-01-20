package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"package/pkg"
)

func TestCopyParent(t *testing.T) {
	a := assert.New(t)

	// nil case
	cp := pkg.CopyParent(nil)
	a.Equal(pkg.Parent{}, cp)

	// filled struct case
	p := &pkg.Parent{
		Name: "Harry",
		Children: []pkg.Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}

	cp = pkg.CopyParent(p)

	a.Equal("Harry", cp.Name)
	a.Equal("Harry", p.Name)

	a.Len(p.Children, 1)
	a.Len(cp.Children, 1)

	// mutate copy
	origin_child := []pkg.Child{
		{
			Name: "Andy",
			Age:  18,
		},
	}
	a.Equal(cp.Children, origin_child)
	a.Equal(p.Children, origin_child)

	cp.Children[0] = pkg.Child{}
	a.Equal(p.Children, origin_child)
	a.Equal([]pkg.Child{
		{},
	}, cp.Children)
}

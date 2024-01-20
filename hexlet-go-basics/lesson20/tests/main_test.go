package solution

import (
	"testing"

	"package/pkg"

	"github.com/stretchr/testify/assert"
)

func TestNextASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal(byte('b'), pkg.NextASCII(byte('a')))
	a.Equal(byte('y'), pkg.NextASCII(byte('x')))
	a.Equal(byte('6'), pkg.NextASCII(byte('5')))
}

func TestPrevASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal(byte('7'), pkg.PrevASCII(byte('8')))
	a.Equal(byte('c'), pkg.PrevASCII(byte('d')))
	a.Equal(byte('m'), pkg.PrevASCII(byte('n')))
}

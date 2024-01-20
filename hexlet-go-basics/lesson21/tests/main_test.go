package solution

import (
	"package/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal("abc", pkg.ShiftASCII("abc", 0))
	a.Equal("bcd2", pkg.ShiftASCII("abc1", 1))
	a.Equal("abc1", pkg.ShiftASCII("bcd2", -1))
	a.Equal("rs", pkg.ShiftASCII("hi", 10))
	a.Equal("abc", pkg.ShiftASCII("abc", 256))
	a.Equal("bcd", pkg.ShiftASCII("abc", -511))
}

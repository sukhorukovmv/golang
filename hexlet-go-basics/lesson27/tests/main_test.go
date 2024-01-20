package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"package/pkg"
)

func TestIncDec(t *testing.T) {
	a := assert.New(t)

	c := pkg.Counter{}
	c.Inc(0)
	c.Inc(0)
	c.Inc(40)
	a.Equal(c.Value, 42)
	c.Dec(0)
	c.Dec(30)
	a.Equal(c.Value, 11)
	c.Dec(100)
	// Counter value cannot be negative
	a.Equal(c.Value, 0)
}

func TestCounterIndependence(t *testing.T) {
	a := assert.New(t)

	c1 := pkg.Counter{}
	c2 := pkg.Counter{}

	c1.Inc(10)
	c2.Inc(20)

	a.Equal(c1.Value, 10)
	a.Equal(c2.Value, 20)
}

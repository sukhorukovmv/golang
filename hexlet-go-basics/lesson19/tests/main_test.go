package solution

import (
	"testing"

	"package/pkg"

	"github.com/stretchr/testify/assert"
)

func TestMostPopularWord(t *testing.T) {
	a := assert.New(t)
	a.Equal("hello", pkg.MostPopularWord([]string{"hello", "world", "hello"}))
	a.Equal("world", pkg.MostPopularWord([]string{"hello", "world", "hello", "world", "world"}))
	a.Equal("one", pkg.MostPopularWord([]string{"one", "two", "three", "four", "five"}))
	a.Equal("c", pkg.MostPopularWord([]string{"a", "b", "c", "c", "d", "e", "e", "d"}))
	a.Equal("a", pkg.MostPopularWord([]string{"a", "c", "c", "a"}))
}

func TestMostPopularWord2(t *testing.T) {
	a := assert.New(t)
	a.Equal("hello", pkg.MostPopularWord2([]string{"hello", "world", "hello"}))
	a.Equal("world", pkg.MostPopularWord2([]string{"hello", "world", "hello", "world", "world"}))
	a.Equal("one", pkg.MostPopularWord2([]string{"one", "two", "three", "four", "five"}))
	a.Equal("c", pkg.MostPopularWord2([]string{"a", "b", "c", "c", "d", "e", "e", "d"}))
	a.Equal("a", pkg.MostPopularWord2([]string{"a", "c", "c", "a"}))
}

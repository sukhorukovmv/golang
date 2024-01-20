package solution

import (
	"package/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatinLetters(t *testing.T) {
	a := assert.New(t)
	a.Equal("abc", pkg.LatinLetters(" abc1"))
	a.Equal("", pkg.LatinLetters(" привет"))
	a.Equal("test", pkg.LatinLetters("11 ! t e    s t %)"))
	a.Equal("John", pkg.LatinLetters("John Уолтер"))
	a.Equal("word", pkg.LatinLetters("wo[r]d"))
}

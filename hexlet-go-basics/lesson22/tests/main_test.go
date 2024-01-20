package solution

import (
	"package/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, pkg.IsASCII(" abc1"))
	a.Equal(false, pkg.IsASCII("хекслет"))
	a.Equal(false, pkg.IsASCII("hello \U0001F970"))
}

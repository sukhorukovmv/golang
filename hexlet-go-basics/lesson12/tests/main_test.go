package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"package/pkg"
)

func TestErrorMessageToCode(t *testing.T) {
	a := assert.New(t)
	a.Equal(0, pkg.ErrorMessageToCode("OK"))
	a.Equal(1, pkg.ErrorMessageToCode("CANCELLED"))
	a.Equal(2, pkg.ErrorMessageToCode("UNKNOWN"))
	a.Equal(2, pkg.ErrorMessageToCode("err"))
}

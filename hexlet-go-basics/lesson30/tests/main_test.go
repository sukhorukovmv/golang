package solution

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"package/pkg"
)

func TestGetErrorMsg(t *testing.T) {
	a := assert.New(t)
	a.Equal(pkg.ErrBadRequest.Error(), pkg.GetErrorMsg(pkg.ErrBadRequest))
	a.Equal(pkg.ErrBadConnection.Error(), pkg.GetErrorMsg(pkg.ErrBadConnection))
	a.Equal("", pkg.GetErrorMsg(pkg.NonCriticalError{}))
	a.Equal(pkg.UnknownErrorMsg, pkg.GetErrorMsg(errors.New("unknown")))
	a.Equal(pkg.ErrBadRequest.Error(), pkg.GetErrorMsg(fmt.Errorf("wrap: %w", pkg.ErrBadRequest)))
	a.Equal(
		pkg.ErrBadConnection.Error(),
		pkg.GetErrorMsg(fmt.Errorf("wrap: %w", pkg.ErrBadConnection)),
	)
	a.Equal("", pkg.GetErrorMsg(fmt.Errorf("wrap: %w", pkg.NonCriticalError{})))
}

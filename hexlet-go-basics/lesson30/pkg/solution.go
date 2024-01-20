package pkg

import (
	"errors"
)

type NonCriticalError struct{}

func (e NonCriticalError) Error() string {
	return "validation error"
}

var (
	ErrBadConnection = errors.New("bad connection")
	ErrBadRequest    = errors.New("bad request")
)

const UnknownErrorMsg = "unknown error"

// BEGIN (write your solution here)
func GetErrorMsg(err error) string {
	if errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}
	switch {
	case errors.Is(err, NonCriticalError{}):
		return ""
	case errors.Is(err, ErrBadRequest):
		return err.Error()
	case errors.Is(err, ErrBadConnection):
		return err.Error()
	default:
		return UnknownErrorMsg
	}
}

var criticalErrs = []error{ErrBadRequest, ErrBadConnection}

// GetErrorMsg returns the err message if the error is critical. Otherwise it returns an empty string.
func GetErrorMsg1(err error) string {
	for _, crErr := range criticalErrs {
		if errors.Is(err, crErr) {
			return crErr.Error()
		}
	}

	if errors.As(err, &NonCriticalError{}) {
		return ""
	}

	return UnknownErrorMsg
}

// END

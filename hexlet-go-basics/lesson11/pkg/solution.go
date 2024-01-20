package pkg

import (
	"strings"
)

type UserCreateRequest struct {
	FirstName string
	Age       int
}

var invalidRequest = "invalid request"

// BEGIN (write your solution here)
func Validate(req UserCreateRequest) string {
	if req.Age <= 0 || req.Age > 150 || strings.Contains(req.FirstName, " ") || req.FirstName == "" {
		return invalidRequest
	}

	return ""
}

// END

package solution

import (
	"github.com/stretchr/testify/assert"
	"package/pkg"
	"testing"
)

func TestValidate(t *testing.T) {
	a := assert.New(t)
	a.Equal("invalid request", pkg.Validate(pkg.UserCreateRequest{
		FirstName: "",
		Age:       0,
	}))
	a.Equal("invalid request", pkg.Validate(pkg.UserCreateRequest{
		FirstName: "John",
		Age:       -5,
	}))
	a.Equal("invalid request", pkg.Validate(pkg.UserCreateRequest{
		FirstName: "Andy",
		Age:       0,
	}))
	a.Equal("invalid request", pkg.Validate(pkg.UserCreateRequest{
		FirstName: "Karl",
		Age:       151,
	}))
	a.Equal("invalid request", pkg.Validate(pkg.UserCreateRequest{
		FirstName: "",
		Age:       5,
	}))
	a.Equal("invalid request", pkg.Validate(pkg.UserCreateRequest{
		FirstName: " Henry",
		Age:       15,
	}))
	a.Equal("invalid request", pkg.Validate(pkg.UserCreateRequest{
		FirstName: "John Smith",
		Age:       15,
	}))
	a.Equal("", pkg.Validate(pkg.UserCreateRequest{
		FirstName: "John",
		Age:       150,
	}))
	a.Equal("", pkg.Validate(pkg.UserCreateRequest{
		FirstName: "Susan",
		Age:       30,
	}))
}


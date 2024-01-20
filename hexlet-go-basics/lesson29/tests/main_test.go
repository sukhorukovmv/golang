package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"package/pkg"
)

func TestDecodeAndValidateRequest(t *testing.T) {
	a := assert.New(t)

	req, err := pkg.DecodeAndValidateRequest(nil)
	a.Error(err)
	a.Equal(pkg.CreateUserRequest{}, req)

	req, err = pkg.DecodeAndValidateRequest([]byte(""))
	a.Error(err)
	a.Equal(pkg.CreateUserRequest{}, req)

	req, err = pkg.DecodeAndValidateRequest([]byte("{}"))
	a.Error(err)
	a.Equal(pkg.CreateUserRequest{}, req)

	req, err = pkg.DecodeAndValidateRequest(
		[]byte("{\"email\":\"\",\"password\":\"test\",\"password_confirmation\":\"test\"}"),
	)
	a.Equal(pkg.ErrEmailRequired, err)
	a.Equal(pkg.CreateUserRequest{}, req)

	req, err = pkg.DecodeAndValidateRequest(
		[]byte("{\"email\":\"test\",\"password\":\"\",\"password_confirmation\":\"test\"}"),
	)
	a.Equal(pkg.ErrPasswordRequired, err)
	a.Equal(pkg.CreateUserRequest{}, req)

	req, err = pkg.DecodeAndValidateRequest(
		[]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"\"}"),
	)
	a.Equal(pkg.ErrPasswordConfirmationRequired, err)
	a.Equal(pkg.CreateUserRequest{}, req)

	req, err = pkg.DecodeAndValidateRequest(
		[]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"test2\"}"),
	)
	a.Equal(pkg.ErrPasswordDoesNotMatch, err)
	a.Equal(pkg.CreateUserRequest{}, req)

	req, err = pkg.DecodeAndValidateRequest(
		[]byte(
			"{\"email\":\"email@test.com\",\"password\":\"passwordtest\",\"password_confirmation\":\"passwordtest\"}",
		),
	)
	a.NoError(err)
	a.Equal(pkg.CreateUserRequest{
		Email:                "email@test.com",
		Password:             "passwordtest",
		PasswordConfirmation: "passwordtest",
	}, req)
}

package pkg

import (
	"encoding/json"
	"errors"
)

type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

var (
	ErrEmailRequired                = errors.New("email is required")
	ErrPasswordRequired             = errors.New("password is required")
	ErrPasswordConfirmationRequired = errors.New("password confirmation is required")
	ErrPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

// BEGIN (write your solution here)
func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	reqU := CreateUserRequest{}
	if err := json.Unmarshal(requestBody, &reqU); err != nil {
		return reqU, err
	}
	if reqU.Email == "" {
		return CreateUserRequest{}, ErrEmailRequired
	}
	if reqU.Password == "" {
		return CreateUserRequest{}, ErrPasswordRequired
	}
	if reqU.PasswordConfirmation == "" {
		return CreateUserRequest{}, ErrPasswordConfirmationRequired
	}
	if reqU.Password != reqU.PasswordConfirmation {
		return CreateUserRequest{}, ErrPasswordDoesNotMatch
	}
	return reqU, nil
}

// END

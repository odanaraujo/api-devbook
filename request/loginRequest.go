package request

import "errors"

type LoginRequest struct {
	ID       uint64
	Email    string
	Password string
}

func (login *LoginRequest) ValidatePassword() error {
	if login.Password == "" {
		return errors.New("Please inform the password")
	}
	return nil
}

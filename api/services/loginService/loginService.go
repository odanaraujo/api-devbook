package loginService

import "github.com/odanaraujo/api-devbook/request"

type Service interface {
	Login(email string) (request.LoginRequest, error)
}

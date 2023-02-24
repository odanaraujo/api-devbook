package services

import (
	"github.com/odanaraujo/api-devbook/domain"
)

type Service interface {
	SaveUser(user domain.User) (uint64, error)
	GetAll() ([]domain.User, error)
	GetUserID(ID uint64) (domain.User, error)
	UpdateUser(ID uint64, newUser domain.User) (domain.User, error)
	DeleteUser(ID uint64) error
}

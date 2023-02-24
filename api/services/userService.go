package services

import (
	"github.com/odanaraujo/api-devbook/domain"
)

type Service interface {
	SaveUser(user domain.Users) (uint64, error)
	GetAll() ([]domain.Users, error)
	GetUserID(ID uint64) (domain.Users, error)
	UpdateUser(ID uint64, newUser domain.Users) (domain.Users, error)
	DeleteUser(ID uint64) error
}

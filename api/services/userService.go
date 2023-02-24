package services

import (
	"github.com/odanaraujo/api-devbook/domain"
)

type Service interface {
	SaveUser(user domain.Users) (uint64, error)
	GetAll() ([]domain.Users, error)
	GetUserID(ID uint64) (domain.Users, error)
}

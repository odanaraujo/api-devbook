package userService

import (
	"github.com/odanaraujo/api-devbook/domain"
)

type Service interface {
	SaveUser(user domain.User) (uint64, error)
	GetAll() ([]domain.User, error)
	GetUserID(ID uint64) (domain.User, error)
	UpdateUser(ID uint64, newUser domain.User) (domain.User, error)
	DeleteUser(ID uint64) error
	FollowUser(userID uint64, followID uint64) error
	UnfollowUser(userID uint64, followID uint64) error
	GetfollowersUser(userID uint64) ([]domain.User, error)
	GetFollowingUser(userID uint64) ([]domain.User, error)
	UpdatePassword(userID uint64, passwordHash string) error
	GetPassword(ID uint64) (string, error)
}

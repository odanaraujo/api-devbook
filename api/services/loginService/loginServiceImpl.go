package loginService

import (
	"github.com/odanaraujo/api-devbook/infrastructure/database"
	"github.com/odanaraujo/api-devbook/infrastructure/repository"
	"github.com/odanaraujo/api-devbook/request"
)

func Login(email string) (request.LoginRequest, error) {
	db, err := database.Connection()

	if err != nil {
		return request.LoginRequest{}, err
	}

	defer db.Close()

	repo := repository.NewRepositoryLogin(db)
	request, err := repo.GetUserWithEmail(email)

	if err != nil {
		return request, err
	}

	return request, nil
}

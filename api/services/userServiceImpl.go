package services

import (
	"github.com/odanaraujo/api-devbook/domain"
	"github.com/odanaraujo/api-devbook/infrastructure/database"
	"github.com/odanaraujo/api-devbook/infrastructure/repository"
)

func SaveUSer(user domain.Users) (uint64, error) {

	db, err := database.Connection()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	userId, err := repo.Save(user)

	if err != nil || userId == 0 {
		return 0, err
	}
	return userId, nil
}

func GetAll() ([]domain.Users, error) {
	db, err := database.Connection()

	if err != nil {
		return []domain.Users{}, err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	users, err := repo.GetAll()

	if err != nil {
		return []domain.Users{}, err
	}

	return users, nil
}

func GetUserID(ID uint64) (domain.Users, error) {
	db, err := database.Connection()

	if err != nil {
		return domain.Users{}, err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	user, err := repo.GetUserId(ID)

	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

package userService

import (
	"github.com/odanaraujo/api-devbook/domain"
	"github.com/odanaraujo/api-devbook/infrastructure/database"
	"github.com/odanaraujo/api-devbook/infrastructure/repository"
)

func SaveUSer(user domain.User) (uint64, error) {

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

func GetAll(nameOrNick string) ([]domain.User, error) {
	db, err := database.Connection()

	if err != nil {
		return []domain.User{}, err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	users, err := repo.GetAll(nameOrNick)

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func GetUserID(ID uint64) (domain.User, error) {
	db, err := database.Connection()

	if err != nil {
		return domain.User{}, err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	user, err := repo.GetUserId(ID)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func GetPassword(ID uint64) (string, error) {
	db, err := database.Connection()

	if err != nil {
		return "", err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	password, err := repo.GetPassword(ID)

	if err != nil {
		return "", err
	}

	return password, nil
}

func UpdateUser(ID uint64, newUser domain.User) (domain.User, error) {
	db, err := database.Connection()

	if err != nil {
		return domain.User{}, err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	newUser, err = repo.UpdateUser(ID, newUser)

	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}

func DeleteUser(ID uint64) error {
	db, err := database.Connection()

	if err != nil {
		return err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	err = repo.DeleteUser(ID)

	if err != nil {
		return err
	}

	return nil
}

func FollowUser(userID uint64, followID uint64) error {
	db, err := database.Connection()

	if err != nil {
		return err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	err = repo.FollowUser(userID, followID)

	if err != nil {
		return err
	}

	return nil
}

func UnfollowUser(userID uint64, followID uint64) error {
	db, err := database.Connection()

	if err != nil {
		return err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	err = repo.UnfollowUser(userID, followID)

	if err != nil {
		return err
	}

	return nil
}

func GetFollowersUser(userID uint64) ([]domain.User, error) {
	db, err := database.Connection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	users, err := repo.GetFollowersUser(userID)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetFollowingUser(userID uint64) ([]domain.User, error) {
	db, err := database.Connection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	users, err := repo.GetFollowingUser(userID)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func UpdatePassword(userID uint64, passwordHash string) error {
	db, err := database.Connection()

	if err != nil {
		return err
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	err = repo.UpdatePassword(userID, passwordHash)

	if err != nil {
		return err
	}

	return nil
}

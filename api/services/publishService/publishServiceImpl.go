package publishService

import (
	"github.com/odanaraujo/api-devbook/domain"
	"github.com/odanaraujo/api-devbook/infrastructure/database"
	"github.com/odanaraujo/api-devbook/infrastructure/repository"
)

func CreaterPublish(publish domain.Publish) (uint64, error) {
	db, err := database.Connection()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	repo := repository.NewRepositoryPublish(db)
	ID, err := repo.CreaterPublish(publish)

	if err != nil {
		return 0, err
	}

	return ID, nil
}

func GetPublish(ID uint64) (domain.Publish, error) {
	db, err := database.Connection()

	if err != nil {
		return domain.Publish{}, err
	}

	repo := repository.NewRepositoryPublish(db)
	publish, err := repo.GetPublish(ID)

	if err != nil {
		return domain.Publish{}, err
	}

	return publish, nil
}

func GetAllPublish(ID uint64) ([]domain.Publish, error) {
	db, err := database.Connection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	repo := repository.NewRepositoryPublish(db)
	users, err := repo.GetAllPublish(ID)

	if err != nil {
		return nil, err
	}

	return users, nil
}

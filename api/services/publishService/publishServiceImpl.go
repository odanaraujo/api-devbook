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

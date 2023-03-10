package publishService

import "github.com/odanaraujo/api-devbook/domain"

type Service interface {
	CreaterPublish(publish domain.Publish) (uint64, error)
	GetPublish(ID uint64) (domain.Publish, error)
}

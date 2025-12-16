package entityService

import (
	"ostadbun/interface/entity"
	entityrepository "ostadbun/repository/entity"
)

type Entity struct {
	repo entityrepository.EntityRepo
}

func Config(r entityrepository.EntityRepo) Entity {
	return Entity{
		repo: r,
	}
}

func (a Entity) Search(entity string) []entity.IEntity {

	e, err := a.repo.Search(entity)

	if err != nil {
		return nil
	}
	return e

}

package service

import (
	"errors"
	"gobasic/entity"
	"gobasic/repository"
)

type ServiceCategory struct {
	Repository repository.CategoryRepository
}

func (service ServiceCategory) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}
package repository

import "gobasic/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
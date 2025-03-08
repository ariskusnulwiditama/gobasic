package test

import (
	"gobasic/repository"
	"gobasic/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = service.ServiceCategory{Repository: categoryRepository}

func TestCategoryServiceGet(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}


package Service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"say_hello_unit_test/Repository"
	"say_hello_unit_test/entity"
	"testing"
)

var categoryRepository = &Repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

// unit test with return value as category(id and name)

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Android",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	r, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, category.Id, r.Id)
	assert.Equal(t, category.Name, r.Name)
}

//unit test if no return value

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

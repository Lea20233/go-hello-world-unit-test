package Service

import (
	"errors"
	"say_hello_unit_test/Repository"
	"say_hello_unit_test/entity"
)

type CategoryService struct {
	Repository Repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}

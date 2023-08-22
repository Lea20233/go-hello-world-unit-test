package Repository

import "say_hello_unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

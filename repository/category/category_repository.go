package category

import (
	"learning/restapi/model/domain"
)

type CategoryRepository interface {
	Create(category *domain.Category) *domain.Category
	Update(category *domain.Category) (*domain.Category, error)
	Delete(id int)
	GetById(id int) (*domain.Category, error)
	GetAll() []*domain.Category
	GetAllByUserId(userId string) []*domain.Category
	GetByIdAndUserId(id int, userId string) (*domain.Category, error)
}

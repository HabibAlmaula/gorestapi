package user

import (
	"learning/restapi/model/domain"
)

type UserRepository interface {
	Create(user *domain.User) *domain.User
	Update(user *domain.User) *domain.User
	Delete(id int)
	GetById(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	GetAll() []*domain.User
}

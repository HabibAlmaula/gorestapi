package user

import (
	"errors"
	"gorm.io/gorm"
	"learning/restapi/helper"
	"learning/restapi/model/domain"
)

type UserRepositoryImpl struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn}
}

func (u *UserRepositoryImpl) Create(user *domain.User) *domain.User {
	tx := u.Conn.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		helper.PanicIfError(err)
	}
	tx.Commit()
	return user
}

func (u *UserRepositoryImpl) Update(user *domain.User) *domain.User {
	tx := u.Conn.Begin()
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		helper.PanicIfError(err)
	}
	tx.Commit()
	return user
}

func (u *UserRepositoryImpl) Delete(id int) {
	tx := u.Conn.Begin()
	if err := tx.Delete(&domain.User{}, id).Error; err != nil {
		tx.Rollback()
		helper.PanicIfError(err)
	}
	tx.Commit()
}

func (u *UserRepositoryImpl) GetById(id int) (*domain.User, error) {
	tx := u.Conn.Begin()
	user := &domain.User{}
	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		} else {
			helper.PanicIfError(err)
		}
	}
	tx.Commit()
	return user, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*domain.User, error) {
	tx := u.Conn.Begin()
	user := &domain.User{}
	if err := tx.Find(&user, "email = ?", email).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		} else {
			helper.PanicIfError(err)
		}
	}
	tx.Commit()
	return user, nil
}

func (u *UserRepositoryImpl) GetAll() []*domain.User {
	tx := u.Conn.Begin()
	var users []*domain.User
	if err := tx.Find(&users).Error; err != nil {
		tx.Rollback()
		helper.PanicIfError(err)
	}
	tx.Commit()
	return users
}

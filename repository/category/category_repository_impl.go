package category

import (
	"errors"
	"gorm.io/gorm"
	"learning/restapi/helper"
	"learning/restapi/model/domain"
)

type CategoryRepositoryImpl struct {
	Conn *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{Conn: conn}
}

func (c *CategoryRepositoryImpl) Create(category *domain.Category) *domain.Category {
	tx := c.Conn.Begin()
	if er := tx.Create(&category).Error; er != nil {
		tx.Rollback()
		helper.PanicIfError(er)
	}
	tx.Commit()
	return category
}

func (c *CategoryRepositoryImpl) Update(category *domain.Category) (cat *domain.Category, err error) {
	tx := c.Conn.Begin()
	if er := tx.Model(category).Update("name", category.Name).Error; er != nil {
		tx.Rollback()
		if errors.Is(er, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		} else {
			helper.PanicIfError(er)
		}
	}
	return category, nil
}

func (c CategoryRepositoryImpl) Delete(id int) {
	tx := c.Conn.Begin()
	if er := tx.Delete(&domain.Category{}, id).Error; er != nil {
		tx.Rollback()
		helper.PanicIfError(er)
	}
	tx.Commit()
}

func (c CategoryRepositoryImpl) GetById(id int) (*domain.Category, error) {
	tx := c.Conn.Begin()
	category := &domain.Category{}
	if er := tx.First(&category, id).Error; er != nil {
		tx.Rollback()
		if errors.Is(er, gorm.ErrRecordNotFound) {
			return category, errors.New("category not found")
		} else {

			helper.PanicIfError(er)
		}
	}
	tx.Commit()
	return category, nil
}

func (c CategoryRepositoryImpl) GetAll() []*domain.Category {

	var categories []*domain.Category
	tx := c.Conn.Begin()
	if err := tx.Find(&categories).Error; err != nil {
		tx.Rollback()
		helper.PanicIfError(err)
	}
	return categories
}

func (c *CategoryRepositoryImpl) GetAllByUserId(userId string) []*domain.Category {
	tx := c.Conn.Begin()
	var categories []*domain.Category
	if err := tx.Where("user_id = ?", userId).Find(&categories).Error; err != nil {
		tx.Rollback()
		helper.PanicIfError(err)
	}
	return categories
}

func (c *CategoryRepositoryImpl) GetByIdAndUserId(id int, userId string) (*domain.Category, error) {
	tx := c.Conn.Begin()
	category := &domain.Category{}
	if er := tx.Where("id = ? AND user_id = ?", id, userId).First(&category).Error; er != nil {
		tx.Rollback()
		if errors.Is(er, gorm.ErrRecordNotFound) {
			return category, errors.New("category not found")
		} else {
			helper.PanicIfError(er)
		}
	}
	tx.Commit()
	return category, nil
}

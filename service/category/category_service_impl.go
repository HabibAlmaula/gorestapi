package category

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"learning/restapi/exception"
	"learning/restapi/helper"
	"learning/restapi/model/domain"
	"learning/restapi/model/web/request"
	"learning/restapi/model/web/response"
	"learning/restapi/repository/category"
)

type CategoryServiceImpl struct {
	CategoryRepository category.CategoryRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository category.CategoryRepository, db *gorm.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}

}

func (c *CategoryServiceImpl) Create(request request.CategoryCreateRequest, userId string) response.CategoryResponse {
	// Validate request
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	category := &domain.Category{
		Name:   request.Name,
		UserId: userId,
	}

	category = c.CategoryRepository.Create(category)

	return helper.ToCategoryResponse(category)

}

func (c *CategoryServiceImpl) Update(updateRequest request.CategoryUpdateRequest, id int) response.CategoryResponse {
	// Validate request
	err := c.Validate.Struct(updateRequest)
	helper.PanicIfError(err)

	// Get category by id
	category, err := c.CategoryRepository.GetById(id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Update category
	category.Name = updateRequest.Name

	// Save to database
	category, err = c.CategoryRepository.Update(category)
	if err != nil {
		helper.PanicIfError(err)
	}

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) Delete(id int) {
	// Get category by id
	category, err := c.CategoryRepository.GetById(id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	// Delete category
	c.CategoryRepository.Delete(category.Id)
}

func (c *CategoryServiceImpl) GetById(id int) response.CategoryResponse {
	// Get category by id
	category, err := c.CategoryRepository.GetById(id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) GetAll(ctx context.Context) []response.CategoryResponse {

	categories := c.CategoryRepository.GetAll()

	var categoryResponses []response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}

	return categoryResponses
}

func (c *CategoryServiceImpl) GetAllByUserId(userId string) []response.CategoryResponse {

	categories := c.CategoryRepository.GetAllByUserId(userId)

	var categoryResponses []response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}

	return categoryResponses
}

func (c *CategoryServiceImpl) GetByIdAndUserId(id int, userId string) response.CategoryResponse {

	// Get category by id and user id
	category, err := c.CategoryRepository.GetByIdAndUserId(id, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResponse(category)
}

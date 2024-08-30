package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"learning/restapi/exception"
	"learning/restapi/helper"
	"learning/restapi/model/domain"
	"learning/restapi/model/web/request"
	"learning/restapi/model/web/response"
	"learning/restapi/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}

}

func (c *CategoryServiceImpl) Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse {
	// Validate request
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = c.CategoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (c *CategoryServiceImpl) Update(ctx context.Context, updateRequest request.CategoryUpdateRequest, id int) response.CategoryResponse {
	// Validate request
	err := c.Validate.Struct(updateRequest)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// Get category by id
	category, err := c.CategoryRepository.GetById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Update category
	category.Name = updateRequest.Name

	// Save to database
	category = c.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// Get category by id
	category, err := c.CategoryRepository.GetById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	// Delete category
	c.CategoryRepository.Delete(ctx, tx, category.Id)
}

func (c *CategoryServiceImpl) GetById(ctx context.Context, id int) response.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// Get category by id
	category, err := c.CategoryRepository.GetById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) GetAll(ctx context.Context) []response.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories := c.CategoryRepository.GetAll(ctx, tx)

	var categoryResponses []response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}

	return categoryResponses
}

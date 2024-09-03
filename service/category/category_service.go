package category

import (
	"context"
	"learning/restapi/model/web/request"
	"learning/restapi/model/web/response"
)

type CategoryService interface {
	Create(request request.CategoryCreateRequest, userId string) response.CategoryResponse
	Update(updateRequest request.CategoryUpdateRequest, id int) response.CategoryResponse
	Delete(id int)
	GetById(id int) response.CategoryResponse
	GetAll(ctx context.Context) []response.CategoryResponse
	GetAllByUserId(userId string) []response.CategoryResponse
	GetByIdAndUserId(id int, userId string) response.CategoryResponse
}

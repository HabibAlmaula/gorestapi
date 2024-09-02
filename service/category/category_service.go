package category

import (
	"context"
	"learning/restapi/model/web/request"
	"learning/restapi/model/web/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest, userId string) response.CategoryResponse
	Update(ctx context.Context, updateRequest request.CategoryUpdateRequest, id int) response.CategoryResponse
	Delete(ctx context.Context, id int)
	GetById(ctx context.Context, id int) response.CategoryResponse
	GetAll(ctx context.Context) []response.CategoryResponse
	GetAllByUserId(ctx context.Context, userId string) []response.CategoryResponse
	GetByIdAndUserId(ctx context.Context, id int, userId string) response.CategoryResponse
}

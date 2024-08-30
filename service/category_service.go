package service

import (
	"context"
	"learning/restapi/model/web/request"
	"learning/restapi/model/web/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, updateRequest request.CategoryUpdateRequest, id int) response.CategoryResponse
	Delete(ctx context.Context, id int)
	GetById(ctx context.Context, id int) response.CategoryResponse
	GetAll(ctx context.Context) []response.CategoryResponse
}

package controller

import (
	"github.com/julienschmidt/httprouter"
	"learning/restapi/helper"
	"learning/restapi/model/base"
	"learning/restapi/model/web/request"
	"learning/restapi/service/category"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	Service category.CategoryService
}

func NewCategoryController(categoryService category.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Service: categoryService,
	}
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := request.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &req)

	userId := r.Header.Get("X-User-ID")

	res := c.Service.Create(req, userId)
	response := base.BaseResponse{
		Code:    201,
		Message: "Success Create Category",
		Succes:  true,
		Data:    res,
	}
	helper.WriteToResponseBody(w, response)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := request.CategoryUpdateRequest{}

	helper.ReadFromRequestBody(r, &req)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	res := c.Service.Update(req, id)
	response := base.BaseResponse{
		Code:    200,
		Message: "Success Update Category",
		Succes:  true,
		Data:    res,
	}

	helper.WriteToResponseBody(w, response)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	c.Service.Delete(id)
	response := base.BaseResponse{
		Code:    200,
		Message: "Success Delete Category",
		Succes:  true,
	}

	helper.WriteToResponseBody(w, response)
}

func (c *CategoryControllerImpl) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	res := c.Service.GetById(id)
	response := base.BaseResponse{
		Code:    200,
		Message: "Success Get Category By Id",
		Succes:  true,
		Data:    res,
	}
	helper.WriteToResponseBody(w, response)

}

func (c *CategoryControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res := c.Service.GetAll(r.Context())
	response := base.BaseResponse{
		Code:    200,
		Message: "Success Get All Category",
		Succes:  true,
		Data:    res,
	}

	helper.WriteToResponseBody(w, response)

}

func (c *CategoryControllerImpl) GetAllByUserId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := r.Header.Get("X-User-ID")
	res := c.Service.GetAllByUserId(userId)
	response := base.BaseResponse{
		Code:    200,
		Message: "Success Get All Category",
		Succes:  true,
		Data:    res,
	}

	helper.WriteToResponseBody(w, response)
}

func (c *CategoryControllerImpl) GetByIdAndUserId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := r.Header.Get("X-User-ID")
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	res := c.Service.GetByIdAndUserId(id, userId)
	response := base.BaseResponse{
		Code:    200,
		Message: "Success Get All Category",
		Succes:  true,
		Data:    res,
	}

	helper.WriteToResponseBody(w, response)
}

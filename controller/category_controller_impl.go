package controller

import (
	"github.com/julienschmidt/httprouter"
	"learning/restapi/helper"
	"learning/restapi/model/base"
	"learning/restapi/model/web/request"
	"learning/restapi/service"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Service: categoryService,
	}
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	req := request.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &req)

	res := c.Service.Create(r.Context(), req)
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

	res := c.Service.Update(r.Context(), req, id)
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

	c.Service.Delete(r.Context(), id)
	response := base.BaseResponse{
		Code:    200,
		Message: "Success Delete Category",
		Succes:  true,
	}

	helper.WriteToResponseBody(w, response)
	//categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	//helper.PanicIfError(err)
	//
	//c.Service.Delete(r.Context(), categoryId)
	//response := base.BaseResponse{
	//	Code:    200,
	//	Message: "Success Delete Category",
	//	Succes:  true,
	//}
	//w.Header().Set("Content-Type", "application/json")
	//errs := json.NewEncoder(w).Encode(response)
	//helper.PanicIfError(errs)
}

func (c *CategoryControllerImpl) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	res := c.Service.GetById(r.Context(), id)
	response := base.BaseResponse{
		Code:    200,
		Message: "Success Get Category By Id",
		Succes:  true,
		Data:    res,
	}
	helper.WriteToResponseBody(w, response)
	//categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	//helper.PanicIfError(err)
	//
	//res := c.Service.GetById(r.Context(), categoryId)
	//response := base.BaseResponse{
	//	Code:    200,
	//	Message: "Success Get Category By Id",
	//	Succes:  true,
	//	Data:    res,
	//}
	//w.Header().Set("Content-Type", "application/json")
	//errs := json.NewEncoder(w).Encode(response)
	//helper.PanicIfError(errs)
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

	//res := c.Service.GetAll(r.Context())
	//response := base.BaseResponse{
	//	Code:    200,
	//	Message: "Success Get All Category",
	//	Succes:  true,
	//	Data:    res,
	//}
	//w.Header().Set("Content-Type", "application/json")
	//encoder := json.NewEncoder(w)
	//errs := encoder.Encode(response)
	//helper.PanicIfError(errs)
}

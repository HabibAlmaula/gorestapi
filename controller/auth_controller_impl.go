package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"learning/restapi/helper"
	"learning/restapi/model/base"
	"learning/restapi/model/web/request"
	"learning/restapi/service/user"
	"net/http"
)

type AuthControllerImpl struct {
	Service user.UserService
}

func NewAuthController(service user.UserService) AuthController {
	return &AuthControllerImpl{Service: service}
}

func (a AuthControllerImpl) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := request.RegisterUserRequest{}
	helper.ReadFromRequestBody(r, &req)

	fmt.Println("Request_controller: ", req)

	res := a.Service.Register(req)
	response := base.BaseResponse{
		Code:    201,
		Message: "Success Register User",
		Succes:  true,
		Data:    res,
	}
	helper.WriteToResponseBody(w, response)
}

func (a AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := request.LoginRequest{}
	helper.ReadFromRequestBody(r, &req)
	fmt.Println("Request_controller: ", req)

	res := a.Service.Login(req)

	fmt.Println("Response_controller: ", res)

	response := base.BaseResponse{
		Code:    200,
		Message: "Success Login",
		Succes:  true,
		Data:    res,
	}
	helper.WriteToResponseBody(w, response)
}

package user

import (
	"learning/restapi/model/web/request"
	"learning/restapi/model/web/response"
)

type UserService interface {
	Register(request request.RegisterUserRequest) response.UserResponse
	Login(request request.LoginRequest) response.LoginResponse
}

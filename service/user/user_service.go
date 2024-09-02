package user

import (
	"context"
	"learning/restapi/model/web/request"
	"learning/restapi/model/web/response"
)

type UserService interface {
	Register(ctx context.Context, request request.RegisterUserRequest) response.UserResponse
	Login(ctx context.Context, request request.LoginRequest) response.LoginResponse
}

package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"learning/restapi/exception"
	"learning/restapi/helper"
	"learning/restapi/model/domain"
	"learning/restapi/model/web/request"
	"learning/restapi/model/web/response"
	"learning/restapi/repository/user"
)

type UserServiceImpl struct {
	UserRepository user.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository user.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (u *UserServiceImpl) Register(ctx context.Context, request request.RegisterUserRequest) response.UserResponse {
	//validate request
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)

	fmt.Println("Request_service: ", request)

	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	//Get user by email
	userByEmail, err := u.UserRepository.GetByEmail(ctx, tx, request.Email)

	if userByEmail != (domain.User{}) {
		panic(exception.NewDataExistError("email already exist"))
	}

	passworHash, err := helper.HashPassword(request.Password)
	fmt.Println("Password Hash: ", passworHash)
	helper.PanicIfError(err)
	usr := domain.User{
		FullName: request.Fullname,
		Email:    request.Email,
		Password: passworHash,
	}
	fmt.Println("User: ", usr)
	usr = u.UserRepository.Create(ctx, tx, usr)
	fmt.Println("User After Create: ", usr)
	return helper.ToUserResponse(usr)
}

func (u *UserServiceImpl) Login(ctx context.Context, request request.LoginRequest) response.LoginResponse {
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)
	fmt.Println("Request_service: ", request)

	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	//Get user by email
	userByEmail, err := u.UserRepository.GetByEmail(ctx, tx, request.Email)
	fmt.Println("User By Email: ", userByEmail)
	if err != nil {
		fmt.Println("Error: ", err)
		helper.PanicIfError(err)
	}

	if userByEmail == (domain.User{}) {
		//unauthorized
		fmt.Println("User Not Found")
		panic(exception.NewUnAuthorizedError("email or password is invalid"))
	} else {
		//compare password
		isPasswordValid := helper.CheckPasswordHash(request.Password, userByEmail.Password)
		fmt.Println("Password Valid: ", isPasswordValid)
		if !isPasswordValid {
			//unauthorized
			panic(exception.NewUnAuthorizedError("email or password is invalid"))
		} else {
			accToken, errAcc := helper.GenerateAccessTokenJWT(&userByEmail)
			helper.PanicIfError(errAcc)
			refToken, errRef := helper.GenerateRefreshTokenJWT(&userByEmail)
			helper.PanicIfError(errRef)

			return helper.ToLoginResponse(userByEmail, accToken, refToken)
		}
	}
}

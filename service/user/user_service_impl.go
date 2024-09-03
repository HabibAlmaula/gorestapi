package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"learning/restapi/exception"
	"learning/restapi/helper"
	"learning/restapi/model/domain"
	"learning/restapi/model/web/request"
	"learning/restapi/model/web/response"
	"learning/restapi/repository/user"
)

type UserServiceImpl struct {
	UserRepository user.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository user.UserRepository, DB *gorm.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (u *UserServiceImpl) Register(request request.RegisterUserRequest) response.UserResponse {
	//validate request
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)

	fmt.Println("Request_service: ", request)

	//Get user by email
	userByEmail, err := u.UserRepository.GetByEmail(request.Email)

	if userByEmail != (&domain.User{}) {
		panic(exception.NewDataExistError("email already exist"))
	}

	passworHash, err := helper.HashPassword(request.Password)
	fmt.Println("Password Hash: ", passworHash)
	helper.PanicIfError(err)
	usr := &domain.User{
		FullName: request.Fullname,
		Email:    request.Email,
		Password: passworHash,
	}
	fmt.Println("User: ", usr)
	usr = u.UserRepository.Create(usr)
	fmt.Println("User After Create: ", usr)
	return helper.ToUserResponse(usr)
}

func (u *UserServiceImpl) Login(request request.LoginRequest) response.LoginResponse {
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)
	fmt.Println("Request_service: ", request)

	//Get user by email
	userByEmail, err := u.UserRepository.GetByEmail(request.Email)
	fmt.Println("User By Email: ", userByEmail)
	if err != nil {
		fmt.Println("Error: ", err)
		helper.PanicIfError(err)
	}

	if userByEmail == (&domain.User{}) {
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
			accToken, expAccess, errAcc := helper.GenerateAccessTokenJWT(userByEmail)
			helper.PanicIfError(errAcc)
			refToken, _, errRef := helper.GenerateRefreshTokenJWT(userByEmail)
			helper.PanicIfError(errRef)

			return helper.ToLoginResponse(userByEmail, accToken, refToken, expAccess)
		}
	}
}

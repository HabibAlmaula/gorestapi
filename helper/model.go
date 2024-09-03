package helper

import (
	"learning/restapi/model/domain"
	"learning/restapi/model/web/response"
	"time"
)

func ToCategoryResponse(category *domain.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToUserResponse(user *domain.User) response.UserResponse {
	return response.UserResponse{
		Id:       user.Id,
		Fullname: user.FullName,
		Email:    user.Email,
	}
}

func ToLoginResponse(user *domain.User, accessToken string, refreshToken string, expTime int64) response.LoginResponse {
	return response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    time.Unix(expTime, 0),
		User:         ToUserResponse(user),
	}
}

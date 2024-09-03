package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"learning/restapi/configs"
	"learning/restapi/model/domain"
	"net/http"
	"strconv"
	"time"
)

var configJWT = configs.Config.JWT

func GenerateAccessTokenJWT(user *domain.User) (string, int64, error) {
	// Define the claims for the JWT token
	expTime := time.Now().Add(time.Hour * configJWT.TTLAccess).Unix()
	claims := jwt.MapClaims{
		"sub": user.Id,           // Subject
		"iat": time.Now().Unix(), // Issued At
		"exp": expTime,           // Expiration time (2 hours)
	}
	fmt.Println("exp: ", expTime)

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(configJWT.SecretAccess))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expTime, nil
}

func GenerateRefreshTokenJWT(user *domain.User) (string, int64, error) {
	// Define the claims for the JWT token
	expTime := time.Now().Add(time.Hour * configJWT.TTLAccess).Unix()

	claims := jwt.MapClaims{
		"sub": user.Id,           // Subject
		"iat": time.Now().Unix(), // Issued At
		"exp": expTime,           // Expiration time (2 hours)
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(configJWT.SecretRefresh))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expTime, nil
}

func GetUserID(r *http.Request) (int32, error) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		return 0, err
	}
	return int32(userID), nil
}

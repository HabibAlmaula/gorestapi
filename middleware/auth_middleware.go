//package middleware
//
//import (
//	"fmt"
//	"github.com/golang-jwt/jwt"
//	"github.com/julienschmidt/httprouter"
//	"learning/restapi/configs"
//	"learning/restapi/helper"
//	"learning/restapi/model/base"
//	"net/http"
//	"strings"
//)
//
//type JwtClaims struct {
//	ID string `json:"sub"`
//	jwt.StandardClaims
//}
//
//type AuthMiddleware struct {
//	Handler httprouter.Handle
//}
//
//func NewAuthMiddleware(handler httprouter.Handle) *AuthMiddleware {
//	return &AuthMiddleware{Handler: handler}
//}
//
//func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	fmt.Println("AuthMiddleware")
//	auth := r.Header.Get("authorization")
//	if !strings.Contains(auth, "Bearer") {
//		fmt.Print("Error line42")
//		w.WriteHeader(http.StatusUnauthorized)
//		response := base.BaseResponse{
//			Code:    http.StatusUnauthorized,
//			Succes:  false,
//			Message: http.StatusText(http.StatusUnauthorized),
//		}
//		helper.WriteToResponseBody(w, response)
//	}
//
//	tokenString := strings.Split(auth, " ")
//
//	fmt.Println("TOKEN_STRING", tokenString)
//
//	//token, errToken := jwt.ParseWithClaims(tokenString[1], &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
//	//	// Ensure the token method is HMAC before using the key
//	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//	//		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//	//	}
//	//	return []byte(configs.Config.JWT.SecretAccess), nil
//	//})
//
//	token, errToken := jwt.ParseWithClaims(tokenString[1], &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(configs.Config.JWT.SecretAccess), nil
//	})
//
//	if errToken != nil {
//		helper.PanicIfError(errToken)
//	}
//
//	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
//		r.Header.Add("X-User-ID", claims.ID)
//		middleware.ServeHTTP(w, r, ps)
//	} else {
//		fmt.Println("Error line71")
//		w.WriteHeader(http.StatusUnauthorized)
//		response := base.BaseResponse{
//			Code:    http.StatusUnauthorized,
//			Succes:  false,
//			Message: http.StatusText(http.StatusUnauthorized),
//		}
//		helper.WriteToResponseBody(w, response)
//	}
//}

package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
	"learning/restapi/configs"
	"learning/restapi/helper"
	"learning/restapi/model/base"
	"net/http"
	"strings"
)

type JwtClaims struct {
	ID string `json:"sub"`
	jwt.StandardClaims
}

type AuthMiddleware struct {
	Handler httprouter.Handle
}

func NewAuthMiddleware(handler httprouter.Handle) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("AuthMiddleware")
	auth := r.Header.Get("authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		fmt.Println("Error line42")
		w.WriteHeader(http.StatusUnauthorized)
		response := base.BaseResponse{
			Code:    http.StatusUnauthorized,
			Succes:  false,
			Message: http.StatusText(http.StatusUnauthorized),
		}
		helper.WriteToResponseBody(w, response)
		return
	}

	tokenString := strings.TrimPrefix(auth, "Bearer ")

	fmt.Println("TOKEN_STRING", tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method is HMAC before using the key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.Config.JWT.SecretAccess), nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Error parsing token", err)
		w.WriteHeader(http.StatusUnauthorized)
		response := base.BaseResponse{
			Code:    http.StatusUnauthorized,
			Succes:  false,
			Message: "Unauthorized",
		}
		helper.WriteToResponseBody(w, response)
		return
	}

	// If everything is fine, call the next handler
	middleware.Handler(w, r, ps)
}

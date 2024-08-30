package middleware

import (
	"learning/restapi/helper"
	"learning/restapi/model/base"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "KEY" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
		response := base.BaseResponse{
			Code:    http.StatusUnauthorized,
			Succes:  false,
			Message: http.StatusText(http.StatusUnauthorized),
		}
		helper.WriteToResponseBody(writer, response)
	}
}

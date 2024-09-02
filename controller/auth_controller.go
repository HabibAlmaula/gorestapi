package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthController interface {
	Register(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

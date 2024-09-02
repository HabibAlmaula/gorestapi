package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetAllByUserId(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetByIdAndUserId(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

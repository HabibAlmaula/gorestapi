package exception

import (
	"github.com/go-playground/validator/v10"
	"learning/restapi/helper"
	"learning/restapi/model/base"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if dataExistError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func dataExistError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	dataError, ok := err.(DataExistError)
	if !ok {
		return false
	} else {
		writer.WriteHeader(http.StatusConflict)
		response := base.BaseResponse{
			Code:    http.StatusConflict,
			Succes:  false,
			Message: dataError.Error,
		}
		helper.WriteToResponseBody(writer, response)

		return true
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		response := base.BaseResponse{
			Code:    http.StatusBadRequest,
			Succes:  false,
			Message: http.StatusText(http.StatusBadRequest),
		}
		helper.WriteToResponseBody(writer, response)

		return true
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(NotFoundError)
	if !ok {
		return false
	} else {
		writer.WriteHeader(http.StatusNotFound)
		response := base.BaseResponse{
			Code:    http.StatusNotFound,
			Succes:  false,
			Message: http.StatusText(http.StatusNotFound),
		}
		helper.WriteToResponseBody(writer, response)

		return true
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.WriteHeader(http.StatusInternalServerError)
	response := base.BaseResponse{
		Code:    http.StatusInternalServerError,
		Succes:  false,
		Message: http.StatusText(http.StatusInternalServerError),
	}
	helper.WriteToResponseBody(writer, response)
}

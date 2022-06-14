package exception

import (
	"golearning/restapi/helper"
	"golearning/restapi/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if NotFoundError404(writer, request, err) {
		return
	}
	if ValidationError400(writer, request, err) {
		return
	}
	InternalServerError500(writer, request, err)
}

func ValidationError400(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		apiResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "ERROR BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}
}

func NotFoundError404(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(Error404)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		apiResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "ERROR NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}
}

func InternalServerError500(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	apiResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

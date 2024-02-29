package Helper

import (
	"UserLoginAPI/Model/Web"
	"net/http"
)

func WriteOk(writer http.ResponseWriter, response any) {
	web_response := Web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}

	WriteToResponseBody(writer, web_response)
}

func WriteUnAuthorized(writer http.ResponseWriter) {
	web_response := Web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "Unauthorized",
		Data:   UNAUTHORIZED,
	}

	WriteToResponseBody(writer, web_response)
}

func WriteBadRequest(writer http.ResponseWriter) {
	web_response := Web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "Bad Request",
		Data:   BAD_REQUEST,
	}
	WriteToResponseBody(writer, web_response)
}

func WriteNotFound(writer http.ResponseWriter) {
	web_response := Web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   NOT_FOUND,
	}
	WriteToResponseBody(writer, web_response)
}

func WriteFailValidation(writer http.ResponseWriter) {
	web_response := Web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   FAIL_VALIDATION,
	}
	WriteToResponseBody(writer, web_response)
}

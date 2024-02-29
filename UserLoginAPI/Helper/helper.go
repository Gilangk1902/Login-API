package Helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, writer http.ResponseWriter, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&result)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

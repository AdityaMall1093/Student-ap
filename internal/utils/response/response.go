package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	Status      = "Ok"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func GenralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

package http

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}

func handleError(w http.ResponseWriter, err error, status int) {
	response := Data{
		Message: err.Error(),
		Error:   true,
	}

	payload, _ := json.Marshal(response)

	w.WriteHeader(status)
	w.Write(payload)
}

func handle(w http.ResponseWriter, message string, data interface{}, status int) {
	response := Data{
		Message: message,
		Error:   false,
		Data:    data,
	}

	payload, _ := json.Marshal(response)

	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json")
	w.Write(payload)
}

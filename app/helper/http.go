package helper

import (
	"encoding/json"
	"net/http"
)

// errorResponse type
type errorResponse struct {
	Error string `json:"error"`
}

// parses any type of response to JSON and writes to the ResponseWriter
func WriteJSON(writer http.ResponseWriter, status int, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(response)
}

// Writes error message to responseWriter
func WriteError(writer http.ResponseWriter, status int, response interface{}) error {
	return WriteJSON(writer, status, errorResponse{Error: response.(string)})
}

package response

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type jsonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func JSON(w http.ResponseWriter, status int, data any) {

	buffer := new(bytes.Buffer)

	err := json.NewEncoder(buffer).Encode(data)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to encode json response."))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buffer.Bytes())
}

func JSONError(w http.ResponseWriter, status int, message string) {

	payload := jsonError{
		Code:    status,
		Message: message,
	}

	buffer := new(bytes.Buffer)

	err := json.NewEncoder(buffer).Encode(payload)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(status)
		w.Write([]byte("Failed to encode json response."))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buffer.Bytes())
}

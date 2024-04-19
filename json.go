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
		http.Error(w, "Failed to encode JSON response.", http.StatusInternalServerError)
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
		http.Error(w, "Failed to encode JSON response.", status)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buffer.Bytes())
}

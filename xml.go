package response

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

type xmlError struct {
	xml.Name `xml:"Error"`
	Code     int    `xml:"code"`
	Message  string `xml:"message"`
}

func XML(w http.ResponseWriter, status int, data any) {

	buffer := new(bytes.Buffer)

	err := xml.NewEncoder(buffer).Encode(data)
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to encode xml response."))
		return
	}

	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(status)
	w.Write(buffer.Bytes())
}

func XMLError(w http.ResponseWriter, status int, message string) {

	payload := xmlError{
		Code:    status,
		Message: message,
	}

	buffer := new(bytes.Buffer)

	err := xml.NewEncoder(buffer).Encode(payload)
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

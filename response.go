package response

import "net/http"

func Status(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func Error(w http.ResponseWriter, status int, message string) {
	http.Error(w, message, status)
}

func Content(w http.ResponseWriter, r *http.Request, status int, data any) {

	header := r.Header.Get("Accept")

	if header == "application/json" {
		JSON(w, status, data)
		return
	}

	if header == "application/xml" {
		XML(w, status, data)
		return
	}

	Status(w, http.StatusUnsupportedMediaType)
}

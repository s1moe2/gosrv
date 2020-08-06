package handlers

import (
	"encoding/json"
	"net/http"
)

type appError struct {
	Status int
	Message string
}

// internalError returns an appError with 500 code and default message
func internalError() appError {
	return appError{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
	}
}

// respond is an helper that takes care of the
// HTTP response part of a request handler
func respond(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		// TODO add log on error
	}

}

// respondError is an helper similar to respond but only used for internal errors
func respondError(w http.ResponseWriter, e appError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.Status)
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		// TODO add log on error
	}
}
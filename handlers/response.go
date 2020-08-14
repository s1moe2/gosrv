package handlers

import (
	"encoding/json"
	"net/http"
)

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

// respondError is an helper similar to respond but only used for custom errors
func respondError(w http.ResponseWriter, e customError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.StatusCode())
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		// TODO add log on error
	}
}

// respondInternalError is an helper similar to respondError but responds
// with a default internal error code and payload
func respondInternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(newInternalError())
	if err != nil {
		// TODO add log on error
	}
}
package api

import (
	"encoding/json"
	"net/http"
)

type coinBalance struct {
	Username string
}

type coinBalanceResponse struct {
	Code    int
	Balance int64
}

type Error struct {
	Code int

	message string
}

func writerError(w http.ResponseWriter, message string, statusCode int) {
	resp := Error{
		Code:    statusCode,
		message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(resp)
}

var RequestErrorHandler = func(w http.ResponseWriter, err error) {
	writerError(w, err.Error(), http.StatusBadRequest)
}

var InternalErrorHandler = func(w http.ResponseWriter) {
	writerError(w, "An unexpectd error happened", http.StatusInternalServerError)
}

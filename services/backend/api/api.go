package api

import (
	"encoding/json"
	"net/http"
)

// Balance params
type BalanceParams struct {
	Username string
}

type AccountParams struct {
	Username string
	Balance  int
}

type BalanceResponse struct {
	// Response code
	Code int

	// Account balance
	Balance int
}

type UserParams struct {
	Username string
	Token    string
}

type UserResponse struct {
	// Response code
	Code int

	// Username
	Username string

	// Token
	Token string

	// Account Balance
	Balance int
}

type Response struct {
	// Response code
	Code int

	// Message
	Message string
}

type Error struct {
	// Error response code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)

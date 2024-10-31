package api

import (
	"encoding/json"
	"net/http"
)

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

func raiseRiskStoreError(w http.ResponseWriter, code int, message string) {
	err := Error{
		Code:    int32(code),
		Message: message,
	}

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(err)
}

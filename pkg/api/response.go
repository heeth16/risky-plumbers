package api

import (
	"encoding/json"
	"net/http"
)

// Error defines model for Error.
type Response struct {
	// Message Error message
	Message string `json:"message"`
}

func returnRiskStoreResponse(w http.ResponseWriter, statusCode int, message string) {
	err := Response{
		Message: message,
	}

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(err)
}

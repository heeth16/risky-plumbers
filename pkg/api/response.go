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

// returnRiskStoreResponse is a utility function to return a JSON encoded message
// to the given http.ResponseWriter with the given HTTP status code and
// message.
func returnRiskStoreResponse(w http.ResponseWriter, statusCode int, message string) {
	err := Response{
		Message: message,
	}

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(err)
}

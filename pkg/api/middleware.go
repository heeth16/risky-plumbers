package api

import "net/http"

// MiddlewareContentTypeJSON is a middleware that sets the Content-Type header
// of the response to "application/json" for all incoming HTTP requests.
// It wraps the provided http.Handler and ensures that the Content-Type is
// set before passing control to the next handler in the chain.
func MiddlewareContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

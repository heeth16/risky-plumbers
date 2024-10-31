package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heeth16/risky-plumbers/pkg/api"
)

// main is the entry point for the Risk Server.
//
// It creates a new RiskStore and a new GorillaServerOptions instance.
// The GorillaServerOptions are configured as follows:
// - BaseURL is set to "/v1"
// - BaseRouter is set to a new mux.Router.
// - Middlewares is set to a slice containing the MiddlewareContentTypeJSON function.
//
// It then calls the HandlerWithOptions function to get a handler for the server.
// Finally, it creates a new http.Server instance and calls ListenAndServe on it.
func main() {
	store := api.NewRiskStore()
	options := api.GorillaServerOptions{
		BaseURL:     "/v1",
		BaseRouter:  mux.NewRouter(),
		Middlewares: []api.MiddlewareFunc{api.MiddlewareContentTypeJSON},
	}

	hander := api.HandlerWithOptions(store, options)

	s := &http.Server{
		Handler: hander,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}

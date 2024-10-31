package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heeth16/risky-plumbers/pkg/api"
)

func main() {
	fmt.Println("Starting Risk Server")
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

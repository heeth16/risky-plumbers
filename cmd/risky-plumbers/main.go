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
	server := api.NewRiskStore()
	options := api.GorillaServerOptions{
		BaseURL:    "/v1",
		BaseRouter: mux.NewRouter(),
	}

	hander := api.HandlerWithOptions(server, options)

	s := &http.Server{
		Handler: hander,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}

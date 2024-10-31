package main

import (
	"flag"
	"log"
	"net"
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
	// Define command-line flags for IP and port
	ip := flag.String("ip", "127.0.0.1", "IP address to listen on")
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()
	addr := net.JoinHostPort(*ip, *port)

	store := api.NewRiskStore()
	options := api.GorillaServerOptions{
		BaseURL:     "/v1",
		BaseRouter:  mux.NewRouter(),
		Middlewares: []api.MiddlewareFunc{api.MiddlewareContentTypeJSON},
	}

	hander := api.HandlerWithOptions(store, options)

	// Create a new HTTP server and listen on the specified address
	s := &http.Server{
		Handler: hander,
		Addr:    addr,
	}

	log.Printf("Starting server on %s", addr)
	log.Fatal(s.ListenAndServe())
}

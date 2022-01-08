package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sushi/api"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil
	evidentlyApp := api.NewEvidentlyApp()

	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidator(swagger))

	api.HandlerFromMux(evidentlyApp, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1:%d", *port),
	}

	log.Fatal(s.ListenAndServe())
}

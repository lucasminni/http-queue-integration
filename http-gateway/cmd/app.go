package cmd

import (
	httpServer "http-gateway/internal/infra/server"
	"http-gateway/internal/settings"
	"log"
)

var Mode string

func Start() {
	if err := settings.LoadEnvs(); err != nil {
		log.Fatal("Error trying to load environment variables")
	}

	if err := httpServer.StartHTTPServer(); err != nil {
		log.Fatal("Error trying to start the HTTP server")
	}
}

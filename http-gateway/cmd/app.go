package cmd

import (
	httpServer "http-gateway/internal/infra/server"
	"log"
)

func Start() {
	log.Println("Starting HTTP server...")
	if err := httpServer.StartHTTPServer(); err != nil {
		log.Fatal("Error trying to start the HTTP server")
	}
}

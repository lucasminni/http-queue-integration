package cmd

import (
	postgres "http-gateway/internal/infra/database/sql"
	httpServer "http-gateway/internal/infra/server"
	"log"
)

func Start() {
	log.Println("Starting postgres database connection...")
	if err := postgres.StartPostgresDatabase(); err != nil {
		log.Fatal("Error trying to start Postgres database")
	}

	log.Println("Starting HTTP server...")
	if err := httpServer.StartHTTPServer(); err != nil {
		log.Fatal("Error trying to start the HTTP server")
	}
}

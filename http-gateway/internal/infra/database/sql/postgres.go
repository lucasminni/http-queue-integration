package infra

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "orders"
)

func StartPostgresDatabase() (*sql.DB, error) {
	config := fmt.Sprintf(
		"host=%s "+
			"port=%d "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", config)

	if err != nil {
		log.Panic("Error while database start - " + err.Error())
		return nil, err
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Panic("Database health check failed - " + err.Error())
		return nil, err
	}

	return db, nil
}

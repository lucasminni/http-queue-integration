package infra

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "orders"
)

func StartPostgresDatabase() error {
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

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Panic("Error while database start - " + err.Error())
		return nil
	}

	sqlDB, _ := db.DB()

	err = sqlDB.Ping()

	if err != nil {
		log.Panic("Database health check failed - " + err.Error())
		return nil
	}

	return nil
}

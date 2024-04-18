package common

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func SetUpDatabaseConnection() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", GetEnv().POSTGRESDB_HOST, GetEnv().POSTGRESDB_USERNAME, GetEnv().POSTGRESDB_PASSWORD, GetEnv().POSTGRESDB_DB, GetEnv().POSTGRESDB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal("Failed to connect database.")
	}

	DBInstance = db
}

package database

import (
	"fmt"
	"log"

	"github.com/SaiHLu/rest-template/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpPostgresDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", common.GetEnv().POSTGRESDB_HOST, common.GetEnv().POSTGRESDB_USERNAME, common.GetEnv().POSTGRESDB_PASSWORD, common.GetEnv().POSTGRESDB_DB, common.GetEnv().POSTGRESDB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal("Failed to connect database.")
	}

	return db
}

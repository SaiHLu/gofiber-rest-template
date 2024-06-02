package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SaiHLu/rest-template/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,   // Slow SQL threshold
		LogLevel:                  logger.Silent, // Log level
		IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      false,         // Don't include params in the SQL log
		Colorful:                  true,          // Disable color
	},
)

func SetUpPostgresDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", common.GetEnv().POSTGRESDB_HOST, common.GetEnv().POSTGRESDB_USERNAME, common.GetEnv().POSTGRESDB_PASSWORD, common.GetEnv().POSTGRESDB_DB, common.GetEnv().POSTGRESDB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true, Logger: newLogger})

	if err != nil {
		log.Fatal("Failed to connect database.")
	}

	return db
}

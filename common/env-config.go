package common

import (
	"log"
	"os"
	"strconv"
)

type envConfig struct {
	// DB
	POSTGRESDB_HOST     string
	POSTGRESDB_PORT     int
	POSTGRESDB_DB       string
	POSTGRESDB_USERNAME string
	POSTGRESDB_PASSWORD string

	//  Redis
	REDIS_HOST string
	REDIS_PORT string
}

func GetEnv() *envConfig {
	DB_PORT, err := strconv.Atoi(os.Getenv("POSTGRESDB_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	// REDIS_PORT, err := ConvertStringToInt(os.Getenv("REDIS_PORT"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return &envConfig{
		// DB
		POSTGRESDB_HOST:     os.Getenv("POSTGRESDB_HOST"),
		POSTGRESDB_PORT:     DB_PORT,
		POSTGRESDB_DB:       os.Getenv("POSTGRESDB_DB"),
		POSTGRESDB_USERNAME: os.Getenv("POSTGRESDB_USERNAME"),
		POSTGRESDB_PASSWORD: os.Getenv("POSTGRESDB_PASSWORD"),
		REDIS_HOST:          os.Getenv("REDIS_HOST"),
		REDIS_PORT:          os.Getenv("REDIS_PORT"),
	}
}

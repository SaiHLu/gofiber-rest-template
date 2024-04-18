package common

import (
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
}

func GetEnv() *envConfig {
	DB_PORT, err := strconv.Atoi(os.Getenv("POSTGRESDB_PORT"))
	if err != nil {
		panic(err)
	}

	return &envConfig{
		// DB
		POSTGRESDB_HOST:     os.Getenv("POSTGRESDB_HOST"),
		POSTGRESDB_PORT:     DB_PORT,
		POSTGRESDB_DB:       os.Getenv("POSTGRESDB_DB"),
		POSTGRESDB_USERNAME: os.Getenv("POSTGRESDB_USERNAME"),
		POSTGRESDB_PASSWORD: os.Getenv("POSTGRESDB_PASSWORD"),
	}
}

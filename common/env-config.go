package common

import (
	"fmt"
	"os"
	"strconv"

	"github.com/SaiHLu/rest-template/common/logger"
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
	REDIS_PORT int
	REDIS_ADDR string

	// AccessToken
	ACCESS_TOKEN_SECRET string
	ACCESS_TOKEN_TTL    int

	// RefreshToken
	REFRESH_TOKEN_SECRET string
	REFRESH_TOKEN_TTL    int
}

func GetEnv() *envConfig {
	DB_PORT, err := strconv.Atoi(os.Getenv("POSTGRESDB_PORT"))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	REDIS_ADDR := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	REDIS_PORT, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	ACCESS_TOKEN_TTL, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TTL"))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	REFRESH_TOKEN_TTL, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	return &envConfig{
		// DB
		POSTGRESDB_HOST:     os.Getenv("POSTGRESDB_HOST"),
		POSTGRESDB_PORT:     DB_PORT,
		POSTGRESDB_DB:       os.Getenv("POSTGRESDB_DB"),
		POSTGRESDB_USERNAME: os.Getenv("POSTGRESDB_USERNAME"),
		POSTGRESDB_PASSWORD: os.Getenv("POSTGRESDB_PASSWORD"),
		// REDIS
		REDIS_HOST: os.Getenv("REDIS_HOST"),
		REDIS_PORT: REDIS_PORT,
		REDIS_ADDR: REDIS_ADDR,
		// ACCESS_TOKEN
		ACCESS_TOKEN_SECRET: os.Getenv("ACCESS_TOKEN_SECRET"),
		ACCESS_TOKEN_TTL:    ACCESS_TOKEN_TTL,
		// REFRESH_TOKEN
		REFRESH_TOKEN_SECRET: os.Getenv("REFRESH_TOKEN_SECRET"),
		REFRESH_TOKEN_TTL:    REFRESH_TOKEN_TTL,
	}
}

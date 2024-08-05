package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/SaiHLu/rest-template/common"
	"github.com/SaiHLu/rest-template/common/logger"
	"github.com/SaiHLu/rest-template/internal/infrastructure/cache"
	api "github.com/SaiHLu/rest-template/internal/infrastructure/interface"
	"github.com/SaiHLu/rest-template/internal/infrastructure/persistence/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis/v3"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		sig         = make(chan os.Signal, 1)
	)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	app := fiber.New(fiber.Config{
		ReadTimeout: time.Minute * 5,
		ColorScheme: fiber.Colors{
			Green: fiber.DefaultColors.Green,
		},
		AppName: "Rest Template",
	})

	go func() {
		<-sig
		cancel()

		if err := app.Shutdown(); err != nil {
			logger.Error(err.Error())
		}
	}()

	postgresDB, err := database.SetUpPostgresDatabaseConnection()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	cacheStorage := cache.NewRedisCache(redis.Config{
		Host:      common.GetEnv().REDIS_HOST,
		Port:      common.GetEnv().REDIS_PORT,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})
	if err := cacheStorage.Ping(); err != nil {
		logger.Error("Redis connection failed.")
		os.Exit(1)
	}

	defer func() {
		sqlDb, err := postgresDB.DB()
		if err != nil {
			logger.Error(err.Error())
		}

		if err := sqlDb.Close(); err != nil {
			logger.Error(err.Error())
		}

		if err := cacheStorage.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()

	api.SetupRoutes(app, postgresDB, cacheStorage)

	go func() {
		if err := app.Listen(":8000"); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	logger.Info("Exited")
}

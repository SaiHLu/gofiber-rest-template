package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/SaiHLu/rest-template/common"
	database "github.com/SaiHLu/rest-template/database"
	"github.com/SaiHLu/rest-template/internal/infrastructure/cache"
	api "github.com/SaiHLu/rest-template/internal/infrastructure/interface"
	"github.com/SaiHLu/rest-template/internal/infrastructure/queue"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis/v3"
	_ "github.com/joho/godotenv/autoload"
)

var (
	wg sync.WaitGroup
)

func main() {
	postgresDB := database.SetUpPostgresDatabaseConnection()
	newQueue := queue.NewQueue(common.GetEnv().REDIS_ADDR)
	cacheStore := cache.NewRedisClient(redis.Config{
		Host:      common.GetEnv().REDIS_HOST,
		Port:      common.GetEnv().REDIS_PORT,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})

	wg.Add(1)
	go func() {
		defer wg.Done()

		newQueue.MonitorQueues()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		newQueue.ExecuteQueue()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		app := fiber.New(fiber.Config{
			ReadTimeout: time.Second * 5,
			ColorScheme: fiber.Colors{
				Green: fiber.DefaultColors.Green,
			},
			AppName: "Rest Template",
		})

		api.SetupRoutes(app, postgresDB, cacheStore.Client)

		log.Fatalln(app.Listen(":8000"), "Server is running on port: 8000")
	}()

	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of active goroutines: %d\n", numGoroutines)

	wg.Wait()
}

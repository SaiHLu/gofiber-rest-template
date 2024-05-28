package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/SaiHLu/rest-template/common"
	database "github.com/SaiHLu/rest-template/database"
	"github.com/SaiHLu/rest-template/internal/external/api"
	"github.com/SaiHLu/rest-template/internal/external/queue"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

var (
	wg sync.WaitGroup
)

func main() {
	var postgresDB = database.SetUpPostgresDatabaseConnection()

	newQueue := queue.NewQueue(common.GetEnv().REDIS_ADDR)

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

		api.SetupRoutes(app, postgresDB)

		log.Fatalln(app.Listen(":8000"), "Server is running on port: 8000")
	}()

	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of active goroutines: %d\n", numGoroutines)

	wg.Wait()
}

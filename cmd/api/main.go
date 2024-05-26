package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/SaiHLu/rest-template/common"
	database "github.com/SaiHLu/rest-template/database"
	http "github.com/SaiHLu/rest-template/internal/external/http"
	"github.com/SaiHLu/rest-template/internal/external/queue"
	"github.com/SaiHLu/rest-template/internal/external/queue/task"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	_ "github.com/joho/godotenv/autoload"
)

var (
	wg sync.WaitGroup
)

func main() {
	var postgresDB = database.SetUpPostgresDatabaseConnection()

	RedisAddr := fmt.Sprintf("%s:%s", common.GetEnv().REDIS_HOST, common.GetEnv().REDIS_PORT)
	newQueue := queue.NewQueue(RedisAddr)
	queueClient := newQueue.SetupQueue()
	defer queueClient.Close()

	task, err := task.NewEmailDeliveryTask(1, "my:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := queueClient.Enqueue(task, asynq.Queue("low"), asynq.ProcessIn(time.Second*10))
	if err != nil {
		log.Fatalf("could not start queue: %v", err)
	}

	log.Printf("enqueued task: type=%s queue=%s\n", info.Type, info.Queue)

	wg.Add(1)
	go func() {
		defer wg.Done()

		newQueue.StartMonitoring()
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

		http.SetupRoutes(app, postgresDB)

		log.Fatalln(app.Listen(":8000"), "Server is running on port: 8000")
	}()

	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Number of active goroutines: %d\n", numGoroutines)

	wg.Wait()
}

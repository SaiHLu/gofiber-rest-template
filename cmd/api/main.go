package main

import (
	"log"
	"time"

	database "github.com/SaiHLu/rest-template/database"
	http "github.com/SaiHLu/rest-template/internal/external/http"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var postgresDB = database.SetUpPostgresDatabaseConnection()

	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * 5,
		ColorScheme: fiber.Colors{
			Green: fiber.DefaultColors.Green,
		},
		AppName: "Rest Template",
	})

	http.SetupRoutes(app, postgresDB)

	log.Fatalln(app.Listen(":8000"), "Server is running on port: 8000")
}

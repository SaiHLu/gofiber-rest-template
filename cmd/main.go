package main

import (
	"log"
	"time"

	"github.com/SaiHLu/rest-template/common"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	common.SetUpDatabaseConnection()

	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * 5,
		ColorScheme: fiber.Colors{
			Green: fiber.DefaultColors.Green,
		},
		AppName: "Rest Template",
	})

	setupRoutes(app)

	log.Fatalln(app.Listen(":8000"), "Server is running on port: 8000")
}

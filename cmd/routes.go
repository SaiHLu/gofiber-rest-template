package main

import (
	users "github.com/SaiHLu/rest-template/internal/users"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	users.SetupRoutes(app)
}

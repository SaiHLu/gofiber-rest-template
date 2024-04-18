package users

import (
	"github.com/SaiHLu/rest-template/common"
	user_handler "github.com/SaiHLu/rest-template/internal/users/handlers"
	user_model "github.com/SaiHLu/rest-template/internal/users/models"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userHandler := user_handler.NewUserHandler()

	routeV1 := app.Group("/api/v1")

	routeV1.Get("/", userHandler.Get)
	routeV1.Post("/", common.RequestValidationMiddleware[user_model.UserModel], userHandler.Create)
}

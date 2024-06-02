package route

import (
	"github.com/SaiHLu/rest-template/internal/infrastructure/auth/dto"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/controller"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App, authController *controller.AuthController) {
	authRouteGroup := app.Group("/auth")

	authRouteGroup.Post("/login", middleware.CreateRequestValidationMiddleware[dto.LoginDto], authController.Login)
}

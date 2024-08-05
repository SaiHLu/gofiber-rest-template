package route

import (
	"github.com/SaiHLu/rest-template/internal/core/dto"
	"github.com/SaiHLu/rest-template/internal/core/repository/user"
	"github.com/SaiHLu/rest-template/internal/core/service"
	"github.com/SaiHLu/rest-template/internal/infrastructure/cache"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/controller"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(app *fiber.App, postgresDB *gorm.DB, cacheStorage cache.Cache) {
	userRepository := user.NewPostgresRepository(postgresDB)
	userService := service.NewUserService(userRepository)
	authController := controller.NewAuthController(userService, cacheStorage)

	authRouteGroup := app.Group("/auth")

	authRouteGroup.Post("/login", middleware.CreateRequestValidationMiddleware[dto.LoginDto], authController.Login)
}

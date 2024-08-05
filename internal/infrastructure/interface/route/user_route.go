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

func UserRouter(app *fiber.App, postgresDB *gorm.DB, cacheStorage cache.Cache) {
	userRepository := user.NewPostgresRepository(postgresDB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	app.Use(middleware.JwtMiddleware(userService, cacheStorage))

	userRouteGroup := app.Group("/users")

	userRouteGroup.Get("/", middleware.QueryValidationMiddleware[dto.QueryUserDto], userController.GetAll)
	userRouteGroup.Get("/:id", userController.GetOne)
	userRouteGroup.Post("/", middleware.CreateRequestValidationMiddleware[dto.CreateUserDto], userController.Create)
	userRouteGroup.Patch("/:id", middleware.ParamValidationMiddleware[dto.ParamUserDto], middleware.UpdateRequestValidationMiddleware[dto.UpdateUserDto], userController.Update)
	userRouteGroup.Delete("/:id", middleware.ParamValidationMiddleware[dto.ParamUserDto], userController.Delete)

	userRouteGroup.Get("/queues", userController.ExecuteQueue)
}

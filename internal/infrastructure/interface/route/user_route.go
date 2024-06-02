package route

import (
	"github.com/SaiHLu/rest-template/internal/app/dto"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/controller"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App, userController *controller.UserController) {
	userRouteGroup := app.Group("/users")

	userRouteGroup.Get("/", middleware.QueryValidationMiddleware[dto.QueryUserDto], userController.GetAll)
	userRouteGroup.Get("/:id", userController.GetOne)
	userRouteGroup.Post("/", middleware.CreateRequestValidationMiddleware[dto.CreateUserDto], userController.Create)
	userRouteGroup.Patch("/:id", middleware.ParamValidationMiddleware[dto.ParamUserDto], middleware.UpdateRequestValidationMiddleware[dto.UpdateUserDto], userController.Update)
	userRouteGroup.Delete("/:id", middleware.ParamValidationMiddleware[dto.ParamUserDto], userController.Delete)

	userRouteGroup.Get("/queues", userController.ExecuteQueue)
}

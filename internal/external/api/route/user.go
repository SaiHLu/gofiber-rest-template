package route

import (
	"github.com/SaiHLu/rest-template/internal/app/domain/user/dto"
	postgresrepository "github.com/SaiHLu/rest-template/internal/app/domain/user/repository/postgres_repository"
	"github.com/SaiHLu/rest-template/internal/app/domain/user/service"
	"github.com/SaiHLu/rest-template/internal/external/api/controller"
	"github.com/SaiHLu/rest-template/internal/external/api/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRouter(app *fiber.App, db *gorm.DB) {
	userRepo := postgresrepository.NewPostgresRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	app.Get("/users", middleware.QueryValidationMiddleware[dto.QueryUserDto], userController.GetAll)
	app.Post("/users", middleware.CreateRequestValidationMiddleware[dto.CreateUserDto], userController.Create)
	app.Patch("/users/:id", middleware.ParamValidationMiddleware[dto.ParamUserDto], middleware.UpdateRequestValidationMiddleware[dto.UpdateUserDto], userController.Update)
	app.Delete("/users/:id", middleware.ParamValidationMiddleware[dto.ParamUserDto], userController.Delete)

	app.Get("/queues", userController.ExecuteQueue)
}

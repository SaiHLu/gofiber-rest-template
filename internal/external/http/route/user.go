package route

import (
	"github.com/SaiHLu/rest-template/internal/app/domain/user/dto"
	postgresrepository "github.com/SaiHLu/rest-template/internal/app/domain/user/repository/postgres_repository"
	"github.com/SaiHLu/rest-template/internal/app/domain/user/service"
	"github.com/SaiHLu/rest-template/internal/external/http/controller"
	"github.com/SaiHLu/rest-template/internal/external/http/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRouter(app *fiber.App, db *gorm.DB) {
	userRepo := postgresrepository.NewPostgresRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	app.Get("/users", middleware.QueryValidationMiddleware[dto.QueryUserDto], userController.GetAll)
	app.Post("/users", middleware.CreateRequestValidationMiddleware[dto.CreateUserDto], userController.Create)
}

package api

import (
	"log"

	"github.com/SaiHLu/rest-template/common"
	"github.com/SaiHLu/rest-template/internal/app/repository/user"
	"github.com/SaiHLu/rest-template/internal/app/service"
	"github.com/SaiHLu/rest-template/internal/infrastructure/cache/repository"
	cacheService "github.com/SaiHLu/rest-template/internal/infrastructure/cache/service"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/controller"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/middleware"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/route"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis/v3"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, postgresDB *gorm.DB, cacheStore *redis.Storage) {
	userRepository := user.NewPostgresRepository(postgresDB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)

	redisRepo := repository.NewRedisCacheService(cacheStore)
	redisService := cacheService.NewCacheService(redisRepo)

	route.AuthRouter(app, authController)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(common.GetEnv().JWT_SECRET)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println("err: ", err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
		},
	}))

	app.Use(middleware.JwtMiddleware(userService, redisService))

	route.UserRouter(app, userController)
}

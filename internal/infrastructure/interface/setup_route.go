package api

import (
	"github.com/SaiHLu/rest-template/common"
	"github.com/SaiHLu/rest-template/common/logger"
	"github.com/SaiHLu/rest-template/internal/core/repository/user"
	"github.com/SaiHLu/rest-template/internal/core/service"
	"github.com/SaiHLu/rest-template/internal/infrastructure/cache"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/middleware"
	"github.com/SaiHLu/rest-template/internal/infrastructure/interface/route"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, postgresDB *gorm.DB, cacheStorage cache.Cache) {
	// ======================
	// Public Routes
	// ======================
	route.AuthRouter(app, postgresDB, cacheStorage)

	// ======================
	// Public Routes
	// ======================

	// ======================
	// Private Routes
	// ======================
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(common.GetEnv().ACCESS_TOKEN_SECRET)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Error(err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
		},
	}))

	userRepository := user.NewPostgresRepository(postgresDB)
	userService := service.NewUserService(userRepository)
	app.Use(middleware.JwtMiddleware(userService, cacheStorage))

	route.UserRouter(app, postgresDB)
	// ======================
	// Private Routes
	// ======================
}

package api

import (
	"github.com/SaiHLu/rest-template/internal/external/api/route"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	route.UserRouter(app, db)
}

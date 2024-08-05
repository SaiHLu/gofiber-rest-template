package middleware

import (
	"encoding/json"
	"time"

	"github.com/SaiHLu/rest-template/common/logger"
	"github.com/SaiHLu/rest-template/internal/core/entity"
	"github.com/SaiHLu/rest-template/internal/core/service"
	"github.com/SaiHLu/rest-template/internal/infrastructure/cache"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func JwtMiddleware(userService service.UserService, cacheStorage cache.Cache) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenInfo := c.Locals("user").(*jwt.Token)
		claims := tokenInfo.Claims.(jwt.MapClaims)
		userId := claims["id"].(uuid.UUID)
		var (
			user       entity.User
			err        error
			cachedUser []byte
			expireTime *jwt.NumericDate
		)

		expireTime, err = claims.GetExpirationTime()
		if err != nil {
			logger.Error(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
		}

		cachedUser, err = cacheStorage.Get(userId.String())
		if err != nil {
			logger.Error(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
		}

		if len(cachedUser) != 0 {
			if err = json.Unmarshal(cachedUser, &user); err != nil {
				logger.Error(err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
			}

			c.Locals("authUser", user)
		} else {
			user, err = userService.GetOneById(userId)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
			}

			c.Locals("authUser", user)
			cachedUser, err = json.Marshal(user)
			if err != nil {
				logger.Error(err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
			}

			cacheStorage.Set(user.ID.String(), cachedUser, time.Until(expireTime.Time))
		}

		return c.Next()
	}
}

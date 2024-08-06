package middleware

import (
	"encoding/json"
	"time"

	"github.com/SaiHLu/rest-template/common/constant"
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
		id := claims["id"]
		userId, _ := uuid.Parse(id.(string))

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

		cachedUser, err = cacheStorage.Get(constant.GetUser(userId))
		if err != nil {
			logger.Error(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
		}

		if len(cachedUser) != 0 {
			if err = json.Unmarshal(cachedUser, &user); err != nil {
				logger.Error(err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
			}

			c.Locals(constant.AuthUserCtx, user)
		} else {
			user, err = userService.GetOneById(userId)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
			}

			c.Locals(constant.AuthUserCtx, user)
			cachedUser, err = json.Marshal(user)
			if err != nil {
				logger.Error(err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
			}

			cacheStorage.Set(constant.GetUser(userId), cachedUser, time.Until(expireTime.Time))
		}

		cacheStorage.Set(string(constant.UserIdCtx), []byte(user.ID.String()), -1)

		return c.Next()
	}
}

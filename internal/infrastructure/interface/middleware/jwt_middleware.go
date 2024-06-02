package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/SaiHLu/rest-template/internal/app/entity"
	"github.com/SaiHLu/rest-template/internal/app/service"
	cache "github.com/SaiHLu/rest-template/internal/infrastructure/cache/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(userService service.UserService, cacheService cache.CacheService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenInfo := c.Locals("user").(*jwt.Token)
		claims := tokenInfo.Claims.(jwt.MapClaims)
		userId := int64(claims["id"].(float64))
		var (
			user       entity.User
			err        error
			cachedUser []byte
			expireTime *jwt.NumericDate
		)

		expireTime, err = claims.GetExpirationTime()
		if err != nil {
			log.Println("Expire Time Error: ", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
		}

		cachedUser, err = cacheService.Get(fmt.Sprintf("%d", userId))
		if err != nil {
			log.Println("cache Error: ", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
		}

		if len(cachedUser) != 0 {
			if err = json.Unmarshal(cachedUser, &user); err != nil {
				log.Println("Json Unmarshal Error: ", err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
			}

			c.Locals("authUser", user)
		} else {
			user, err = userService.GetOne(map[string]interface{}{"id": userId})
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
			}

			c.Locals("authUser", user)
			cachedUser, err = json.Marshal(user)
			if err != nil {
				log.Println("Json Marshal Error: ", err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON("Something went wrong")
			}

			cacheService.Set(fmt.Sprintf("%d", user.ID), cachedUser, time.Until(expireTime.Time))
		}

		return c.Next()
	}
}

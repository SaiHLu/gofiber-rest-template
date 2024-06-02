package controller

import (
	"time"

	"github.com/SaiHLu/rest-template/common"
	"github.com/SaiHLu/rest-template/internal/app/domain/user/service"
	"github.com/SaiHLu/rest-template/internal/infrastructure/auth/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authService service.UserService
}

func NewAuthController(authService service.UserService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (a *AuthController) Login(c *fiber.Ctx) error {
	var body dto.LoginDto
	_ = c.BodyParser(&body)

	user, err := a.authService.GetOne(map[string]interface{}{"email": body.Email})
	if err != nil {
		return c.JSON(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.JSON(fiber.Map{"password": err.Error()})
	}

	expireTime := time.Now().Add(time.Second * time.Duration(common.GetEnv().JWT_TTL)).Unix()
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": expireTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(common.GetEnv().JWT_SECRET))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(t)
}

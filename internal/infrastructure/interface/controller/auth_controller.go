package controller

import (
	"time"

	"github.com/SaiHLu/rest-template/common"
	"github.com/SaiHLu/rest-template/common/constant"
	"github.com/SaiHLu/rest-template/common/logger"
	"github.com/SaiHLu/rest-template/common/presenter"
	"github.com/SaiHLu/rest-template/internal/core/dto"
	"github.com/SaiHLu/rest-template/internal/core/service"
	"github.com/SaiHLu/rest-template/internal/infrastructure/application/token"
	"github.com/SaiHLu/rest-template/internal/infrastructure/cache"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authService  service.UserService
	cacheStorage cache.Cache
}

func NewAuthController(authService service.UserService, cacheStorage cache.Cache) *AuthController {
	return &AuthController{
		authService:  authService,
		cacheStorage: cacheStorage,
	}
}

var (
	accessTokenTTL  = time.Second * time.Duration(common.GetEnv().ACCESS_TOKEN_TTL)
	refreshTokenTTL = time.Second * time.Duration(common.GetEnv().REFRESH_TOKEN_TTL)
)

func (a *AuthController) Login(c *fiber.Ctx) error {
	var body dto.LoginDto
	_ = c.BodyParser(&body)

	user, err := a.authService.GetOneByEmail(body.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(presenter.ErrorJsonResponse("Invalid Credentials.", err.Error()))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(presenter.ErrorJsonResponse("Invalid Credentials.", err.Error()))
	}

	accessToken, err := token.GenerateAccessToken(user.ID, time.Now().Add(accessTokenTTL).Unix())
	if err != nil {
		logger.Error(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(presenter.ErrorJsonResponse("Invalid Credentials.", err.Error()))
	}

	refreshToken, err := token.GenerateRefreshToken(user.ID, time.Now().Add(refreshTokenTTL).Unix())
	if err != nil {
		logger.Error(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(presenter.ErrorJsonResponse("Invalid Credentials.", err.Error()))
	}

	a.cacheStorage.Set(constant.GetRefreshToken(user.ID), []byte(refreshToken), refreshTokenTTL)

	return c.Status(fiber.StatusOK).JSON(presenter.SuccessJsonResponse(map[string]string{"accessToken": accessToken}, "Login Success"))
}

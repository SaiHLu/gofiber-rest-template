package controller

import (
	"github.com/SaiHLu/rest-template/internal/app/domain/user/dto"
	"github.com/SaiHLu/rest-template/internal/app/domain/user/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.Service
}

func NewUserController(userService service.Service) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) GetAll(c *fiber.Ctx) error {
	var query dto.QueryUserDto

	if err := c.QueryParser(&query); err != nil {
		return c.JSON(err)
	}

	users, err := u.userService.GetAll(query)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(users)
}

func (u *UserController) Create(c *fiber.Ctx) error {
	var body dto.CreateUserDto

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(err.Error())
	}

	user, err := u.userService.Create(body)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(user)
}

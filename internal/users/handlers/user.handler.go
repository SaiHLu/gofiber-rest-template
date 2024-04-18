package user_handler

import (
	"net/http"

	"github.com/SaiHLu/rest-template/common"
	user_model "github.com/SaiHLu/rest-template/internal/users/models"
	user_repository "github.com/SaiHLu/rest-template/internal/users/repository"
	user_service "github.com/SaiHLu/rest-template/internal/users/services"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct{}

type CustomErrorMessage struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

var (
	userRepository user_repository.UserRepository
	userService    user_service.UserService
)

func NewUserHandler() *userHandler {
	userRepository = user_repository.NewUserRepository(common.DBInstance)
	userService = user_service.NewUserService(userRepository)

	return &userHandler{}
}

func (u *userHandler) Get(ctx *fiber.Ctx) error {
	result, err := userService.Get()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return ctx.JSON(common.DefaultErrorResponse("Get users error.", err))
	}

	return ctx.JSON(common.DefaultSuccessResponse(result, "Get Users"))
}

func (u *userHandler) Create(ctx *fiber.Ctx) error {
	data := user_model.UserModel{}

	if err := ctx.BodyParser(&data); err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(common.DefaultErrorResponse("Incorrect Payload", err.Error()))
	}

	result, err := userService.Create(&data)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(common.DefaultErrorResponse("Create user error.", err.Error()))
	}

	return ctx.JSON(common.DefaultSuccessResponse(result, "Created User"))
}

package controller

import (
	"log"
	"time"

	"github.com/SaiHLu/rest-template/common"
	"github.com/SaiHLu/rest-template/internal/core/dto"
	"github.com/SaiHLu/rest-template/internal/core/service"
	"github.com/SaiHLu/rest-template/internal/infrastructure/queue"
	"github.com/SaiHLu/rest-template/internal/infrastructure/queue/task"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) GetAll(c *fiber.Ctx) error {
	var query dto.QueryUserDto

	_ = c.QueryParser(&query)

	users, err := u.userService.GetAll(query)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(users)
}

func (u *UserController) GetOne(c *fiber.Ctx) error {
	var id dto.IdParamUserDto

	_ = c.ParamsParser(&id)

	user, err := u.userService.GetOneById(id.ID)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(user)
}

func (u *UserController) Create(c *fiber.Ctx) error {
	var body dto.CreateUserDto

	_ = c.BodyParser(&body)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return c.JSON(err.Error())
	}

	body.Password = string(hashedPassword)
	user, err := u.userService.Create(body)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(user)
}

func (u *UserController) Update(c *fiber.Ctx) error {
	var body dto.UpdateUserDto
	var userId dto.ParamUserDto

	_ = c.ParamsParser(&userId)
	_ = c.BodyParser(&body)

	user, err := u.userService.Update(userId.Id, body)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(user)
}

func (u *UserController) Delete(c *fiber.Ctx) error {
	var param dto.ParamUserDto

	_ = c.ParamsParser(&param)

	user, err := u.userService.Delete(param.Id)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(user)
}

func (u *UserController) ExecuteQueue(c *fiber.Ctx) error {
	newQueue := queue.NewQueue(common.GetEnv().REDIS_ADDR)
	queueClient := newQueue.SetupQueueClient()
	defer queueClient.Close()

	log.Println("newQueue: ", &newQueue)

	task, err := task.NewEmailDeliveryTask(1, "my:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := queueClient.Enqueue(task, asynq.Queue("low"), asynq.ProcessIn(time.Second*10))
	if err != nil {
		log.Fatalf("could not start queue: %v", err)
	}

	log.Printf("enqueued task: type=%s queue=%s\n", info.Type, info.Queue)

	return c.JSON("Execute Job")
}

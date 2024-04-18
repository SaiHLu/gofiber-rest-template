package common

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CustomErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func RequestValidationMiddleware[T any](c *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	body := new(T)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(DefaultErrorResponse("Incorrect Payload", err.Error()))
	}

	if err := validate.Struct(body); err != nil {
		var validationErrors validator.ValidationErrors

		if errors.As(err, &validationErrors) {
			output := make([]CustomErrorMessage, len(validationErrors))

			for index, field := range validationErrors {
				output[index] = CustomErrorMessage{field.Field(), FormatValidationMessage(field.Tag())}
			}

			return c.Status(http.StatusBadRequest).JSON(DefaultErrorResponse("Validation Errors", output))
		}

		return c.Status(http.StatusInternalServerError).JSON(DefaultErrorResponse("Something went wrong", err))
	}

	return c.Next()
}

package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/SaiHLu/rest-template/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParamValidationMiddleware[T any](c *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	param := new(T)

	if err := c.ParamsParser(param); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.DefaultErrorResponse("Invalid Payload", err.Error()))
	}

	if err := validate.Struct(param); err != nil {
		var validationErrors validator.ValidationErrors

		if errors.As(err, &validationErrors) {
			customErrorsFormat := make(map[string]string)

			for _, field := range validationErrors {
				customErrorsFormat[strings.ToLower(field.Field())] = common.FormatValidationMessage(field.Tag(), field.Value())
			}

			return c.Status(http.StatusBadRequest).JSON(common.DefaultErrorResponse("Validation Errors", customErrorsFormat))
		}

		return c.Status(http.StatusInternalServerError).JSON(common.DefaultErrorResponse("Something went wrong", err))
	}

	return c.Next()
}

func QueryValidationMiddleware[T any](c *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	query := new(T)

	if err := c.QueryParser(query); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.DefaultErrorResponse("Invalid Payload", err.Error()))
	}

	if err := validate.Struct(query); err != nil {
		var validationErrors validator.ValidationErrors

		if errors.As(err, &validationErrors) {
			customErrorsFormat := make(map[string]string)

			for _, field := range validationErrors {
				customErrorsFormat[strings.ToLower(field.Field())] = common.FormatValidationMessage(field.Tag(), field.Value())
			}

			return c.Status(http.StatusBadRequest).JSON(common.DefaultErrorResponse("Validation Errors", customErrorsFormat))
		}

		return c.Status(http.StatusInternalServerError).JSON(common.DefaultErrorResponse("Something went wrong", err))
	}

	return c.Next()
}

func CreateRequestValidationMiddleware[T any](c *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	body := new(T)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.DefaultErrorResponse("Incorrect Payload", err.Error()))
	}

	if err := validate.Struct(body); err != nil {
		var validationErrors validator.ValidationErrors

		if errors.As(err, &validationErrors) {
			customErrorsFormat := make(map[string]string)

			for _, field := range validationErrors {
				customErrorsFormat[strings.ToLower(field.Field())] = common.FormatValidationMessage(field.Tag(), field.Value())
			}

			return c.Status(http.StatusBadRequest).JSON(common.DefaultErrorResponse("Validation Errors", customErrorsFormat))
		}

		return c.Status(http.StatusInternalServerError).JSON(common.DefaultErrorResponse("Something went wrong", err))
	}

	return c.Next()
}

func UpdateRequestValidationMiddleware[T any](c *fiber.Ctx) error {
	var validate = validator.New(validator.WithRequiredStructEnabled())

	validate.SetTagName("updatereq")

	body := new(T)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]error{"Incorrect Payload": err})
	}

	if err := validate.Struct(body); err != nil {
		var validationErrors validator.ValidationErrors
		customErrorsFormat := make(map[string]string)

		if errors.As(err, &validationErrors) {
			for _, field := range validationErrors {
				customErrorsFormat[strings.ToLower(field.Field())] = common.FormatValidationMessage(field.Tag(), field.Value())
			}

			return c.Status(http.StatusBadRequest).JSON(common.DefaultErrorResponse("validation errors", customErrorsFormat))
		}

		return c.Status(http.StatusInternalServerError).JSON(common.DefaultErrorResponse("Something went wrong", err))
	}

	return c.Next()
}

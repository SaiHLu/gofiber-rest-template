package common

import "github.com/gofiber/fiber/v2"

func DefaultSuccessResponse(data any, message string) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
		"error":   nil,
	}
}

func DefaultErrorResponse(message string, err any) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"message": message,
		"data":    nil,
		"error":   err,
	}
}

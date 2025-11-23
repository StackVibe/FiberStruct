package api

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// CreatedResponse 200 , 201
func CreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse 400 , 500
func ErrorResponse(c *fiber.Ctx, message string, code int) error {
	if code == 0 {
		code = fiber.StatusBadRequest
	}
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}

// NotFoundResponse  404
func NotFoundResponse(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Resource not found"
	}
	return ErrorResponse(c, message, fiber.StatusNotFound)
}

// UnauthorizedResponse  401
func UnauthorizedResponse(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Unauthorized"
	}
	return ErrorResponse(c, message, fiber.StatusUnauthorized)
}

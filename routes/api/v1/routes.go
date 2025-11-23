package v1

import "github.com/gofiber/fiber/v2"

func RegisterV1ApiRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	RegisterUserRoutes(api)
}

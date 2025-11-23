package v1

import (
	"Test/core/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router) {
	users := router.Group("users")
	users.Get("list", controllers.GetUsers)
}

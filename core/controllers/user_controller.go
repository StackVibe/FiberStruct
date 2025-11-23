package controllers

import (
	"Test/core/services"
	"Test/core/utilities/api"

	"github.com/gofiber/fiber/v2"
)

var userService services.UserService

func GetUsers(ctx *fiber.Ctx) error {
	users, err := userService.GetAllUsers()
	if err != nil {
		return api.ErrorResponse(ctx, "Failed to Load Users !", 500)
	}

	return api.SuccessResponse(ctx, "Users Fetched Successfully", users)
}

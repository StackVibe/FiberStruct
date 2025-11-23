package main

import (
	config2 "Test/database/config"
	"Test/database/seeders"
	v1 "Test/routes/api/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config2.ConnectDB()
	config2.MigrateModels()
	seeders.SeedAll()

	app := fiber.New(fiber.Config{})

	// API V1
	v1.RegisterV1ApiRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"hello":  "world",
			"status": 200,
		})
	})

	app.Listen(":3000")
}

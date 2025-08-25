package main

import (
	"my-fin-go/routes"

	"my-fin-go/database"
	"my-fin-go/database/migrations"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Inisialisasi Database
	database.DatabaseInit()

	// Inisialisasi Migration
	migrations.Migration()

	// Tambahkan middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello",
		})
	})

	// Inisialisasi Rute
	routes.RouteInit(app)

	app.Listen(":8081")
}

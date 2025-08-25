package controllers

import (
	"my-fin-go/database"
	"my-fin-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllFinances(c *fiber.Ctx) error {
	var finance []*models.Finance

	database.DB.Debug().Find(&finance)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Get Finances",
		"finance": finance,
	})
}

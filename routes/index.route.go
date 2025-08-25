package routes

import (
	"my-fin-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Get("/users", controllers.GetAllUsers)

	r.Get("/user/:id", controllers.GetUserByID)

	r.Get("/finances", controllers.GetAllFinances)

}

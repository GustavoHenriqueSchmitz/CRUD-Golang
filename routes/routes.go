package routes

import (
	"CRUD-Golang/controller"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	router := app.Group("/api")

	router.Post("/create", controller.Create)
}

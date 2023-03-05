package routes

import (
	"CRUD-Golang/controller"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	router := app.Group("/api")

	router.Post("/create", controller.Create)
	router.Get("/read", controller.Read)
	router.Get("/readOne/:id", controller.ReadOne)
	router.Put("/update/:id", controller.Update)
	router.Delete("/delete/:id", controller.Delete)
}

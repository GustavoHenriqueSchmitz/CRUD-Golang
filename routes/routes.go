package routes

import (
	"CRUD-Golang/controller"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	router := app.Group("/api")

	router.Post("/create", controller.Create)
	router.Get("/read", controller.Read)
	router.Get("/readOne", controller.ReadOne)
	router.Put("/update", controller.Update)
	router.Delete("/delete", controller.Delete)
}

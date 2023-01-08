package main

import (
	"CRUD-Golang/models"

	"github.com/gofiber/fiber/v2"
)

func main() {

	server := models.Server{
		App:  fiber.New(),
		Port: ":5000",
	}

	server.App.Listen(server.Port)
}

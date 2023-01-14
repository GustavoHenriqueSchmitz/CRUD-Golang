package main

import (
	"CRUD-Golang/database"
	"CRUD-Golang/models"
	"CRUD-Golang/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	server := models.Server{
		App:  fiber.New(),
		Port: ":5000",
	}

	routes.Routes(server.App)
	database.Database()

	err := server.App.Listen(server.Port)
	if err != nil {
		fmt.Println(err)
	}
}

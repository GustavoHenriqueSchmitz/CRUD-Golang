package controller

import (
	"CRUD-Golang/models"
	"CRUD-Golang/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Create(request *fiber.Ctx) error {

	person := models.Person{}
	err := request.BodyParser(&person)
	fmt.Println(person)

	if err != nil {
		return request.Status(500).JSON(models.Response{
			Data:    nil,
			Message: "Erro interno na aplicação.",
			Error:   true,
		})
	}

	err = services.CreatePerson(person)

	if err != nil {
		return request.Status(400).JSON(models.Response{
			Data:    nil,
			Message: "Falha ao criar pessoa",
			Error:   true,
		})
	}

	return request.Status(201).JSON(models.Response{
		Data:    nil,
		Message: "Pessoa criada com sucesso.",
		Error:   false,
	})
}

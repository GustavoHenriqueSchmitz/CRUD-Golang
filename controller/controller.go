package controller

import (
	"CRUD-Golang/models"
	"CRUD-Golang/services"

	"github.com/gofiber/fiber/v2"
)

func Create(request *fiber.Ctx) error {

	person := models.Person{}
	err := request.BodyParser(&person)

	if err != nil {
		return request.Status(500).JSON(models.Response{
			Data:    nil,
			Message: "Internal error in the aplication.",
			Error:   true,
		})
	}

	err = services.CreatePerson(person)

	if err != nil {
		return request.Status(400).JSON(models.Response{
			Data:    nil,
			Message: "Failed while trying to create person.",
			Error:   true,
		})
	}

	return request.Status(201).JSON(models.Response{
		Data:    nil,
		Message: "Person created with success.",
		Error:   false,
	})
}

func Read(request *fiber.Ctx) error {

	people, err := services.ReadPeople()

	if err != nil {
		return request.Status(400).JSON(models.Response{
			Data:    nil,
			Message: "Error while trying to get people.",
			Error:   true,
		})
	}

	return request.Status(200).JSON(models.Response{
		Data:    people,
		Message: "Datas got with success.",
		Error:   false,
	})
}

func ReadOne(request *fiber.Ctx) error {
	return nil
}

func Delete(request *fiber.Ctx) error {
	return nil
}

func Update(request *fiber.Ctx) error {
	return nil
}

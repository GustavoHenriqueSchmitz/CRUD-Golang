package controller

import (
	"CRUD-Golang/models"
	"CRUD-Golang/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Create(request *fiber.Ctx) error {

	person := models.Person{}
	err := request.BodyParser(&person)

	if err != nil {
		return request.Status(500).JSON(models.Response{
			Data:    nil,
			Message: "internal error in the application",
			Error:   true,
		})
	}

	err = services.CreatePerson(person)

	if err != nil {
		return request.Status(400).JSON(models.Response{
			Data:    nil,
			Message: err.Error(),
			Error:   true,
		})
	}

	return request.Status(201).JSON(models.Response{
		Data:    nil,
		Message: "person created with success",
		Error:   false,
	})
}

func Read(request *fiber.Ctx) error {

	people, err := services.ReadPeople()

	if err != nil {
		return request.Status(400).JSON(models.Response{
			Data:    nil,
			Message: err.Error(),
			Error:   true,
		})
	}

	return request.Status(200).JSON(models.Response{
		Data:    people,
		Message: "datas got with success",
		Error:   false,
	})
}

func ReadOne(request *fiber.Ctx) error {

	urlId := request.Params("id")
	personId, err := strconv.Atoi(urlId)

	if err != nil {
		return request.Status(500).JSON(models.Response{
			Data:    nil,
			Message: "internal error in the application",
			Error:   true,
		})
	}

	person, err := services.ReadOnePerson(personId)

	if err != nil {
		return request.Status(400).JSON(models.Response{
			Data:    nil,
			Message: err.Error(),
			Error:   true,
		})
	}

	return request.Status(200).JSON(models.Response{
		Data:    person,
		Message: "person got with success",
		Error:   false,
	})
}

func Delete(request *fiber.Ctx) error {

	urlId := request.Params("id")
	personId, err := strconv.Atoi(urlId)

	if err != nil {
		return request.Status(500).JSON(models.Response{
			Data:    nil,
			Message: "internal error in the application",
			Error:   true,
		})
	}

	err = services.DeletePerson(personId)

	if err != nil {
		return request.Status(400).JSON(models.Response{
			Data:    nil,
			Message: err.Error(),
			Error:   true,
		})
	}

	return request.Status(200).JSON(models.Response{
		Data:    nil,
		Message: "person deleted with success",
		Error:   false,
	})
}

func Update(request *fiber.Ctx) error {

	urlId := request.Params("id")
	personId, err := strconv.Atoi(urlId)

	if err != nil {
		return request.Status(500).JSON(models.Response{
			Data:    nil,
			Message: "internal error in the application",
			Error:   true,
		})
	}

	personUpdateData := models.Person{}
	err = request.BodyParser(&personUpdateData)

	if err != nil {
		return request.Status(500).JSON(models.Response{
			Data:    nil,
			Message: "internal error in the application",
			Error:   true,
		})
	}

	err = services.UpdatePerson(personId, personUpdateData)

	if err != nil {
		return request.Status(400).JSON(models.Response{
			Data:    nil,
			Message: err.Error(),
			Error:   true,
		})
	}

	return request.Status(200).JSON(models.Response{
		Data:    nil,
		Message: "person updated with success",
		Error:   false,
	})
}

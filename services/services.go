package services

import (
	"CRUD-Golang/database"
	"CRUD-Golang/models"
)

func CreatePerson(body models.Person) error {

	result := database.DB.Table("peoples").Create(&body)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ReadPerson() {

}

func ReadOnePerson() {

}

func DeletePerson() {

}

func UpdatePerson() {

}

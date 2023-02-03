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

func ReadPeople() ([]models.Person, error) {

	people := []models.Person{}
	results := database.DB.Table("peoples").Select("*").Scan(&people)

	if results.Error != nil {
		return nil, results.Error
	}

	return people, nil
}

func ReadOnePerson() {

}

func DeletePerson() {

}

func UpdatePerson() {

}

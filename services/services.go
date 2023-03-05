package services

import (
	"CRUD-Golang/database"
	"CRUD-Golang/models"
	"errors"
)

func CreatePerson(body models.Person) error {

	result := database.DB.Table("peoples").Create(&body)

	if result.Error != nil {
		return errors.New("error while trying to create person")
	}

	return nil
}

func ReadPeople() ([]models.Person, error) {

	people := []models.Person{}
	results := database.DB.Table("peoples").Select("*").Scan(&people)

	if results.Error != nil {
		return nil, errors.New("there was an error while trying to get the people")
	}

	return people, nil
}

func ReadOnePerson(id int) (models.Person, error) {

	person := models.Person{}
	results := database.DB.Table("peoples").Select("*").Where("id", id).Scan(&person)

	if results.Error != nil {
		return models.Person{}, errors.New("there was an error while trying to find the person")
	} else if (person == models.Person{}) {
		return models.Person{}, errors.New("the person that you want to find, doesn't exist")
	}

	return person, nil
}

func DeletePerson(id int) error {
	person := models.Person{}
	results := database.DB.Table("peoples").Select("*").Where("id", id).Scan(&person)

	if results.Error != nil {
		return errors.New("there was an error while trying to delete user")
	} else if (person == models.Person{}) {
		return errors.New("the person that you want to delete, doesn't exist")
	}

	results = database.DB.Table("peoples").Delete(&models.People{}, id)
	if results.Error != nil {
		return errors.New("there was an error while trying to delete user")
	}

	return nil
}

func UpdatePerson(personId int, personDataUpdate models.Person) error {
	person := models.Person{}
	results := database.DB.Table("peoples").Select("*").Where("id", personId).Scan(&person)

	if results.Error != nil {
		return errors.New("there was an error while trying to delete user")
	} else if (person == models.Person{}) {
		return errors.New("the person that you want to delete, doesn't exist")
	}

	results = database.DB.Table("peoples").Where("id", personId).UpdateColumns(personDataUpdate)

	if results.Error != nil {
		return errors.New("there was an error while trying to delete user")
	}

	return nil
}

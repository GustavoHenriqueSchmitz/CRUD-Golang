package services

import (
	"CRUD-Golang/models"
)

func CreatePerson(body models.Person) error {

	/*_, err := database.DB.Query(`
		insert into people(name, age, function)
		values("`, body.Name, `","`, body.Age, `","`, body.Function, `")
	`)

	if err != nil {
		return err
	}*/

	return nil
}

package database

import (
	"CRUD-Golang/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

const (
	user         = "go"
	password     = "go"
	host         = "localhost"
	port         = "5432"
	databaseName = "crud_golang"
	sslmode      = "disable"
)

func Database() {

	connection := "host=" + host + " user=" + user + " password=" + password + " dbname=" + databaseName + " port=" + port + " sslmode=" + sslmode
	DB, err = gorm.Open(postgres.Open(connection))

	go func() {
		for {
			err = DB.AutoMigrate(&models.People{})
			if err != nil {

				for {
					DB, _ = gorm.Open(postgres.Open(connection))

					err = DB.AutoMigrate(&models.People{})
					if err != nil {
						time.Sleep(8 * time.Second)
					} else {
						time.Sleep(8 * time.Second)
						break
					}
				}
			} else {
				time.Sleep(60 * time.Second)
			}
		}
	}()
}

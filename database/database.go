package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

const (
	database     = "postgres"
	user         = "go"
	password     = "go"
	host         = "localhost"
	port         = 5432
	databaseName = "crud_golang"
)

func Database() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, databaseName)
	DB, err := sql.Open(database, psqlconn)

	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			// Use ping for verify if connection is active, if the connection is inactive try to reconnect.
			err = DB.Ping()
			if err != nil {
				fmt.Println("\nDatabase connection lost!")

				for {
					// Try to reconnect.
					DB, err = sql.Open(database, psqlconn)

					// Use ping for verify if the reconnection is succeed.
					err = DB.Ping()
					if err != nil {
						fmt.Println("\nReconnection failed!")
						time.Sleep(10 * time.Second)
					} else {
						fmt.Println("\nReconnected!")
						break
					}
				}
			} else {
				time.Sleep(20 * time.Second)
			}
		}
	}()
}

package dao

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/go-gl/gl/all-core/gl"
)

const (
	USER = "postgres"
	PASSWORD = "postgres"
	DATABASE = "jug"
)

func GetConnection() (*sql.DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		USER, PASSWORD, DATABASE)

	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	return conn, err
}

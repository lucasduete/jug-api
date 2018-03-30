package dao

import (
	"fmt"
	"log"
	"database/sql"

	"gopkg.in/mgo.v2"
	_ "github.com/lib/pq"
)

const (
	USER_POSTGRES     = "postgres"
	PASSWORD_POSTGRES = "postgres"
	SERVER_MONGO      = "172.17.0.2"
	DATABASE          = "jug"
)

func GetConnectionPostgres() (*sql.DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		USER_POSTGRES, PASSWORD_POSTGRES, DATABASE)

	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	return conn, err
}

func GetConnectionMongo() (*mgo.Database, error) {
	session, err := mgo.Dial(SERVER_MONGO)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		conn := session.DB(DATABASE)
		return conn, nil
	}

}

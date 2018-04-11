package dao

import (
	"fmt"
	"log"
	"database/sql"

	"gopkg.in/mgo.v2"
	_ "github.com/lib/pq"
	"github.com/garyburd/redigo/redis"
)

const (
	userPostgres     = "postgres"
	passwordPostgres = "postgres"
	serverMongo      = "172.17.0.2"
	database         = "jug"
)

func GetConnectionPostgres() (*sql.DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		userPostgres, passwordPostgres, database)

	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	return conn, err
}

func GetConnectionMongo() (*mgo.Database, error) {
	session, err := mgo.Dial(serverMongo)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		conn := session.DB(database)
		return conn, nil
	}

}

func getConnectionRedis() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	log.Fatal(err)
	return c
}

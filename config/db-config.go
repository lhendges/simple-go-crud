package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
var server = "server-host"
var port = 5432
var user = "user"
var password = "password"
var database = "database"

func Connect() {
	// Build connection string
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", server, port, user, password, database)
	var err error
	// Create connection pool
	db, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
}

func GetDb() *sql.DB {
	return db
}

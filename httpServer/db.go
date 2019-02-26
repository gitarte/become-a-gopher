package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mydbuser"
	password = "mydbpass"
	dbname   = "mydb"
)

// InitDB - create abstraction to database
func InitDB() *sql.DB {
	PsqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Open does not actually open the connection - it just creates abstraction for use in future
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Ping worms up the connections
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

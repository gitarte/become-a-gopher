package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// InitDB - create abstraction to database
func InitDB() *sql.DB {
	host := os.Getenv("POSTGRES_HOST")         //	"localhost"
	port := os.Getenv("POSTGRES_PORT")         //	5432
	user := os.Getenv("POSTGRES_USER")         //	"mydbuser"
	password := os.Getenv("POSTGRES_PASSWORD") //	"mydbpass"
	dbname := os.Getenv("POSTGRES_DB")         //	"mydb"

	dPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Faulty DB_PORT %s", err.Error())
	}

	PsqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		dPort,
		user,
		password,
		dbname,
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

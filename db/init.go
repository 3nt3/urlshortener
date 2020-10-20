package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func connectToDB(host string, port int, username string, password string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, username)
	db, err := sql.Open("postgres", psqlInfo)
	return db, err
}

func Init() error {
	log.Printf("[ ~ ] (db/init) connecting to db...\n")

	// why???
	foo, err := connectToDB("localhost", 5432, "postgres", "postgres")
	database = foo

	if err != nil {
		log.Printf("[ - ] error connecting to database: %s\n", err.Error())
		return err
	}
	log.Printf("[ + ] (db/init) successfully connected to database!!!!!11elf!\n")

	// initialize urls table
	if err = initURLsTable(); err != nil {
		log.Printf("[ - ] (db/init) error initializing urls table: %s", err.Error())
		return err
	}

	// init users table
	if err = initUsersTable(); err != nil {
		log.Printf("[ - ] (db/init) error initializing users table: %s\n", err.Error())
		return err
	}

	return nil
}

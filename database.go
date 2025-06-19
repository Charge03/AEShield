package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var postgresDB *sql.DB

func InitDB(connStr string) error {
	var err error
	postgresDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return postgresDB.Ping()
}

func StoreFile(filename, operation string, data []byte) error {
	if postgresDB == nil {
		return fmt.Errorf("database not initialized")
	}
	_, err := postgresDB.Exec(
		"INSERT INTO files (filename, operation, data) VALUES ($1, $2, $3)",
		filename, operation, data,
	)
	if err != nil {
		log.Println("Failed to store file: ", err)
	}
	return err
}

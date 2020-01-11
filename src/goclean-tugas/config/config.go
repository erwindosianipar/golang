package config

import (
	"database/sql"
	"fmt"
	"log"

	// Call package driver for mysql
	_ "github.com/go-sql-driver/mysql"
)

// ConnectHandler is to open connection to database
func ConnectHandler() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:erwindo123@tcp(127.0.0.1:3306)/restoran")
	if err != nil {
		fmt.Println("Error: cannot connect to database.")
		log.Fatal(err.Error())
	}

	return db, nil
}

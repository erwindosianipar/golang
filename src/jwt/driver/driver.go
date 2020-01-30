package driver

import (
	"database/sql"
	"log"
	"os"

	// importing sql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// ConnectDB is used to connection
func ConnectDB() *sql.DB {
	connString := os.Getenv("MYSQL_URL")
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

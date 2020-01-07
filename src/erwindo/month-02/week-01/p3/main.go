package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:erwindo123@tcp(127.0.0.1:3306)/book_store")
	handleError(err)

	defer db.Close()

	tx, err := db.Begin()
	handleError(err)

	// insert a record into table1
	res, err := tx.Exec("INSERT INTO customers (first_name, last_name, email) VALUES(?)", "Erwindo", "Sianipar", "erwindoq@gmail.com")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// fetch the auto incremented id
	id, err := res.LastInsertId()
	handleError(err)

	// insert record into table2, referencing the first record from table1
	res, err = tx.Exec("INSERT INTO orders (order_date, ammount, customer_id) VALUES(?, ?, ?)", "2020-01-06", 200000, id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// commit the transaction
	handleError(tx.Commit())

	log.Println("Done.")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type siswa struct {
	id   int
	name string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	createData(queryValues.Get("name"))
	fmt.Println("Insert success!")
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:erwindo123@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createData(param string) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("insert into siswa (nama) values (?)", param)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Insert success!")
}

func readAllData() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, _ := db.Query("select * from siswa")

	defer rows.Close()

	var result []siswa

	for rows.Next() {
		var each = siswa{}
		var err = rows.Scan(&each.id, &each.name)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	for _, each := range result {
		fmt.Println(each.id, each.name)
	}
}

func readData(param int) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var result = siswa{}

	err = db.
		QueryRow("select * from siswa where id = (?)", param).
		Scan(&result.id, &result.name)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%v %v\n", result.id, result.name)
}

func updateData(param1 string, param2 int) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("update siswa set nama = (?) where id = (?)", param1, param2)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Update success!")
}

func deleteData(param int) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("delete from siswa where id = (?)", param)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Delete success!")
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error")
	}
}

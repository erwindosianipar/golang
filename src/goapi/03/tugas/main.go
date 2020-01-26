package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type stPeople struct {
	ID     int
	Nama   string
	Gender string
}

var people []stPeople
var peoples = stPeople{}

type stPerson struct {
	Nama   string
	Gender string
}

var person []stPerson
var persons = stPerson{}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:erwindo123@tcp(127.0.0.1:3306)/people")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	port := "8080"

	http.HandleFunc("/student", getStudentHandler)
	http.HandleFunc("/students", getStudentsHandler)

	fmt.Printf("Starting web server at http://localhost:%v/student\n", port)
	fmt.Printf("Starting web server at http://localhost:%v/students", port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func getStudentsHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getAllStudents(res, req)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(req.Method + " method is not supported"))
		return
	}
}

func getAllStudents(res http.ResponseWriter, req *http.Request) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, _ := db.Query("select * from peoples")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	people = nil
	for rows.Next() {
		var each = stPeople{}
		var err = rows.Scan(&each.ID, &each.Nama, &each.Gender)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		people = append(people, each)
	}

	json, err := json.Marshal(people)

	res.Header().Set("Content-type", "application/json")
	res.Write(json)
}

func getStudentHandler(res http.ResponseWriter, req *http.Request) {

	id := req.FormValue("id")
	ID, _ := strconv.Atoi(id)

	nama := req.FormValue("nama")
	gender := req.FormValue("gender")

	if req.Method == "POST" {
		insertStudent(res, req)
	} else if req.Method == "GET" {
		getStudentByID(ID, res)
	} else if req.Method == "PUT" {
		updateStudentByID(ID, nama, gender, res)
	} else if req.Method == "DELETE" {
		deleteStudentByID(ID, res)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(req.Method + " method is not supported"))
		return
	}
}

func insertStudent(res http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Oops, something when wrong."))
		fmt.Println(err.Error())
		return
	}

	var std stPerson
	err = json.Unmarshal(reqBody, &std)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Oops, something when wrong."))
		fmt.Println(err.Error())
		return
	}

	person = append(person, std)
	savedStudentToDatabase()
}

func savedStudentToDatabase() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	for _, each := range person {
		_, err = db.Exec("insert into peoples (nama, gender) values(?, ?)", each.Nama, each.Gender)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}

func getStudentByID(ID int, res http.ResponseWriter) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var resp = stPeople{}

	err = db.QueryRow("select id, nama, gender from peoples where id = (?) group by id", ID).Scan(&resp.ID, &resp.Nama, &resp.Gender)

	if err == sql.ErrNoRows {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("No data found"))
		return
	}

	if err != nil {
		fmt.Println("No data")
		return
	}

	people = nil
	json, err := json.Marshal(resp)

	res.Header().Set("Content-type", "application/json")
	res.Write(json)
}

func updateStudentByID(ID int, nama, gender string, res http.ResponseWriter) {
	if ID == 0 && nama == "" && gender == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Oops, something went wrong."))
		return
	}

	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("update peoples set nama = ?, gender = ? where id = ?", nama, gender, ID)

	if err == sql.ErrNoRows {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("No data found to update"))
		return
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}

func deleteStudentByID(ID int, res http.ResponseWriter) {
	if ID == 0 {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Oops, something went wrong."))
		return
	}

	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rowID := 0

	err1 := db.QueryRow("select id from peoples where id = ?", ID).Scan(&rowID)

	if err1 == sql.ErrNoRows {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("No data found to delete"))
		return
	}

	_, err = db.Exec("delete from peoples where id = ?", ID)

	if err != nil {
		log.Fatal(err.Error())
	}
}

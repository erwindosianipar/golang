package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gomux/p1/middleware"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

func getStudentByID(res http.ResponseWriter, req *http.Request) {
	pathVer := mux.Vars(req)

	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var resp = stPeople{}

	err = db.QueryRow("select id, nama, gender from peoples where id = (?) group by id", pathVer["id"]).Scan(&resp.ID, &resp.Nama, &resp.Gender)

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
	people = append(people, stPeople{ID: resp.ID, Nama: resp.Nama, Gender: resp.Gender})
	json, err := json.Marshal(people)

	res.Header().Set("Content-type", "application/json")
	res.Write(json)
}

func updateStudentByID(res http.ResponseWriter, req *http.Request) {
	pathVer := mux.Vars(req)

	ID := pathVer["id"]
	nama := pathVer["nama"]
	gender := pathVer["gender"]

	if ID == "" && nama == "" && gender == "" {
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

func deleteStudentByID(res http.ResponseWriter, req *http.Request) {
	pathVer := mux.Vars(req)

	if pathVer["id"] == "" {
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

	err1 := db.QueryRow("select id from peoples where id = ?", pathVer["id"]).Scan(&rowID)

	if err1 == sql.ErrNoRows {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("No data found to delete"))
		return
	}

	_, err = db.Exec("delete from peoples where id = ?", pathVer["id"])

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	port := "8080"

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/students", getAllStudents).Methods(http.MethodGet)
	router.HandleFunc("/student", insertStudent).Methods(http.MethodPost)
	router.HandleFunc("/student/{id}", getStudentByID).Methods(http.MethodGet)
	router.HandleFunc("/student/{id}/{nama}/{gender}", updateStudentByID).Methods(http.MethodPut)
	router.HandleFunc("/student/{id}", deleteStudentByID).Methods(http.MethodDelete)

	router.Use(middleware.Logger) //middleware all

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

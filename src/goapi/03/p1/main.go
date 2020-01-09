package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Student struct {
	ID     int
	Nama   string
	Gender string
}

var students = []Student{
	{1, "Erwindo", "M"},
	{2, "Sianipar", "M"},
}

func helloHandlerGet(res http.ResponseWriter, req *http.Request) {
	// get parsing URL -> req.FormValue()
	valName := req.FormValue("name")
	valGender := strings.ToLower(req.FormValue("gender"))

	gender := ""

	if valGender == "m" {
		gender = "laki-laki"
	} else if valGender == "f" {
		gender = "perempuan"
	} else {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid Gender"))
		return
	}

	// get request method HTTP request -> req.Method
	res.Write([]byte("HTTP method: " + req.Method + " hello " + valName + " kamu adalah seorang " + gender))
}

func helloHandlerPost(res http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Oops, something when wrong."))
		fmt.Println(err.Error())
		return
	}

	res.Write([]byte("HTTP method: " + req.Method + " isi dari body adalah: "))
	res.Write(reqBody)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		helloHandlerGet(res, req)
	} else if req.Method == "POST" {
		helloHandlerPost(res, req)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(req.Method + " method is not supported"))
		return
	}
}

func getAllStudentHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		json, err := json.Marshal(students)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Oops, something went wrong."))
			fmt.Println(err.Error())
			return
		}
		res.Header().Set("Content-type", "application/json")
		res.Write(json)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(req.Method + " method is not supported"))
		return
	}
}

func studentHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

	} else if req.Method == "GET" {
		insertStudent(res, req)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(req.Method + " method is not supported"))
		return
	}
}

func getStudentHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		insertStudent(res, req)
	} else if req.Method == "GET" {
		res.Write([]byte(req.Method + " you done requested"))
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

	var std Student
	err = json.Unmarshal(reqBody, &std)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Oops, something when wrong."))
		fmt.Println(err.Error())
		return
	}

	students = append(students, std)
}

func main() {
	port := "8080"

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/student", getStudentHandler)
	http.HandleFunc("/students", getAllStudentHandler)

	fmt.Printf("Starting web server at http://localhost:%v/hello\n", port)
	fmt.Printf("Starting web server at http://localhost:%v/student\n", port)
	fmt.Printf("Starting web server at http://localhost:%v/students", port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

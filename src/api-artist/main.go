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

type stArtist struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Debut    string `json:"debut"`
	Category string `json:"category"`
}

var artist []stArtist
var artists = stArtist{}

type stArtis struct {
	Nama     string
	Debut    string
	Category string
}

var artis []stArtis
var artiss = stArtis{}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:erwindo123@tcp(127.0.0.1:3306)/artist")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	port := "8080"

	http.HandleFunc("/artist", getArtistHandler)
	http.HandleFunc("/artists", getArtistsHandler)

	fmt.Println("Starting web server at port: ", port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func getArtistsHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getAllArtist(res, req)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(req.Method + " method is not supported"))
		return
	}
}

func getAllArtist(res http.ResponseWriter, req *http.Request) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, _ := db.Query("select * from artist")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	artist = nil
	for rows.Next() {
		var each = stArtist{}
		var err = rows.Scan(&each.ID, &each.Nama, &each.Debut, &each.Category)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		artist = append(artist, each)
	}

	json, err := json.Marshal(artist)

	res.Header().Set("Content-type", "application/json")
	res.Write(json)
}

func getArtistHandler(res http.ResponseWriter, req *http.Request) {

	id := req.FormValue("id")
	ID, _ := strconv.Atoi(id)

	if req.Method == "GET" {
		getArtistByID(ID, res)
	} else if req.Method == "POST" {
		insertArtist(res, req)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte(req.Method + " method is not supported"))
		return
	}
}

func insertArtist(res http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Oops, something when wrong."))
		fmt.Println(err.Error())
		return
	}

	var std stArtis
	err = json.Unmarshal(reqBody, &std)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Oops, something when wrong."))
		fmt.Println(err.Error())
		return
	}

	artis = append(artis, std)
	savedArtistToDatabase()
}

func savedArtistToDatabase() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	for _, each := range artis {
		_, err = db.Exec("insert into artist (nama, debut, category) values(?, ?, ?)", each.Nama, each.Debut, each.Category)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}

func getArtistByID(ID int, res http.ResponseWriter) {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var resp = stArtist{}

	err = db.QueryRow("select id, nama, debut, category from artist where id = (?) group by id", ID).Scan(&resp.ID, &resp.Nama, &resp.Debut, &resp.Category)

	if err == sql.ErrNoRows {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("No data found"))
		return
	}

	if err != nil {
		fmt.Println("No data")
		return
	}

	artist = nil
	json, err := json.Marshal(resp)

	res.Header().Set("Content-type", "application/json")
	res.Write(json)
}

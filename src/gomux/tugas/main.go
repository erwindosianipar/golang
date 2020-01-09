package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gomux/tugas/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:erwindo123@tcp(127.0.0.1:3306)/restoran")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func helloGetHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(req.Method + " helloGetHandler"))
}

type Menu struct {
	ID    int
	Nama  string
	Harga int
}

var stMenu []Menu
var slMenu = Menu{}

type Meja struct {
	ID     int
	Status string
}

var stMeja []Meja
var slMeja = Meja{}

type Pesanan struct {
	MenuID int
	Qty    int
}

type Transaksi struct {
	MejaID int
	Notes  string
	Pesan  []Pesanan
}

var stTransaksi []Transaksi
var stPesanan []Pesanan

func listMenuGetHandler(res http.ResponseWriter, req *http.Request) {
	db, err := connect()

	if err != nil {
		fmt.Println("[listMenuGetHandler] Error: when connect to database:", err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("select * from menu")

	if err != nil {
		fmt.Println("[listMenuGetHandler] Error: when query select menu:", err.Error())
		return
	}

	defer rows.Close()

	stMenu = nil

	for rows.Next() {
		var each = Menu{}
		var err = rows.Scan(&each.ID, &each.Nama, &each.Harga)

		if err != nil {
			fmt.Println("[listMenuGetHandler] Error: when scaning rows from table menu:", err.Error())
			return
		}

		stMenu = append(stMenu, each)
	}

	json, err := json.Marshal(stMenu)

	if err != nil {
		fmt.Println("[listMenuGetHandler] Error: when marshal stMenu to json: ", err.Error())
		return
	}

	res.Header().Set("Content-type", "application/json")
	res.Write(json)
}

func listMejaGetHandler(res http.ResponseWriter, req *http.Request) {
	db, err := connect()

	if err != nil {
		fmt.Println("[listMejaGetHandler] Error: when connect to database:", err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("select * from meja")

	if err != nil {
		fmt.Println("[listMejaGetHandler] Error: when query select meja:", err.Error())
		return
	}

	defer rows.Close()

	stMeja = nil

	for rows.Next() {
		var each = Meja{}
		var err = rows.Scan(&each.ID, &each.Status)

		if err != nil {
			fmt.Println("[listMejaGetHandler] Error: when scaning rows from table meja:", err.Error())
			return
		}

		stMeja = append(stMeja, each)
	}

	json, err := json.Marshal(stMeja)

	if err != nil {
		fmt.Println("[listMejaGetHandler] Error: when marshal stMeja to json: ", err.Error())
		return
	}

	res.Header().Set("Content-type", "application/json")
	res.Write(json)
}

func openMejaPutHandler(res http.ResponseWriter, req *http.Request) {
	db, err := connect()

	if err != nil {
		fmt.Println("[openMejaPutHandler] Error: when connect to database:", err.Error())
		return
	}

	defer db.Close()

	pathVar := mux.Vars(req)
	status := ""
	err = db.
		QueryRow("select status from meja where id = ?", pathVar["id"]).
		Scan(&status)

	if err != nil {
		fmt.Println("[openMejaPutHandler] Error: when select status meja:", err.Error())
		return
	}

	if status == "open" {
		res.Write([]byte("Oops, meja yang anda pesan tidak tersedia"))
		return
	}

	_, err = db.Exec("update meja set status = 'open' where id = ?", pathVar["id"])

	if err == sql.ErrNoRows {
		res.WriteHeader(http.StatusForbidden)
		res.Write([]byte("Oops, id meja : " + pathVar["id"] + " tidak tersedia untuk diupdate"))
		return
	}

	if err != nil {
		fmt.Println("[openMejaPutHandler] Error: something went wrong:", err.Error())
		return
	}

	res.Write([]byte("Sukses: meja berhasil dipesan"))
}

func insertTransaksi(res http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println("[insertTransaksi] Error: when ioutil reading", err.Error())
		return
	}

	var std []Transaksi
	err = json.Unmarshal(reqBody, &std)

	if err != nil {
		fmt.Println("[insertTransaksi] Error: when unmarshal stTransaksi", err.Error())
		return
	}

	insertToDB(std[0].MejaID, std[0].Notes, std)
}

func insertToDB(ID int, notes string, std []Transaksi) {
	db, err := connect()

	if err != nil {
		fmt.Println("[insertToDB] Error: when connect to database", err.Error())
		return
	}

	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		fmt.Println("[insertToDB] Error: when begin database", err.Error())
		return
	}

	currentTime := time.Now()

	res, err := tx.Exec("insert into transaksi (meja_id, tanggal, notes) values (?, ?, ?)", ID, currentTime.Format("2006-01-02 15:04:05"), notes)

	if err != nil {
		tx.Rollback()
		fmt.Println("[insertToDB] Error: when get insert transaction to database", err.Error())
		return
	}

	id, err := res.LastInsertId()

	if err != nil {
		tx.Rollback()
		fmt.Println("[insertToDB] Error: when get lastInsertId", err.Error())
		return
	}

	for i := 0; i < len(std[0].Pesan); i++ {
		_, err := tx.Exec("insert into pesanan (transaksi_id, menu_id, qty) values (?, ?, ?)", id, std[0].Pesan[i].MenuID, std[0].Pesan[i].Qty)

		if err != nil {
			tx.Rollback()
			fmt.Println("[insertToDB] Error: when insert into pesanan", err.Error())
			return
		}
	}

	err = tx.Commit()

	if err != nil {
		fmt.Println("[insertToDB] Error: when get commit database", err.Error())
		return
	}
}

func main() {
	port := "8080"

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", helloGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/listMenu", listMenuGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/listMeja", listMejaGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/openMeja/{id}", openMejaPutHandler).Methods(http.MethodPut)
	router.HandleFunc("/insertTransaksi", insertTransaksi).Methods(http.MethodPost)

	router.Use(middleware.Logger)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

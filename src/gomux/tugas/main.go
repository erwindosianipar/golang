package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gomux/tugas/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

type MenuOrdered struct {
	Nama  string
	Qty   int
	Harga int
	Total int
}

type Bill struct {
	MejaID     int
	Menus      []MenuOrdered
	GrandTotal int
}

var stMenuOrdered []MenuOrdered
var stBill []Bill

type Error struct {
	Error   bool
	Message string
	Data    interface{}
}

var jsonError Error

func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.Header().Set("Content-Type", "application/json")
	responseJSON(w, status, error)
}

func responseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func helloGetHandler(res http.ResponseWriter, req *http.Request) {
	jsonError.Error = false
	jsonError.Message = "Sukses"
	jsonError.Data = "Sukses merender /hello"

	respondWithError(res, http.StatusInternalServerError, jsonError)
}

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

	// json, err := json.Marshal(stMeja)

	if err != nil {
		fmt.Println("[listMejaGetHandler] Error: when marshal stMeja to json: ", err.Error())
		return
	}

	// res.Header().Set("Content-type", "application/json")
	// res.Write(json)

	jsonError.Error = false
	jsonError.Message = "Sukses"
	jsonError.Data = stMeja

	respondWithError(res, http.StatusInternalServerError, jsonError)
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

	jsonError.Error = false
	jsonError.Message = "Sukses: meja berhasil dipesan"
	jsonError.Data = nil

	respondWithError(res, http.StatusInternalServerError, jsonError)
	// res.Write([]byte("Sukses: meja berhasil dipesan"))
}

func insertTransaksiPostHandler(res http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println("[insertTransaksiPostHandler] Error: when ioutil reading", err.Error())
		return
	}

	var std []Transaksi
	err = json.Unmarshal(reqBody, &std)

	if err != nil {
		fmt.Println("[insertTransaksiPostHandler] Error: when unmarshal stTransaksi", err.Error())
		return
	}

	insertToDB(std[0].MejaID, std[0].Notes, std[0].Pesan)

	jsonError.Error = false
	jsonError.Message = "Sukses: Pesanan berhasil dibuat"
	jsonError.Data = nil

	respondWithError(res, http.StatusInternalServerError, jsonError)
}

func insertToDB(ID int, notes string, pesanan []Pesanan) {
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
		fmt.Println("[insertToDB] Error: when insert transaction to database", err.Error())
		return
	}

	id, err := res.LastInsertId()

	if err != nil {
		tx.Rollback()
		fmt.Println("[insertToDB] Error: when get lastInsertId", err.Error())
		return
	}

	for i := 0; i < len(pesanan); i++ {
		_, err := tx.Exec("insert into pesanan (transaksi_id, menu_id, qty) values (?, ?, ?)", id, pesanan[i].MenuID, pesanan[i].Qty)

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

	return
}

func closeMejaGetHandler(res http.ResponseWriter, req *http.Request) {
	db, err := connect()

	if err != nil {
		fmt.Println("[closeMejaGetHandler] Error: when connect to database:", err.Error())
		return
	}

	defer db.Close()

	pathVar := mux.Vars(req)
	ID, _ := strconv.Atoi(pathVar["id"])

	rows, err := db.Query("select nama, qty, harga, qty*harga as total from pesanan ps join menu mn join transaksi tr on ps.menu_id = mn.id and ps.transaksi_id = tr.id where tr.id = ?", ID)

	if err != nil {
		fmt.Println("[closeMejaGetHandler] Error: when query select:", err.Error())
		return
	}

	_, err = db.Exec("update meja set status = 'close' where id = ?", ID)
	if err != nil {
		fmt.Println("[closeMejaGetHandler] Error: when query update status meja:", err.Error())
		return
	}

	defer rows.Close()

	stMenuOrdered = nil
	stBill = nil

	for rows.Next() {
		var each = MenuOrdered{}
		var err = rows.Scan(&each.Nama, &each.Qty, &each.Harga, &each.Total)

		if err != nil {
			fmt.Println("[listMejaGetHandler] Error: when scaning rows from table meja:", err.Error())
			return
		}

		stMenuOrdered = append(stMenuOrdered, each)
	}

	grandTotal := 0

	db.QueryRow("select sum(qty*harga) as grandTotal from pesanan ps join menu mn join transaksi tr on ps.menu_id = mn.id and ps.transaksi_id = tr.id where tr.id = ?", ID).Scan(&grandTotal)

	stBill = append(stBill, Bill{MejaID: ID, Menus: stMenuOrdered, GrandTotal: grandTotal})

	// json, err := json.Marshal(stBill)

	if err != nil {
		fmt.Println("[listMejaGetHandler] Error: when marshal stBill to json: ", err.Error())
		return
	}

	// res.Header().Set("Content-type", "application/json")
	// res.Write(json)

	jsonError.Error = false
	jsonError.Message = "Success: Transaksi selesai"
	jsonError.Data = stBill

	respondWithError(res, http.StatusInternalServerError, jsonError)

}

func main() {
	port := "8080"

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", helloGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/listMenu", listMenuGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/listMeja", listMejaGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/openMeja/{id}", openMejaPutHandler).Methods(http.MethodPut)
	router.HandleFunc("/insertTransaksi", insertTransaksiPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/closeMeja/{id}", closeMejaGetHandler).Methods(http.MethodGet)

	fmt.Println("Web server starting at port:", port)

	router.Use(middleware.Logger)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

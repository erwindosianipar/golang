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

// Menu struct adalah xxxx
type Menu struct {
	ID    int
	Nama  string
	Harga int
}

var stMenu []Menu
var slMenu = Menu{}

// Meja adalah aaaaaa
type Meja struct {
	ID     int
	Status string
}

var stMeja []Meja
var slMeja = Meja{}

// Pesanan adalah aaaaa
type Pesanan struct {
	MenuID int
	Qty    int
}

// Transaksi adalah aaaaa
type Transaksi struct {
	MejaID int
	Notes  string
	Pesan  []Pesanan
}

var stTransaksi []Transaksi
var stPesanan []Pesanan

// MenuOrdered adalah
type MenuOrdered struct {
	Nama  string
	Qty   int
	Harga int
	Total int
}

// Bill adalah aaaa
type Bill struct {
	MejaID     int
	Menus      []MenuOrdered
	GrandTotal int
}

var stMenuOrdered []MenuOrdered
var stBill []Bill

// Error adalah aaaa
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

	respondWithError(res, http.StatusOK, jsonError)
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

	//json, err := json.Marshal(stMenu)

	if err != nil {
		fmt.Println("[listMenuGetHandler] Error: when marshal stMenu to json: ", err.Error())
		return
	}

	// res.Header().Set("Content-type", "application/json")
	// res.Write(json)

	jsonError.Error = false
	jsonError.Message = "Sukses"
	jsonError.Data = stMenu

	respondWithError(res, http.StatusOK, jsonError)
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

	respondWithError(res, http.StatusOK, jsonError)
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

	if err == sql.ErrNoRows {
		jsonError.Error = true
		jsonError.Message = "Error: nomor meja yang anda pesan tidak tersedia"
		jsonError.Data = nil

		respondWithError(res, http.StatusInternalServerError, jsonError)
		return
	}

	if err != nil {
		fmt.Println("[openMejaPutHandler] Error: when select status meja:", err.Error())
		return
	}

	if status == "open" {
		jsonError.Error = true
		jsonError.Message = "Error: meja telah dipesan sebelumnya"
		jsonError.Data = nil

		respondWithError(res, http.StatusInternalServerError, jsonError)
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

	respondWithError(res, http.StatusOK, jsonError)
	// res.Write([]byte("Sukses: meja berhasil dipesan"))
}

func insertTransaksiPostHandler(res http.ResponseWriter, req *http.Request) {
	db, err := connect()

	if err != nil {
		fmt.Println("[insertToDB] Error: when connect to database", err.Error())
		return
	}

	defer db.Close()

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

	meja := 0
	_ = db.QueryRow("select count(*) as ada from meja where id = ?", std[0].MejaID).Scan(&meja)

	if meja < 1 {
		jsonError.Error = true
		jsonError.Message = "Error: MejaID tidak ditemukan"
		jsonError.Data = nil

		respondWithError(res, http.StatusOK, jsonError)
		return
	}

	insertToDB(res, std[0].MejaID, std[0].Notes, std[0].Pesan)
}

func cekAdaMenuID(menuID int) bool {
	db, err := connect()

	if err != nil {
		fmt.Println("[insertToDB] Error: when connect to database", err.Error())
		return false
	}

	defer db.Close()

	id := 0
	err = db.QueryRow("select id from menu where id = ?", menuID).Scan(id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

func insertToDB(resp http.ResponseWriter, MejaID int, notes string, pesanan []Pesanan) {
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

	res, err := tx.Exec("insert into transaksi (meja_id, tanggal, notes) values (?, ?, ?)", MejaID, currentTime.Format("2006-01-02 15:04:05"), notes)

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

	for i := 0; i < len(pesanan); i++ {
		adaMenu := cekAdaMenuID(pesanan[i].MenuID)

		if adaMenu == false {
			jsonError.Error = true
			jsonError.Message = "Error: terdapat MenuID yang tidak terdaftar"
			jsonError.Data = nil

			respondWithError(resp, http.StatusInternalServerError, jsonError)
			return
		}

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
}

func closeTransaksiGetHandler(res http.ResponseWriter, req *http.Request) {
	db, err := connect()

	if err != nil {
		fmt.Println("[closeTransaksiGetHandler] Error: when connect to database:", err.Error())
		return
	}

	defer db.Close()

	pathVar := mux.Vars(req)
	ID, _ := strconv.Atoi(pathVar["id"])

	id := 0
	getIDTrx := db.QueryRow("select id from transaksi where id = ?", ID).Scan(&id)

	if getIDTrx == sql.ErrNoRows {
		jsonError.Error = true
		jsonError.Message = "Error: ID transaksi tidak ditemukan"
		jsonError.Data = nil

		respondWithError(res, http.StatusOK, jsonError)
		return
	}

	rows, err := db.Query("select nama, qty, harga, qty*harga as total from pesanan ps join menu mn join transaksi tr on ps.menu_id = mn.id and ps.transaksi_id = tr.id where tr.id = ?", ID)

	if err != nil {
		fmt.Println("[closeTransaksiGetHandler] Error: when query select:", err.Error())
		return
	}

	_, err = db.Exec("update meja set status = 'close' where id = ?", ID)
	if err != nil {
		fmt.Println("[closeTransaksiGetHandler] Error: when query update status meja:", err.Error())
		return
	}

	defer rows.Close()

	stMenuOrdered = nil
	stBill = nil

	for rows.Next() {
		var each = MenuOrdered{}
		var err = rows.Scan(&each.Nama, &each.Qty, &each.Harga, &each.Total)

		if err != nil {
			fmt.Println("[closeTransaksiGetHandler] Error: when scaning rows from table meja:", err.Error())
			return
		}

		stMenuOrdered = append(stMenuOrdered, each)
	}

	grandTotal := 0

	err = db.
		QueryRow("select sum(qty*harga) as grandTotal from pesanan ps join menu mn join transaksi tr on ps.menu_id = mn.id and ps.transaksi_id = tr.id where tr.id = ?", ID).
		Scan(&grandTotal)

	stBill = append(stBill, Bill{MejaID: ID, Menus: stMenuOrdered, GrandTotal: grandTotal})

	// json, err := json.Marshal(stBill)

	if err != nil {
		fmt.Println("[closeTransaksiGetHandler] Error: when marshal stBill to json: ", err.Error())
		return
	}

	// res.Header().Set("Content-type", "application/json")
	// res.Write(json)

	jsonError.Error = false
	jsonError.Message = "Success: Transaksi selesai"
	jsonError.Data = stBill

	respondWithError(res, http.StatusOK, jsonError)
	return
}

func main() {
	port := "8080"

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", helloGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/listMenu", listMenuGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/listMeja", listMejaGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/openMeja/{id}", openMejaPutHandler).Methods(http.MethodPut)
	router.HandleFunc("/insertTransaksi", insertTransaksiPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/closeTransaksi/{id}", closeTransaksiGetHandler).Methods(http.MethodGet)

	fmt.Printf("Web server started at http://localhost:%v/hello", port)

	router.Use(middleware.Logger)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

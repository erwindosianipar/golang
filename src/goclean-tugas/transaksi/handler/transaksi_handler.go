package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"goclean-tugas/config"
	"goclean-tugas/library"
	"goclean-tugas/model"
)

// TransaksiHandler is representation to handle menu
type TransaksiHandler struct {
}

func (m *TransaksiHandler) createTransaksi(resp http.ResponseWriter, req *http.Request) {
	db, err := config.ConnectHandler()

	if err != nil {
		fmt.Println("[TransaksiHandler.createTransaksi] /transaksi/handler/ Error: when open connection to database:", err)
		return
	}

	defer db.Close()

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println("[TransaksiHandler.createTransaksi] Error: when ioutil reading", err)

		message := "Error: There is something wrong with your request body."
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	var std []model.Transaksi
	err = json.Unmarshal(reqBody, &std)

	if err != nil {
		fmt.Println("[TransaksiHandler.createTransaksi] Error: when unmarshal reqBody", err)

		message := "Error: There is something with your request body."
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	meja := 0
	_ = db.QueryRow("select count(*) as ada from meja where id = ?", std[0].MejaID).Scan(&meja)

	if meja < 1 {
		message := "Error: MejaID not exist in our database."
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	insertToDB(resp, req, std[0].MejaID, std[0].Notes, std[0].Pesan)
}

func insertToDB(resp http.ResponseWriter, req *http.Request, MejaID int, notes string, pesanan []model.Pesanan) {
	db, err := config.ConnectHandler()

	if err != nil {
		fmt.Println("[TransaksiHandler.insertToDB] /transaksi/handler/ Error: when open connection to database:", err)
		return
	}

	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		fmt.Println("[TransaksiHandler.insertToDB] Error: when begin database", err)
		return
	}

	res, err := tx.Exec("insert into transaksi (meja_id, tanggal, notes) values (?, CURRENT_TIMESTAMP, ?)", MejaID, notes)

	if err != nil {
		tx.Rollback()
		fmt.Println("[TransaksiHandler.insertToDB] Error: when get insert transaction to database", err)
		return
	}

	id, err := res.LastInsertId()

	if err != nil {
		tx.Rollback()
		fmt.Println("[TransaksiHandler.insertToDB] Error: when get lastInsertId", err)
		return
	}

	for i := 0; i < len(pesanan); i++ {
		adaMenu := cekAdaMenuID(pesanan[i].MenuID)

		if adaMenu == false {
			message := "Error: There are MenuID is not in our list menus."
			library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
			return
		}

		_, err := tx.Exec("insert into pesanan (transaksi_id, menu_id, qty) values (?, ?, ?)", id, pesanan[i].MenuID, pesanan[i].Qty)

		if err != nil {
			tx.Rollback()
			fmt.Println("[TransaksiHandler.insertToDB] Error: when insert into pesanan", err)
			return
		}
	}

	err = tx.Commit()

	if err != nil {
		fmt.Println("[TransaksiHandler.insertToDB] Error: when get commit database", err)
		return
	}

	message := "Success: Transaksi has been saved."
	library.JSONoutput(true, message, nil, resp, req, http.StatusOK)

	return
}

func cekAdaMenuID(menuID int) bool {
	db, err := config.ConnectHandler()

	if err != nil {
		fmt.Println("[TransaksiHandler.cekAdaMenuID] /transaksi/handler/ Error: when open connection to database:", err)
		return false
	}

	defer db.Close()

	rowID := 0
	err = db.QueryRow("select id from menu where id = ?", menuID).Scan(&rowID)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

func (m *TransaksiHandler) createBilling(resp http.ResponseWriter, req *http.Request) {
	db, err := config.ConnectHandler()

	if err != nil {
		fmt.Println("[TransaksiHandler.createBilling] /transaksi/handler/ Error: when open connection to database:", err)
		return
	}

	defer db.Close()

	pathVar := mux.Vars(req)
	bilingID, err := strconv.Atoi(pathVar["id"])

	if err != nil {
		message := "Error: Parameter for bilingID must be a number."
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	id := 0
	mejaID := 0
	getIDTrx := db.
		QueryRow("select id, meja_id as mejaID from transaksi where id = ?", bilingID).
		Scan(&id, &mejaID)

	if getIDTrx == sql.ErrNoRows {
		message := "Error: transaksiID was not found in our database."
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	query := "select nama, qty, harga, qty*harga as total from pesanan ps join menu mn join transaksi tr on ps.menu_id = mn.id and ps.transaksi_id = tr.id where tr.id = ?"

	rows, err := db.Query(query, bilingID)

	if err != nil {
		fmt.Println("[TransaksiHandler.createBilling] Error: when query select from transaksi:", err)
		return
	}

	defer rows.Close()

	_, err = db.Exec("update meja set status = 'close' where id = ?", mejaID)

	if err != nil {
		fmt.Println("[closeTransaksiGetHandler] Error: when query update status meja:", err)
		return
	}

	var stMenuOrdered []model.MenuOrdered
	var stBill []model.Bill

	stMenuOrdered = nil
	stBill = nil

	for rows.Next() {
		var each = model.MenuOrdered{}
		var err = rows.Scan(&each.Nama, &each.Qty, &each.Harga, &each.Total)

		if err != nil {
			fmt.Println("[closeTransaksiGetHandler] Error: when scaning rows from table meja:", err.Error())
			return
		}

		stMenuOrdered = append(stMenuOrdered, each)
	}

	grandTotal := 0

	err = db.
		QueryRow("select sum(qty*harga) as grandTotal from pesanan ps join menu mn join transaksi tr on ps.menu_id = mn.id and ps.transaksi_id = tr.id where tr.id = ?", bilingID).
		Scan(&grandTotal)

	if err != nil {
		fmt.Println("[closeTransaksiGetHandler] Error: when select grandTotal: ", err.Error())
		return
	}

	stBill = append(stBill, model.Bill{MejaID: mejaID, Menus: stMenuOrdered, GrandTotal: grandTotal})

	message := "Success: Billing has been print out and table has ben closed"
	library.JSONoutput(false, message, stBill, resp, req, http.StatusOK)
}

// CreateTransaksiHandler is handling for transaksi
func CreateTransaksiHandler(r *mux.Router) {

	TransaksiHandler := TransaksiHandler{}

	r.HandleFunc("/transaksi", TransaksiHandler.createTransaksi).Methods(http.MethodPost)
	r.HandleFunc("/transaksi/{id}", TransaksiHandler.createBilling).Methods(http.MethodPut)
}

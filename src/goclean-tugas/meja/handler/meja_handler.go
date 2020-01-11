package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"goclean-tugas/config"
	"goclean-tugas/library"
	"goclean-tugas/model"
)

// MejaHandler is a handle function for request to meja
type MejaHandler struct {
}

func (m *MejaHandler) getAllMeja(resp http.ResponseWriter, req *http.Request) {
	db, err := config.ConnectHandler()

	if err != nil {
		fmt.Println("[MejaHandler.getAllMeja] /meja/handler/ Error: when open connection to database:", err)
		return
	}

	defer db.Close()

	rows, err := db.Query("select * from meja")

	if err != nil {
		fmt.Println("[MejaHandler.getAllMeja] /meja/handler/ Error: when query select from meja:", err)
	}

	defer rows.Close()

	stMeja := []model.Meja{}

	for rows.Next() {
		var each = model.Meja{}
		var err = rows.Scan(&each.ID, &each.Status)

		if err != nil {
			fmt.Println("[MejaHandler.getAllMeja] /meja/handler/ Error: when scaning rows from table meja:", err)
			return
		}

		stMeja = append(stMeja, each)
	}

	library.JSONoutput(true, "Success", stMeja, resp, req, http.StatusOK)
}

func (m *MejaHandler) openMeja(resp http.ResponseWriter, req *http.Request) {
	db, err := config.ConnectHandler()

	if err != nil {
		fmt.Println("[MejaHandler.openMeja] /meja/handler/ Error: when open connection to database:", err)
		return
	}

	defer db.Close()

	param := mux.Vars(req)
	mejaID, err := strconv.Atoi(param["id"])

	if err != nil {
		message := "Error: Parameter for MejaID must be a number."
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	totalMeja := 0
	statusMeja := ""

	err = db.
		QueryRow("select count(*) as totalMeja, status as statusMeja from meja where id = ? group by id", mejaID).
		Scan(&totalMeja, &statusMeja)

	if err != nil {
		fmt.Println("[MejaHandler.openMeja] /meja/handler/ Error: when select count meja:", err)

		message := "Error: no list in our database meja with mejaID: " + param["id"]
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	if totalMeja < 1 {
		message := "Error: Meja with mejaID: " + param["id"] + " is not exist."
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	if statusMeja == "open" {
		message := "Error: Meja has been opened by someone, please try again with diferent mejaID"
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("update meja set status = 'open' where id = ?", mejaID)

	if err != nil {
		fmt.Println("[MejaHandler.openMeja] /meja/handler/ Error: when update meja status:", err)

		message := "Error: Oops, something went wrong, please try again."
		library.JSONoutput(false, message, nil, resp, req, http.StatusInternalServerError)
		return
	}

	message := "Success: Meja success to opened now you can create a purchase order."
	library.JSONoutput(true, message, nil, resp, req, http.StatusOK)
}

// CreateMejaHandler represent to handling request for meja
func CreateMejaHandler(r *mux.Router) {

	mejaHandler := MejaHandler{}

	r.HandleFunc("/meja/list", mejaHandler.getAllMeja).Methods(http.MethodGet)
	r.HandleFunc("/meja/{id}", mejaHandler.openMeja).Methods(http.MethodPut)
}

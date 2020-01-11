package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"goclean-tugas/config"
	"goclean-tugas/library"
	"goclean-tugas/model"
)

// MenuHandler is representation to handle menu
type MenuHandler struct {
}

func (m *MenuHandler) getAllMenu(resp http.ResponseWriter, req *http.Request) {
	db, err := config.ConnectHandler()

	if err != nil {
		fmt.Println("[MenuHandler.getAllMenu] /menu/handler/ Error: when open connection to database:", err)
		return
	}

	defer db.Close()

	rows, err := db.Query("select * from menu")

	if err != nil {
		fmt.Println("[MenuHandler.getAllMenu] /menu/handler/ Error: when query select from menu:", err)
	}

	defer rows.Close()

	stMenu := []model.Menu{}

	for rows.Next() {
		var each = model.Menu{}
		var err = rows.Scan(&each.ID, &each.Nama, &each.Harga)

		if err != nil {
			fmt.Println("[MenuHandler.getAllMenu] /menu/handler/ Error: when scaning rows from table menu:", err)
			return
		}

		stMenu = append(stMenu, each)
	}

	library.JSONoutput(true, "Success", stMenu, resp, req, http.StatusOK)
}

// CreateMenuHandler is handling for menu
func CreateMenuHandler(r *mux.Router) {

	menuHandler := MenuHandler{}

	r.HandleFunc("/menu/list", menuHandler.getAllMenu).Methods(http.MethodGet)
}

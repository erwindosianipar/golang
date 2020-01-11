package handler

import "github.com/gorilla/mux"

import "net/http"

type MenuHandler struct {
}

func (m *MenuHandler) getAllMenu(resp http.ResponseWriter, req *http.Request) {

}

func (m *MenuHandler) getMenuByID(resp http.ResponseWriter, req *http.Request) {

}

func (m *MenuHandler) deleteMenu(resp http.ResponseWriter, req *http.Request) {

}

func (m *MenuHandler) insertMenu(resp http.ResponseWriter, req *http.Request) {

}

func CreateMenuHandler(r *mux.Router) {

	menu := MenuHandler{}

	r.HandleFunc("/menu", menu.getAllMenu).Methods(http.MethodGet)
	r.HandleFunc("/menu/{id}", menu.getMenuByID).Methods(http.MethodGet)
	r.HandleFunc("/menu/{id}", menu.deleteMenu).Methods(http.MethodDelete)
	r.HandleFunc("/menu", menu.insertMenu).Methods(http.MethodPost)

}

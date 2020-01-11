package handler

import (
	"encoding/json"
	"fmt"
	"goclean/model"
	"net/http"

	"github.com/gorilla/mux"
)

// MejaHandler is isxxxxx
type MejaHandler struct {
}

func (m *MejaHandler) getAllMeja(res http.ResponseWriter, req *http.Request) {

}

func (m *MejaHandler) getMejaByID(res http.ResponseWriter, req *http.Request) {
	meja := model.Meja{
		ID:    1,
		Nomor: 1,
	}

	mejaJSON, err := json.Marshal(meja)

	if err != nil {
		fmt.Println("Error")
	}

	res.Write([]byte(mejaJSON))
}

func (m *MejaHandler) deleteMeja(res http.ResponseWriter, req *http.Request) {

}

func (m *MejaHandler) insertMeja(res http.ResponseWriter, req *http.Request) {

}

// CreateMejaHandler is aaaaaa
func CreateMejaHandler(r *mux.Router) {

	MejaHandler := MejaHandler{}

	r.HandleFunc("/tables", MejaHandler.getAllMeja).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", MejaHandler.getMejaByID).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", MejaHandler.deleteMeja).Methods(http.MethodDelete)
	r.HandleFunc("/table", MejaHandler.insertMeja).Methods(http.MethodPost)
}

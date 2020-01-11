package handler

import (
	"encoding/json"
	"fmt"
	"goclean-joe/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// MejaHandler represent xxxxxxx
type MejaHandler struct {
}

func (m *MejaHandler) getAllMeja(resp http.ResponseWriter, req *http.Request) {

}

func (m *MejaHandler) getMejaByID(resp http.ResponseWriter, req *http.Request) {

	muxVar := mux.Vars(req)

	strID := muxVar["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		data := model.ResponseWrapper{
			Success: false,
			Message: "ID Harus angka",
		}

		mejaJSON, err := json.Marshal(data)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte("Ooops, Something Went Wrong"))
			fmt.Printf("[MejaHandler.getMejaByID] Error when do json Marshalling for error handling : %v \n", err)
		}
		resp.Header().Set("Content-Type", "application/json")
		resp.Write(mejaJSON)
		return
	}

	meja := model.Meja{
		ID:     id,
		Status: "Open",
	}

	data := model.ResponseWrapper{
		Success: true,
		Message: "SUCCESS",
		Data:    meja,
	}

	mejaJSON, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[MejaHandler.getMejaByID] Error when do json Marshalling for meja : %v \n", err)
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(mejaJSON)
}

func (m *MejaHandler) deleteMeja(resp http.ResponseWriter, req *http.Request) {

}

func (m *MejaHandler) insertMeja(resp http.ResponseWriter, req *http.Request) {

}

// CreateMejaHandler represent aaaaaaa
func CreateMejaHandler(r *mux.Router) {

	mejaHandler := MejaHandler{}

	r.HandleFunc("/tables", mejaHandler.getAllMeja).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", mejaHandler.getMejaByID).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", mejaHandler.deleteMeja).Methods(http.MethodDelete)
	r.HandleFunc("/table", mejaHandler.insertMeja).Methods(http.MethodPost)

}

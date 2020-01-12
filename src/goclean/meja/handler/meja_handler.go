package handler

import (
	"encoding/json"
	"fmt"
	"goclean/meja"
	"goclean/model"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// MejaHandler represent xxxxxxx
type MejaHandler struct {
	mejaUsecase meja.MejaUsecase
}

// CreateMejaHandler represent aaaaaaa
func CreateMejaHandler(r *mux.Router, mejaUsecase meja.MejaUsecase) {

	mejaHandler := MejaHandler{mejaUsecase}

	r.HandleFunc("/tables", mejaHandler.getAllMeja).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", mejaHandler.getMejaByID).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", mejaHandler.deleteMeja).Methods(http.MethodDelete)
	r.HandleFunc("/table", mejaHandler.insertMeja).Methods(http.MethodPost)

}

func (m *MejaHandler) getAllMeja(resp http.ResponseWriter, req *http.Request) {

}

func handleSuccess(resp http.ResponseWriter, data interface{}) {
	returnData := model.ResponseWrapper{
		Success: true,
		Message: "SUCCESS",
		Data:    data,
	}

	jsonData, err := json.Marshal(returnData)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[MejaHandler.getMejaByID] Error when do json Marshalling for error handling : %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

func handleError(resp http.ResponseWriter, message string) {
	data := model.ResponseWrapper{
		Success: false,
		Message: message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[MejaHandler.getMejaByID] Error when do json Marshalling for error handling : %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

func (m *MejaHandler) getMejaByID(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)

	strID := muxVar["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		handleError(resp, "ID Harus angka")
		return
	}

	meja, err := m.mejaUsecase.GetById(id)
	if err != nil {
		handleError(resp, err.Error())
		return
	}

	handleSuccess(resp, meja)
}

func (m *MejaHandler) deleteMeja(resp http.ResponseWriter, req *http.Request) {

}

func (m *MejaHandler) insertMeja(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleError(resp, "Ooops, Something went wrong")
		fmt.Println("[MejaHandler.insertMeja] error when reading request body : " + err.Error())
		return
	}

	var meja = model.Meja{}
	err = json.Unmarshal(body, &meja)
	if err != nil {
		handleError(resp, "Ooops, Something went wrong")
		fmt.Println("[MejaHandler.insertMeja] error when unmarshall meja json : " + err.Error())
		return
	}

	if meja.ID == 0 {
		handleError(resp, "Meja Id must exist")
		return
	}

	err = m.mejaUsecase.Insert(&meja)
	if err != nil {
		handleError(resp, err.Error())
		fmt.Println("[MejaHandler.insertMeja] error when call insert service : " + err.Error())
		return
	}

	handleSuccess(resp, nil)
}

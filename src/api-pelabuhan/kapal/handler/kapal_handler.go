package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"api-pelabuhan/kapal"
	"api-pelabuhan/library"
	"api-pelabuhan/model"
)

// KapalHandler for handling request to kapal
type KapalHandler struct {
	kapalUsecase kapal.KapalUsecase
}

// CreateKapalHandler for handling request to all kapal route
func CreateKapalHandler(r *mux.Router, kapalUsecase kapal.KapalUsecase) {

	kapalHandler := KapalHandler{kapalUsecase}

	r.HandleFunc("/kapal/insert", kapalHandler.insertKapal).Methods(http.MethodPost)
	r.HandleFunc("/kapal/update", kapalHandler.updateKapal).Methods(http.MethodPut)
	r.HandleFunc("/kapal/delete/{id}", kapalHandler.deleteKapal).Methods(http.MethodDelete)
	r.HandleFunc("/kapal/view/all", kapalHandler.getAllKapal).Methods(http.MethodGet)
	r.HandleFunc("/kapal/view/id/{id}", kapalHandler.getKapalByID).Methods(http.MethodGet)

}

func (m *KapalHandler) insertKapal(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Check Your Request Body.")
		fmt.Println("[KapalHandler.insertKapal] error when reading request body: " + err.Error())
		return
	}

	var kapal = model.Kapal{}
	err = json.Unmarshal(body, &kapal)
	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Check Your Request Body.")
		fmt.Println("[KapalHandler.insertKapal] error when unmarshall kapal json : " + err.Error())
		return
	}

	if kapal.Kode == "" {
		library.HandleError(resp, "Oops, (Kode) Kapal Must Exist and Cannot be Empty.")
		return
	}

	if kapal.Muatan == 0 {
		library.HandleError(resp, "Oops, (Muatan) Kapal Must Be Exist.")
		return
	}

	if kapal.Status == "" {
		library.HandleError(resp, "Oops, (Status) Kapal Must Exist and Cannot be Empty.")
		return
	}

	if kapal.Status != "berlayar" {
		if kapal.Status != "bersandar" {
			library.HandleError(resp, "Oops, (Status) Kapal Must be (berlayar or bersandar).")
			return
		}
	}

	err = m.kapalUsecase.InsertKapal(&kapal)
	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Try Again.")
		fmt.Println("[KapalHandler.insertKapal] error when call insert service : " + err.Error())
		return
	}

	library.HandleSuccess(resp, nil)
}

func (m *KapalHandler) updateKapal(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Check Your Request Body.")
		fmt.Println("[KapalHandler.updateKapal] error when reading request body: " + err.Error())
		return
	}

	var kapal = model.Kapal{}
	err = json.Unmarshal(body, &kapal)
	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Check Your Request Body.")
		fmt.Println("[KapalHandler.updateKapal] error when unmarshall kapal json : " + err.Error())
		return
	}

	kapalID, err := m.kapalUsecase.GetKapalByID(kapal.ID)

	if kapalID == nil {
		library.HandleError(resp, "Oops, No Data Kapal to Update with ID kapal You Request.")
		return
	}

	if kapal.Kode == "" {
		library.HandleError(resp, "Oops, (Kode) Kapal Must Exist and Cannot be Empty.")
		return
	}

	if kapal.Muatan == 0 {
		library.HandleError(resp, "Oops, (Muatan) Kapal Must Be Exist.")
		return
	}

	if kapal.Status == "" {
		library.HandleError(resp, "Oops, (Status) Kapal Must Exist and Cannot be Empty.")
		return
	}

	if kapal.Status != "berlayar" {
		if kapal.Status != "bersandar" {
			library.HandleError(resp, "Oops, (Status) Kapal Must be (berlayar or bersandar).")
			return
		}
	}

	err = m.kapalUsecase.UpdateKapal(&kapal)
	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Try Again.")
		fmt.Println("[KapalHandler.updateKapal] error when call update service : " + err.Error())
		return
	}

	library.HandleSuccess(resp, nil)
}

func (m *KapalHandler) getAllKapal(resp http.ResponseWriter, req *http.Request) {
	kapal, err := m.kapalUsecase.GetAllKapal()

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong.")
		fmt.Println("[KapalHandler.getAllKapal] error when reading request body: ", err.Error())
		return
	}

	library.HandleSuccess(resp, kapal)
}

func (m *KapalHandler) getKapalByID(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strID := muxVar["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		library.HandleError(resp, "Oops, Parameter ID Must be a Number.")
		return
	}

	kapal, err := m.kapalUsecase.GetKapalByID(id)
	if err != nil {
		library.HandleError(resp, err.Error())
		return
	}

	library.HandleSuccess(resp, kapal)
}

func (m *KapalHandler) deleteKapal(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strID := muxVar["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		library.HandleError(resp, "Oops, Parameter ID Must be a Number.")
		return
	}

	kapal, err := m.kapalUsecase.GetKapalByID(id)

	if kapal == nil {
		library.HandleError(resp, "Oops, No Data Kapal to Delete with ID kapal: "+muxVar["id"]+"")
		return
	}

	err = m.kapalUsecase.DeleteKapal(id)

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Try Again.")
		return

	}

	library.HandleSuccess(resp, nil)
}

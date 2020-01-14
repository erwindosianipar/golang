package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"api-pelabuhan/dock"
	"api-pelabuhan/library"
	"api-pelabuhan/model"
)

// DockHandler for handling request to dock
type DockHandler struct {
	dockUsecase dock.DockUsecase
}

// CreateDockHandler for handling request to all dock route
func CreateDockHandler(r *mux.Router, dockUsecase dock.DockUsecase) {

	dockHandler := DockHandler{dockUsecase}

	r.HandleFunc("/dock/insert", dockHandler.insertDock).Methods(http.MethodPost)
	r.HandleFunc("/dock/update", dockHandler.updateDock).Methods(http.MethodPut)
	r.HandleFunc("/dock/delete/{id}", dockHandler.deleteDock).Methods(http.MethodDelete)
	r.HandleFunc("/dock/view/all", dockHandler.getAllDock).Methods(http.MethodGet)
	r.HandleFunc("/dock/view/id/{id}", dockHandler.getDockByID).Methods(http.MethodGet)

}

func (m *DockHandler) insertDock(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Check Your Request Body.")
		fmt.Println("[DockHandler.insertDock] error when reading request body: " + err.Error())
		return
	}

	var dock = model.Dock{}
	err = json.Unmarshal(body, &dock)
	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Check Your Request Body.")
		fmt.Println("[DockHandler.insertDock] error when unmarshall dock json : " + err.Error())
		return
	}

	if dock.Kode == "" {
		library.HandleError(resp, "Oops, (Kode) Dock Must Exist and Cannot be Empty.")
		return
	}

	err = m.dockUsecase.InsertDock(&dock)
	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Try Again.")
		fmt.Println("[DockHandler.insertDock] error when call insert service : " + err.Error())
		return
	}

	library.HandleSuccess(resp, nil)
}

func (m *DockHandler) updateDock(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Check Your Request Body.")
		fmt.Println("[DockHandler.updateDock] error when reading request body: " + err.Error())
		return
	}

	var dock = model.Dock{}
	err = json.Unmarshal(body, &dock)
	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Check Your Request Body.")
		fmt.Println("[DockHandler.updateDock] error when unmarshall dock json : " + err.Error())
		return
	}

	dockID, err := m.dockUsecase.GetDockByID(dock.ID)

	if dock.ID == 0 {
		library.HandleError(resp, "Oops, (ID) Dock Must Exist and Cannot be Empty.")
		return
	}

	if dockID == nil {
		library.HandleError(resp, "Oops, No Data dock to Update with ID dock You Request.")
		return
	}

	if dock.Kode == "" {
		library.HandleError(resp, "Oops, (Kode) Dock Must Exist and Cannot be Empty.")
		return
	}

	err = m.dockUsecase.UpdateDock(&dock)
	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Try Again.")
		fmt.Println("[DockHandler.updateDock] error when call update service : " + err.Error())
		return
	}

	library.HandleSuccess(resp, nil)
}

func (m *DockHandler) getDockByID(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strID := muxVar["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		library.HandleError(resp, "Oops, Parameter ID Must be a Number.")
		return
	}

	dock, err := m.dockUsecase.GetDockByID(id)
	if err != nil {
		library.HandleError(resp, err.Error())
		return
	}

	library.HandleSuccess(resp, dock)
}

func (m *DockHandler) deleteDock(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strID := muxVar["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		library.HandleError(resp, "Oops, Parameter ID Must be a Number.")
		return
	}

	dock, err := m.dockUsecase.GetDockByID(id)

	if dock == nil {
		library.HandleError(resp, "Oops, No Data Dock to Delete with ID DOck: "+muxVar["id"]+"")
		return
	}

	err = m.dockUsecase.DeleteDock(id)

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong. Please Try Again.")
		return

	}

	library.HandleSuccess(resp, nil)

}

func (m *DockHandler) getAllDock(resp http.ResponseWriter, req *http.Request) {
	dock, err := m.dockUsecase.GetAllDock()

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong.")
		fmt.Println("[DockHandler.getAllDock] error when reading request body: ", err.Error())
		return
	}

	library.HandleSuccess(resp, dock)
}

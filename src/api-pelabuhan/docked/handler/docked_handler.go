package handler

import (
	"api-pelabuhan/docked"
	"net/http"

	"github.com/gorilla/mux"
)

// DockedHandler for handling request to dock
type DockedHandler struct {
	dockedUsecase docked.DockedUsecase
}

// CreateDockedHandler for handling request to all dock route
func CreateDockedHandler(r *mux.Router, dockedUsecase docked.DockedUsecase) {

	dockedHandler := DockedHandler{dockedUsecase}

	r.HandleFunc("/docked/insert", dockedHandler.newDocked).Methods(http.MethodPost)
}

func (m *DockedHandler) newDocked(resp http.ResponseWriter, req *http.Request) {

}

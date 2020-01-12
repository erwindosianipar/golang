package library

import (
	"encoding/json"
	"net/http"

	"api-pelabuhan/model"
)

// HandleSuccess is for handling output json success
func HandleSuccess(resp http.ResponseWriter, data interface{}) {
	returnData := model.ResponseWrapper{
		Success: true,
		Message: "Success",
		Data:    data,
	}

	jsonData, err := json.Marshal(returnData)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong. Check Your Request Body."))
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

// HandleError is for handling output json success
func HandleError(resp http.ResponseWriter, message string) {
	data := model.ResponseWrapper{
		Success: false,
		Message: message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong. Check Your Request Body."))
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

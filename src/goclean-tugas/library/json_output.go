package library

import (
	"encoding/json"
	"fmt"
	"goclean-tugas/model"
	"net/http"
)

// JSONoutput is using to global output handler
func JSONoutput(status bool, message string, data interface{}, resp http.ResponseWriter, req *http.Request, statusCode int) {
	dataToWrap := model.ResponseWrapper{
		Success: status,
		Message: message,
		Data:    data,
	}

	dataJSON, err := json.Marshal(dataToWrap)

	if err != nil {
		fmt.Println("[library.JSONoutput] /library/json_output/ Error: when marshal dataJSON:", err)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(statusCode)
	resp.Write(dataJSON)
	return
}

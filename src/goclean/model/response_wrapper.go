package model

// ResponseWrapper represent standard response message
type ResponseWrapper struct {
	Success bool
	Message string      `json:"successMessage"`
	Data    interface{} `json:"data,omitempty"`
}

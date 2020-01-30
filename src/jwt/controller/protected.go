package controller

import (
	"fmt"
	"jwt/models"
	"jwt/utils"
	"net/http"
)

type Controller struct {
}

// ProtectedEndpoint is used to wrapping route
func (c Controller) ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
	user := models.User{
		ID:       1,
		Email:    "This is a email",
		Password: "This is a password",
	}

	utils.ResponseSuccess(w, user)

	authHeader := req.Header.Get("Authorization")
	fmt.Println(authHeader)
}

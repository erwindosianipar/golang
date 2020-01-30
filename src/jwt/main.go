package main

import (
	"database/sql"
	"fmt"
	"jwt/controller"
	"jwt/driver"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	port := "8000"
	db = driver.ConnectDB()
	defer db.Close()

	controller := controller.Controller{}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/signup", controller.Signup(db)).Methods(http.MethodPost)
	router.HandleFunc("/login", controller.Login(db)).Methods(http.MethodPost)
	router.HandleFunc("/user", controller.TokenVerifyMiddleware(controller.ProtectedEndpoint)).Methods(http.MethodGet)

	fmt.Println("Starting Web Server at Port: " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal()
	}
}

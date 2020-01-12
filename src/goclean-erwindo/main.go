package main

import (
	"fmt"
	"goclean/meja/handler"
	"goclean/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "8080"

	router := mux.NewRouter().StrictSlash(true)

	handler.CreateMejaHandler(router)

	router.Use(middleware.Logger)

	fmt.Println("starting web server")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

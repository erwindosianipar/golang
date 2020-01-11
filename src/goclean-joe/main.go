package main

import (
	"fmt"
	mejaHandler "goclean-joe/meja/handler"
	menuHandler "goclean-joe/menu/handler"
	"goclean/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "8080"

	router := mux.NewRouter().StrictSlash(true)

	mejaHandler.CreateMejaHandler(router)
	menuHandler.CreateMenuHandler(router)

	router.Use(middleware.Logger)

	fmt.Println("Starting Web Server at port : " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

}

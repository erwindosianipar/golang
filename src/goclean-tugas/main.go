package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	mejaHandler "goclean-tugas/meja/handler"
	menuHandler "goclean-tugas/menu/handler"
	"goclean-tugas/middleware"
	transaksiHandler "goclean-tugas/transaksi/handler"
)

func main() {
	port := "8080"

	router := mux.NewRouter().StrictSlash(true)

	mejaHandler.CreateMejaHandler(router)
	menuHandler.CreateMenuHandler(router)
	transaksiHandler.CreateTransaksiHandler(router)

	fmt.Printf("Web server started at http://localhost:%v\n", port)

	router.Use(middleware.Logger)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"gomux/p1/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloGetHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Method GET Hello World"))
}

func helloPostHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Method Post Hello World"))
}

func helloErwindoHandler(res http.ResponseWriter, req *http.Request) {
	pathVer := mux.Vars(req)
	res.Write([]byte("Method Post Hello " + pathVer["nama"]))
}

func main() {
	port := "8080"

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", helloGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/hello", helloPostHandler).Methods(http.MethodPost)

	routerHello := router.PathPrefix("/hello").Subrouter()
	routerHello.HandleFunc("/{nama}", helloErwindoHandler)

	router.Use(middleware.Logger) //middleware all

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

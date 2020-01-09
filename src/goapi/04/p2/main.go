package main

import (
	"fmt"
	"goapi/04/p2/middleware"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println(3)
	res.Write([]byte("Hello World"))
}

func main() {
	port := "8080"

	http.Handle("/hello", middleware.Logger(http.HandlerFunc(helloHandler)))

	fmt.Printf("Starting web server at http://localhost:%v/hello\n", port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

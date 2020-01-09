package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func logger(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		// Header-X
		res.Header().Add("header-x", "XXXXXXXXXX")

		fmt.Printf("URL: %v dipanggil pada jam %v", req.URL.Path, time.Now())
		next.ServeHTTP(res, req)
		fmt.Printf("URL: %v selesai dipanggil pada jam %v", req.URL.Path, time.Now())
	})

}

func logger2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		fmt.Println(3)
		next.ServeHTTP(resp, req)
		fmt.Println(4)
	})
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println(3)
	res.Write([]byte("Hello World"))
}

func main() {
	port := "8080"

	http.Handle("/hello", logger2(logger(http.HandlerFunc(helloHandler))))
	http.HandleFunc("/helloWorld", helloHandler)

	fmt.Printf("Starting web server at http://localhost:%v/hello\n", port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}

}

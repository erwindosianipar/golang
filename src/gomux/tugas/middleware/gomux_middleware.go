package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		// Header-X
		res.Header().Add("header-x", "XXXXXXXXXX")

		fmt.Printf("URL: %v dipanggil pada jam %v\n", req.URL.Path, time.Now())
		next.ServeHTTP(res, req)
		fmt.Printf("URL: %v selesai dipanggil pada jam %v\n", req.URL.Path, time.Now())
	})

}

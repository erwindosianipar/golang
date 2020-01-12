package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	kapalHandler "api-pelabuhan/kapal/handler"
	"api-pelabuhan/kapal/repo"
	"api-pelabuhan/kapal/usecase"
	"api-pelabuhan/middleware"
)

func main() {
	port := "8080"
	conStr := "root:erwindo123@tcp(127.0.0.1:3306)/pelabuhan"

	db, err := sql.Open("mysql", conStr)
	if err != nil {
		log.Fatal("Error When Connect to DB " + conStr + " : " + err.Error())
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	kapalRepo := repo.CreateKapalRepoMysqlImpl(db)
	kapalUsecase := usecase.CreateKapalUsecase(kapalRepo)

	kapalHandler.CreateKapalHandler(router, kapalUsecase)

	router.Use(middleware.Logger)

	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

}

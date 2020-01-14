package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	kapalHandler "api-pelabuhan/kapal/handler"
	kapalRepo "api-pelabuhan/kapal/repo"
	kapalUsecase "api-pelabuhan/kapal/usecase"

	dockHandler "api-pelabuhan/dock/handler"
	dockRepo "api-pelabuhan/dock/repo"
	dockUsecase "api-pelabuhan/dock/usecase"

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

	kapalRepo := kapalRepo.CreateKapalRepoMysqlImpl(db)
	kapalUsecase := kapalUsecase.CreateKapalUsecase(kapalRepo)
	kapalHandler.CreateKapalHandler(router, kapalUsecase)

	dockRepo := dockRepo.CreateDockRepoMysqlImpl(db)
	dockUsecase := dockUsecase.CreateDockUsecase(dockRepo)
	dockHandler.CreateDockHandler(router, dockUsecase)

	router.Use(middleware.Logger)

	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

}

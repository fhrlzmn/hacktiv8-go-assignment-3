package server

import (
	"log"
	"net/http"

	"fhrlzmn/hacktiv8-go/assignment-3/api/database"
	"fhrlzmn/hacktiv8-go/assignment-3/api/handler"
	"fhrlzmn/hacktiv8-go/assignment-3/api/repository"
	"fhrlzmn/hacktiv8-go/assignment-3/api/service"
)

func Start() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dataRepo := repository.Initialize(db)
	dataService := service.Initialize(dataRepo)
	dataHandler := handler.Initialize(dataService)

	if err := dataRepo.CreateTable(); err != nil {
		panic(err)
	}

	// prevent favicon.ico request  from chromium browser
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/", dataHandler.ServeHTTP)
	log.Print("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

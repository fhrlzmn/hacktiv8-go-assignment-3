package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"

	"fhrlzmn/hacktiv8-go/assignment-3/api/server"
	"fhrlzmn/hacktiv8-go/assignment-3/client"
)

func init() {
	// godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	go autoFetchData() // run auto fetch data in background
	server.Start()     // start local service
}

func autoFetchData() {
	// run fetch data every 15 seconds
	for {
		time.Sleep(15 * time.Second) // sleep for 15 seconds
		client.FetchData()           // fetch data from local service
	}
}

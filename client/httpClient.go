package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"fhrlzmn/hacktiv8-go/assignment-3/api/model"
)

func FetchData() {
	var data model.Response

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}

	time := data.Timestamp.Format("2006/01/02 15:04:05")

	fmt.Println("")
	fmt.Println("ID:", data.ID)
	fmt.Println("Time:", time)
	fmt.Println("{")
	fmt.Println("    Wind:", data.Data.Wind, "m/s")
	fmt.Println("    Water:", data.Data.Water, "m")
	fmt.Println("    Wind Status:", data.Data.WindStatus)
	fmt.Println("    Water Status:", data.Data.WaterStatus)
	fmt.Println("}")
}

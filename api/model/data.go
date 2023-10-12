package model

import (
	"time"
)

type Data struct {
	ID          int
	Wind        int
	Water       int
	WindStatus  string
	WaterStatus string
	Timestamp   time.Time
}

type DataResponse struct {
	Wind        int    `json:"Wind"`
	Water       int    `json:"Water"`
	WindStatus  string `json:"WindStatus"`
	WaterStatus string `json:"WaterStatus"`
}

type Response struct {
	ID        int          `json:"ID"`
	Timestamp time.Time    `json:"Timestamp"`
	Data      DataResponse `json:"Data"`
}

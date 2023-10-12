package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fhrlzmn/hacktiv8-go/assignment-3/api/model"
	"fhrlzmn/hacktiv8-go/assignment-3/api/service"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	service service.Service
}

func Initialize(s service.Service) *HandlerImpl {
	return &HandlerImpl{s}
}

func (h *HandlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
		return
	}

	h.service.Generate()
	data := h.service.Get()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dataResponse := model.Response{
		ID:        data.ID,
		Timestamp: data.Timestamp,
		Data: model.DataResponse{
			Wind:        data.Wind,
			Water:       data.Water,
			WindStatus:  data.WindStatus,
			WaterStatus: data.WaterStatus,
		},
	}

	err := json.NewEncoder(w).Encode(dataResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	return
}

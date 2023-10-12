package service

import (
	"math/rand"

	"fhrlzmn/hacktiv8-go/assignment-3/api/model"
	"fhrlzmn/hacktiv8-go/assignment-3/api/repository"
)

type Service interface {
	Generate()
	Get() model.Data
}

type ServiceImpl struct {
	repo repository.Repository
}

func Initialize(r repository.Repository) *ServiceImpl {
	return &ServiceImpl{r}
}

func (s *ServiceImpl) Get() model.Data {
	data, err := s.repo.GetData()
	if err != nil {
		panic(err)
	}

	return data
}

func (s *ServiceImpl) Generate() {
	var data model.Data

	data.Water = rand.Intn(20)
	data.Wind = rand.Intn(20)

	if data.Water <= 5 {
		data.WaterStatus = "Aman"
	} else if data.Water >= 6 && data.Water <= 8 {
		data.WaterStatus = "Siaga"
	} else {
		data.WaterStatus = "Bahaya"
	}

	if data.Wind <= 6 {
		data.WindStatus = "Aman"
	} else if data.Wind >= 7 && data.Wind <= 15 {
		data.WindStatus = "Siaga"
	} else {
		data.WindStatus = "Bahaya"
	}

	s.repo.InsertData(data)
}

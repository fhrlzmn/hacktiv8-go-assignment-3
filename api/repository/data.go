package repository

import (
	"database/sql"

	"fhrlzmn/hacktiv8-go/assignment-3/api/model"
)

type Repository interface {
	CreateTable() error
	InsertData(d model.Data) error
	GetData() (model.Data, error)
}

type RepositoryImpl struct {
	db *sql.DB
}

func Initialize(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) CreateTable() error {
	_, err := r.db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) InsertData(d model.Data) error {
	_, err := r.db.Exec(insertDataQuery, d.Wind, d.Water, d.WindStatus, d.WaterStatus)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) GetData() (model.Data, error) {
	var d model.Data

	err := r.db.QueryRow(`
		SELECT * FROM data ORDER BY id DESC LIMIT 1
	`).Scan(&d.ID, &d.Wind, &d.Water, &d.WindStatus, &d.WaterStatus, &d.Timestamp)
	if err != nil {
		return model.Data{}, err
	}

	return d, nil
}

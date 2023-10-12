package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	var (
		host   = os.Getenv("DB_HOST")
		port   = os.Getenv("DB_PORT")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASS")
		dbname = os.Getenv("DB_NAME")
	)

	db, err := sql.Open(
		"postgres",
		"postgres://"+user+":"+pass+"@"+host+":"+port+"/"+dbname+"?sslmode=disable",
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

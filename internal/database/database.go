package database

import (
	"database/sql"
	"importador_Excel/config"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	conString := config.ConnectDBUrl
	db, err := sql.Open("postgres", conString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
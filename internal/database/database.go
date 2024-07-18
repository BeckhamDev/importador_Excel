package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	conString := "user=postgres dbname=enube_teste sslmode=disable password=rootroot host=localhost port=5432"
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
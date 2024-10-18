package db

import (
	"api/src/config"
	"database/sql"
)

func DBConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBConnectionString)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
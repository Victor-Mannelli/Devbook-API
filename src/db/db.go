package db

import (
	"api/src/config"
	"api/src/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBConnectionString)
	utils.CheckError(err)

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

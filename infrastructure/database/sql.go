package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/odanaraujo/api-devbook/infrastructure/config"
)

func Connection() (*sql.DB, error) {

	db, err := sql.Open("mysql", config.StringConnectionDatabase)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

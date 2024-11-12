package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/crimsonScans.db")
	if err != nil {
		return nil, err
	}
	if err := InitDB(db); err != nil {
		return nil, err
	}
	return db, nil
}

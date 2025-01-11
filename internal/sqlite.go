package internal

import (
	"database/sql"
	"videoverse/pkg/config"
)

func MakeSQLiteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.DATABASE_PATH)
	if err != nil {
		return nil, err
	}
	return db, nil
}

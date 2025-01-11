package internal

import (
	"database/sql"
	"videoverse/pkg/config"
	"videoverse/pkg/logbox"
)

func MakeSQLiteConnection() *sql.DB {
	logbox.NewLogBox().Info().Msg("setting up sqlite connection")
	db, err := sql.Open("sqlite3", config.DATABASE_PATH)
	if err != nil {
		logbox.NewLogBox().Fatal().Err(err).Msg("failed to open sqlite database connection")
	}
	logbox.NewLogBox().Info().Msg("successfully opened sqlite database connection")
	return db
}

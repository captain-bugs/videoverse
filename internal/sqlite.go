package internal

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/mattn/go-sqlite3"
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

func RunMigrations() {
	db := MakeSQLiteConnection()
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		logbox.NewLogBox().Error().Err(err).Msg("failed to create migration driver")
	}

	m, err := migrate.NewWithDatabaseInstance(config.MIGRATIONS_PATH, "videoverse", driver)
	if err != nil {
		logbox.NewLogBox().Error().Err(err).Msg("failed to create migration instance")
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logbox.NewLogBox().Fatal().Err(err).Msg("failed to run migrations")
	}
}

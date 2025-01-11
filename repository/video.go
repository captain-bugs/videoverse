package repository

import (
	"database/sql"
	"videoverse/pkg/models"
)

type VideoRepository struct {
	connection *sql.DB
}

func NewVideoRepository(connection *sql.DB) models.IVideoRepo {
	return &VideoRepository{connection: connection}
}

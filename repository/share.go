package repository

import (
	"database/sql"
	"videoverse/pkg/models"
)

type ShareRepository struct{ connection *sql.DB }

func NewShareRepository(connection *sql.DB) models.IShareRepo {
	return &ShareRepository{connection: connection}
}

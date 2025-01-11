package repository

import (
	"database/sql"
	"videoverse/pkg/models"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) models.IUserRepo {
	return &UserRepository{connection: connection}
}

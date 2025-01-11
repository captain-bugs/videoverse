package repository

import (
	"context"
	"database/sql"
	videoversedb "videoverse/db/videoverse"
	"videoverse/pkg/models"
)

type ShareRepository struct{ db *videoversedb.Queries }

func NewShareRepository(connection *sql.DB) models.IShareRepo {
	return &ShareRepository{db: videoversedb.New(connection)}
}

func (s ShareRepository) GetByID(ctx context.Context, ID string) (*models.Share, error) {
	//TODO implement me
	panic("implement me")
}

func (s ShareRepository) Create(ctx context.Context, share *models.Share) (*models.Share, error) {
	//TODO implement me
	panic("implement me")
}

func (s ShareRepository) Update(ctx context.Context, share *models.Share) (*models.Share, error) {
	//TODO implement me
	panic("implement me")
}

func (s ShareRepository) Delete(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}

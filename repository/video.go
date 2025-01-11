package repository

import (
	"context"
	"database/sql"
	videoversedb "videoverse/db/videoverse"
	"videoverse/pkg/models"
)

type VideoRepository struct {
	db *videoversedb.Queries
}

func NewVideoRepository(connection *sql.DB) models.IVideoRepo {
	return &VideoRepository{db: videoversedb.New(connection)}
}

func (v VideoRepository) GetByID(ctx context.Context, ID string) (*models.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRepository) Create(ctx context.Context, video *models.Video) (*models.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRepository) Update(ctx context.Context, video *models.Video) (*models.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRepository) Delete(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}

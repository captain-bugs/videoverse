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

	var arg = videoversedb.SaveVideoParams{
		Title:       video.Title,
		Description: video.Description,
		UserID:      video.UserID,
		Type:        string(video.Type),
		FilePath:    video.FilePath,
		FileName: sql.NullString{
			String: video.FileName,
			Valid:  true,
		},
		Metadata: sql.NullString{
			String: video.MetadataString(),
			Valid:  len(video.MetadataString()) > 0,
		},
		SizeInBytes: video.SizeInBytes,
		Duration:    int64(video.Metadata.Duration),
		CreatedAt: sql.NullTime{
			Time:  video.CreatedAt,
			Valid: true,
		},
	}

	outcome, err := v.db.SaveVideo(ctx, arg)
	if err != nil {
		return nil, err
	}
	video.ID = outcome.ID
	return video, nil
}

func (v VideoRepository) Update(ctx context.Context, video *models.Video) (*models.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoRepository) Delete(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}

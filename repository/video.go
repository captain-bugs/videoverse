package repository

import (
	"context"
	"database/sql"
	"time"
	"videoverse/av"
	videoversedb "videoverse/db/videoverse"
	"videoverse/pkg/models"
	"videoverse/pkg/utils"
)

type VideoRepository struct {
	db *videoversedb.Queries
}

func NewVideoRepository(connection *sql.DB) models.IVideoRepo {
	return &VideoRepository{db: videoversedb.New(connection)}
}

func (v VideoRepository) GetByID(ctx context.Context, ID int64) (*models.Video, error) {
	outcome, err := v.db.GetVideoByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	var video = models.Video{
		ID:          outcome.ID,
		Title:       outcome.Title,
		Description: outcome.Description,
		UserID:      outcome.UserID,
		Type:        models.VIDEO_TYPE(outcome.Type),
		FilePath:    outcome.FilePath,
		FileName:    outcome.FileName,
		SizeInBytes: outcome.SizeInBytes,
		Duration:    outcome.Duration,

		StartTime: 0,
		EndTime:   0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	if outcome.SourceVideoID.Valid {
		video.SourceVideoID = &outcome.SourceVideoID.Int64
	}
	if outcome.Metadata.Valid {
		video.Metadata = utils.MapToStruct[any, *av.AVFile](utils.StringToMap(outcome.Metadata.String))
	}

	return &video, nil
}

func (v VideoRepository) Create(ctx context.Context, video *models.Video) (*models.Video, error) {

	var arg = videoversedb.SaveVideoParams{
		UserID:      video.UserID,
		Title:       video.Title,
		Description: video.Description,
		Type:        string(video.Type),
		FilePath:    video.FilePath,
		FileName:    video.FileName,
		SizeInBytes: video.SizeInBytes,
		Duration:    video.Metadata.Duration,
		Metadata: sql.NullString{
			String: video.MetadataString(),
			Valid:  len(video.MetadataString()) > 0,
		},
		StartTime: sql.NullFloat64{
			Float64: video.StartTime,
			Valid:   video.StartTime > 0,
		},
		EndTime: sql.NullFloat64{
			Float64: video.EndTime,
			Valid:   video.EndTime > 0,
		},
		CreatedAt: sql.NullTime{
			Time:  video.CreatedAt,
			Valid: true,
		},
	}
	if video.SourceVideoID != nil {
		arg.SourceVideoID = sql.NullInt64{
			Int64: *video.SourceVideoID,
			Valid: true,
		}
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

func (v VideoRepository) Delete(ctx context.Context, ID int64) error {
	//TODO implement me
	panic("implement me")
}

func (v VideoRepository) ListByUserID(ctx context.Context, userID int64) ([]*models.Video, error) {
	outcome, err := v.db.GetVideosByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var videos = make([]*models.Video, 0)
	for _, v := range outcome {
		var video = models.Video{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			UserID:      v.UserID,
			Type:        models.VIDEO_TYPE(v.Type),
			FilePath:    v.FilePath,
			FileName:    v.FileName,
			SizeInBytes: v.SizeInBytes,
			Duration:    v.Duration,

			StartTime: 0,
			EndTime:   0,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		if v.SourceVideoID.Valid {
			video.SourceVideoID = &v.SourceVideoID.Int64
		}
		if v.Metadata.Valid {
			video.Metadata = utils.MapToStruct[any, *av.AVFile](utils.StringToMap(v.Metadata.String))
		}
		videos = append(videos, &video)
	}
	return videos, nil
}

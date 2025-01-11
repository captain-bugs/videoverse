package v1_handler

import (
	"bytes"
	"context"
	"io"
	"time"
	"videoverse/pkg/config"
	"videoverse/pkg/models"
)

type IVideoHandlerV1 interface {
	GetVideo(ctx context.Context, videoID int64) (any, error)
	PostVideo(ctx context.Context, payload *models.ReqSaveVideo) (any, error)
}

func (h *HandlerV1) GetVideo(ctx context.Context, videoID int64) (any, error) {
	return nil, nil
}

func (h *HandlerV1) PostVideo(ctx context.Context, payload *models.ReqSaveVideo) (any, error) {

	if _, err := h.repo.Storage().Upload(io.NopCloser(bytes.NewReader(payload.AVFile.InBytes)), payload.File.Filename, config.FILE_UPLOAD_PATH); err != nil {
		return nil, err
	}
	var video = models.Video{
		Title:       payload.Title,
		Description: payload.Description,
		UserID:      payload.UserID,
		Type:        models.ORIGINAL,
		FilePath:    payload.AVFile.Path,
		FileName:    payload.File.Filename,
		SizeInBytes: int64(len(payload.AVFile.InBytes)),
		Duration:    payload.AVFile.Duration,
		Metadata:    *payload.AVFile,
		CreatedAt:   time.Now().UTC(),
	}
	outcome, err := h.repo.Video().Create(ctx, &video)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

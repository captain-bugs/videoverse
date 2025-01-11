package v1_handler

import (
	"context"
	"videoverse/pkg/models"
)

type IVideoHandlerV1 interface {
	GetVideo(ctx context.Context, videoID int64) (any, error)
	PostVideo(ctx context.Context, payload models.ReqSaveVideo) (any, error)
}

func (h *HandlerV1) GetVideo(ctx context.Context, videoID int64) (any, error) {
	return nil, nil
}

func (h *HandlerV1) PostVideo(ctx context.Context, payload models.ReqSaveVideo) (any, error) {
	return nil, nil
}

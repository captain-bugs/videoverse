package v1_handler

import "videoverse/pkg/models"

type IVideoHandlerV1 interface {
	GetVideo(videoID string) (any, error)
	PostVideo(payload models.ReqSaveVideo) (any, error)
}

func (h *HandlerV1) GetVideo(videoID string) (any, error) {
	return nil, nil
}

func (h *HandlerV1) PostVideo(payload models.ReqSaveVideo) (any, error) {
	return nil, nil
}

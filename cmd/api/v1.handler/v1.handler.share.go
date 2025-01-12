package v1_handler

import (
	"context"
	"errors"
	"fmt"
	"time"
	"videoverse/pkg/auth"
	"videoverse/pkg/config"
	"videoverse/pkg/models"
	"videoverse/response"
)

type IShareHandlerV1 interface {
	GetGenerateShareLink(ctx context.Context, payload *models.ReqShare) (any, error)
	GetViewFile(ctx context.Context, signature string) (any, error)
}

func (h *HandlerV1) GetGenerateShareLink(ctx context.Context, payload *models.ReqShare) (any, error) {

	video, err := h.repo.Video().GetByID(ctx, payload.VideoID)
	if err != nil {
		return nil, response.BadRequest(errors.New("video not found"))
	}
	expiry := time.Now().Add(time.Hour * 24)
	token, err := auth.GenerateSignedToken(payload.UserID, video.ID, video.FilePath, expiry)
	if err != nil {
		return nil, response.BadRequest(errors.New("failed to generate token"))
	}

	var result = make(map[string]any)
	result["link"] = fmt.Sprintf("%s/api/v1/share/view/?signature=%s", config.CDN_ENDPOINT, *token)
	return result, nil

}

func (h *HandlerV1) GetViewFile(ctx context.Context, signature string) (any, error) {
	data, err := auth.VerifySignedToken(signature)
	if err != nil {
		return nil, response.BadRequest(errors.New("invalid signature"))
	}
	return data["file_path"], nil
}

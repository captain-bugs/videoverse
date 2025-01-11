package v1_handler

import (
	"context"
	"videoverse/pkg/models"
)

type IShareHandlerV1 interface {
	PostShare(ctx context.Context, payload models.ReqShare) (any, error)
}

func (h *HandlerV1) PostShare(ctx context.Context, payload models.ReqShare) (any, error) {
	return nil, nil
}

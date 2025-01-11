package v1_controller

import (
	"context"
	"videoverse/pkg/models"
)

type GetUser func(ctx context.Context, userID string) (any, error)
type GetVideo func(ctx context.Context, videoID string) (any, error)
type PostVideo func(ctx context.Context, payload models.ReqSaveVideo) (any, error)
type PostShare func(ctx context.Context, payload models.ReqShare) (any, error)

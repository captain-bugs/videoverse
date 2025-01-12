package v1_controller

import (
	"context"
	"videoverse/pkg/models"
)

type GetUser func(ctx context.Context, userID int64) (any, error)
type PostUser func(ctx context.Context, payload *models.ReqSaveUser) (any, error)
type GetUserVideos func(ctx context.Context, userID int64) (any, error)
type GetVideo func(ctx context.Context, videoID int64) (any, error)
type PostVideo func(ctx context.Context, payload *models.ReqSaveVideo) (any, error)
type PostTrimVideo func(ctx context.Context, payload *models.ReqTrimVideo) (any, error)
type PostMergeVideo func(ctx context.Context, payload *models.ReqMergeVideo) (any, error)
type GetGenerateShareLink func(ctx context.Context, payload *models.ReqShare) (any, error)
type GetViewFile func(ctx context.Context, signature string) (any, error)

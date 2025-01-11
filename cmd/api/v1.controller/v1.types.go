package v1_controller

import "videoverse/pkg/models"

type GetUser func(userID string) (any, error)
type GetVideo func(videoID string) (any, error)
type PostVideo func(payload models.ReqSaveVideo) (any, error)
type PostShare func(payload models.ReqShare) (any, error)

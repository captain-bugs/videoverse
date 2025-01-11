package models

import "context"

type Video struct {
}

type IVideoRepo interface {
	GetByID(ctx context.Context, ID string) (*Video, error)
	Create(ctx context.Context, video *Video) (*Video, error)
	Update(ctx context.Context, video *Video) (*Video, error)
	Delete(ctx context.Context, ID string) error
}

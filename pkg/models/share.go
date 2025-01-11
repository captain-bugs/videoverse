package models

import "context"

type Share struct {
	ID string `json:"id"`
}

type IShareRepo interface {
	GetByID(ctx context.Context, ID string) (*Share, error)
	Create(ctx context.Context, share *Share) (*Share, error)
	Update(ctx context.Context, share *Share) (*Share, error)
	Delete(ctx context.Context, ID string) error
}

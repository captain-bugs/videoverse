package v1_handler

import (
	"context"
	"time"
	"videoverse/pkg/models"
)

type IUserHandlerV1 interface {
	GetUser(ctx context.Context, userID int64) (any, error)
	PostUser(ctx context.Context, payload models.ReqSaveUser) (any, error)
}

func (h *HandlerV1) GetUser(ctx context.Context, userID int64) (any, error) {
	user, err := h.repo.User().GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *HandlerV1) PostUser(ctx context.Context, payload models.ReqSaveUser) (any, error) {
	user, err := h.repo.User().Create(ctx, &models.User{
		Username:  payload.Username,
		Email:     payload.Email,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

package v1_handler

import "context"

type IUserHandlerV1 interface {
	GetUser(ctx context.Context, userID string) (any, error)
}

func (h *HandlerV1) GetUser(ctx context.Context, userID string) (any, error) {
	
	return nil, nil
}

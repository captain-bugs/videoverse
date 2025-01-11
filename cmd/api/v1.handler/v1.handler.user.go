package v1_handler

type IUserHandlerV1 interface {
	GetUser(userID string) (any, error)
}

func (h *HandlerV1) GetUser(userID string) (any, error) {
	return nil, nil
}

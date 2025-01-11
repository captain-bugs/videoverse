package v1_handler

import "videoverse/pkg/models"

type IShareHandlerV1 interface {
	PostShare(payload models.ReqShare) (any, error)
}

func (h *HandlerV1) PostShare(payload models.ReqShare) (any, error) {
	return nil, nil
}

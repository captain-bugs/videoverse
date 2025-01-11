package v1_handler

import "videoverse/repository"

type IHandlerV1 interface {
	IUserHandlerV1
	IVideoHandlerV1
	IShareHandlerV1
}

type HandlerV1 struct {
	repo repository.IRepo
}

func NewHandlerV1(repo repository.IRepo) IHandlerV1 {
	return &HandlerV1{repo: repo}
}

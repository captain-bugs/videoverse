package v1_controller

import "github.com/gin-gonic/gin"

type IVideoControllerV1 interface {
	GetVideo(ctx *gin.Context, fn GetVideo) error
	PostVideo(ctx *gin.Context, fn PostVideo) error
}

func (c *ControllerV1) GetVideo(ctx *gin.Context, fn GetVideo) error {
	return nil
}

func (c *ControllerV1) PostVideo(ctx *gin.Context, fn PostVideo) error {
	return nil
}

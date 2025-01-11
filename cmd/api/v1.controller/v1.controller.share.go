package v1_controller

import "github.com/gin-gonic/gin"

type IShareControllerV1 interface {
	PostShare(ctx *gin.Context, fn PostShare) error
}

func (c *ControllerV1) PostShare(ctx *gin.Context, fn PostShare) error {
	return nil
}

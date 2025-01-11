package v1_controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"videoverse/response"
)

type IUserControllerV1 interface {
	GetUser(c *gin.Context, fn GetUser) error
}

func (c *ControllerV1) GetUser(ctx *gin.Context, fn GetUser) error {
	id, exist := ctx.Get("user_id")
	if !exist {
		return ctx.AbortWithError(400, response.NewAPIError(400, errors.New("user_id not found in context")))
	}

	rc := ctx.Request.Context()
	_, err := fn(rc, id.(string))
	if err != nil {
		return ctx.AbortWithError(400, response.NewAPIError(500, err))
	}
	return nil
}

package v1_controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"videoverse/pkg/models"
	"videoverse/response"
)

type IShareControllerV1 interface {
	GetGenerateShareLink(ctx *gin.Context, fn GetGenerateShareLink) error
	GetViewFile(ctx *gin.Context, fn GetViewFile) error
}

func (c *ControllerV1) GetGenerateShareLink(ctx *gin.Context, fn GetGenerateShareLink) error {
	userID, exist := ctx.Get("user_id")
	if !exist {
		return ctx.AbortWithError(400, response.NewAPIError(400, errors.New("user_id not found in context")))
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return response.BadRequest(err)
	}
	payload := &models.ReqShare{
		VideoID: int64(id),
		UserID:  userID.(int64),
	}
	reply, err := fn(ctx, payload)
	if err != nil {
		return err
	}
	ctx.JSON(200, reply)
	return nil
}

func (c *ControllerV1) GetViewFile(ctx *gin.Context, fn GetViewFile) error {
	signature := ctx.Query("signature")
	if signature == "" {
		return response.BadRequest(errors.New("signature is required"))
	}
	reply, err := fn(ctx, signature)
	if err != nil {
		return err
	}
	ctx.Header("Cache-Control", "max-age=31536000")
	ctx.File(reply.(string))
	return nil
}

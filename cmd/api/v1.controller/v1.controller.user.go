package v1_controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"videoverse/pkg/models"
	"videoverse/response"
)

type IUserControllerV1 interface {
	GetUser(c *gin.Context, fn GetUser) error
	PostUser(c *gin.Context, fn PostUser) error
}

func (c *ControllerV1) GetUser(ctx *gin.Context, fn GetUser) error {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		return response.BadRequest(errors.New("user_id not found in context"))
	}

	rc := ctx.Request.Context()
	_, err := fn(rc, userID)
	if err != nil {
		return response.NewAPIError(500, err)
	}
	return nil
}

func (c *ControllerV1) PostUser(ctx *gin.Context, fn PostUser) error {
	var request *models.ReqSaveUser
	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		return response.BadRequest(err)
	}
	if problems := models.Validate(request, make(map[string]any)); len(problems) > 0 {
		return response.ErrorsInRequestBody(problems)
	}
	rc := ctx.Request.Context()
	reply, err := fn(rc, request)
	if err != nil {
		return err
	}
	ctx.JSON(201, reply)
	return nil
}

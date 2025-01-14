package v1_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"videoverse/av"
	"videoverse/pkg/models"
	"videoverse/response"
)

type IVideoControllerV1 interface {
	GetUserVideos(ctx *gin.Context, fn GetUserVideos) error
	GetVideo(ctx *gin.Context, fn GetVideo) error
	PostVideo(ctx *gin.Context, fn PostVideo) error
	PostTrimVideo(ctx *gin.Context, fn PostTrimVideo) error
	PostMergeVideo(ctx *gin.Context, fn PostMergeVideo) error
}

func (c *ControllerV1) GetUserVideos(ctx *gin.Context, fn GetUserVideos) error {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		return response.BadRequest(errors.New("user_id not found in context"))
	}
	reply, err := fn(ctx, userID)
	if err != nil {
		return err
	}
	ctx.JSON(200, reply)
	return nil
}

func (c *ControllerV1) GetVideo(ctx *gin.Context, fn GetVideo) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return response.BadRequest(err)
	}
	reply, err := fn(ctx, int64(id))
	if err != nil {
		return err
	}
	ctx.JSON(200, reply)
	return nil
}

func (c *ControllerV1) PostVideo(ctx *gin.Context, fn PostVideo) error {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		return response.BadRequest(errors.New("user_id not found in context"))
	}
	var req models.ReqSaveVideo
	if err := ctx.ShouldBind(&req); err != nil {
		return response.BadRequest(err)
	}
	if !req.IsFileSizeValid() {
		return response.BadRequest(errors.New("file size is too large"))
	}
	if !req.IsFileTypeValid() {
		return response.BadRequest(errors.New(fmt.Sprintf("file type is not supported: %s", req.File.Header.Get("Content-Type"))))
	}
	byts, err := req.GetFile()
	if err != nil {
		return response.BadRequest(err)
	}
	avfile := av.NewAVFile(av.WithBytes(byts))
	avfile.FetchMetaInfo()
	if errs := avfile.Validate(); len(errs) > 0 {
		return response.ErrorsInRequestBody(errs)
	}
	req.UserID = userID
	req.AVFile = avfile
	reply, err := fn(ctx, &req)
	if err != nil {
		return err
	}
	ctx.JSON(200, reply)
	return nil
}

func (c *ControllerV1) PostTrimVideo(ctx *gin.Context, fn PostTrimVideo) error {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		return response.BadRequest(errors.New("user_id not found in context"))
	}
	var request models.ReqTrimVideo
	request.UserID = userID
	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		return response.BadRequest(err)
	}
	if problems := models.Validate(request, make(map[string]any)); len(problems) > 0 {
		return response.ErrorsInRequestBody(problems)
	}
	reply, err := fn(ctx, &request)
	if err != nil {
		return err
	}
	ctx.JSON(200, reply)
	return nil
}

func (c *ControllerV1) PostMergeVideo(ctx *gin.Context, fn PostMergeVideo) error {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		return response.BadRequest(errors.New("user_id not found in context"))
	}
	var request models.ReqMergeVideo
	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		return response.BadRequest(err)
	}
	if problems := models.Validate(request, make(map[string]any)); len(problems) > 0 {
		return response.ErrorsInRequestBody(problems)
	}
	request.UserID = userID
	reply, err := fn(ctx, &request)
	if err != nil {
		return err
	}
	ctx.JSON(200, reply)
	return nil
}

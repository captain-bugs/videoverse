package v1_controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"videoverse/av"
	"videoverse/pkg/models"
	"videoverse/response"
)

type IVideoControllerV1 interface {
	GetVideo(ctx *gin.Context, fn GetVideo) error
	PostVideo(ctx *gin.Context, fn PostVideo) error
}

func (c *ControllerV1) GetVideo(ctx *gin.Context, fn GetVideo) error {
	return nil
}

func (c *ControllerV1) PostVideo(ctx *gin.Context, fn PostVideo) error {
	var req *models.ReqSaveVideo
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
	if errs := avfile.Validate(); errs != nil {
		return response.ErrorsInRequestBody(errs)
	}
	req.AVFile = avfile
	reply, err := fn(ctx, req)
	if err != nil {
		return err
	}
	ctx.JSON(200, reply)
	return nil
}

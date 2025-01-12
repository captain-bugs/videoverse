package models

import (
	"mime/multipart"
	"videoverse/av"
	"videoverse/pkg/utils"
)

type ReqSaveVideo struct {
	Title       string                `form:"title"        binding:"required"`
	Description string                `form:"description"  binding:"required"`
	File        *multipart.FileHeader `form:"file"         binding:"required"`
	AVFile      *av.AVFile
	UserID      int64
}

func (r *ReqSaveVideo) IsFileSizeValid() bool {
	// file size should be either greater than 5MB or less than 25MB
	return r.File.Size > 5*1024*1024 && r.File.Size < 25*1024*1024

}

func (r *ReqSaveVideo) IsFileTypeValid() bool {
	return utils.SupportedFileTypes(r.File.Header.Get("Content-Type"))
}

func (r *ReqSaveVideo) IsFileDurationValid() bool {
	return false
}

func (r *ReqSaveVideo) GetFile() ([]byte, error) {
	return utils.ReadMultipartFileHeader(r.File)
}

type ReqTrimVideo struct {
	VideoID     int64   `json:"video_id"     binding:"required"`
	StartTime   float64 `json:"start_time"   binding:"required,gt=0"`
	EndTime     float64 `json:"end_time"     binding:"required,gtfield=StartTime"`
	Title       string  `json:"title"        binding:"omitempty"`
	Description string  `json:"description"  binding:"omitempty"`
	UserID      int64
}

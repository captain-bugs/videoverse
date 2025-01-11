package models

import (
	"context"
	"encoding/json"
	"time"
	"videoverse/av"
	"videoverse/pkg/utils"
)

type VIDEO_TYPE string

const (
	ORIGINAL VIDEO_TYPE = "ORIGINAL"
	TRIMMED  VIDEO_TYPE = "TRIMMED"
	MERGED   VIDEO_TYPE = "MERGED"
)

type Video struct {
	ID            int64      `json:"id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	UserID        int64      `json:"user_id"`
	SourceVideoID int64      `json:"source_video_id"`
	Type          VIDEO_TYPE `json:"type"`
	FilePath      string     `json:"file_path"`
	FileName      string     `json:"file_name"`
	SizeInBytes   int64      `json:"size_in_bytes"`
	Duration      float64    `json:"duration"`
	Metadata      av.AVFile  `json:"metadata"`
	StartTime     int64      `json:"start_time"`
	EndTime       int64      `json:"end_time"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (v *Video) MetadataString() string {
	data := utils.ToMap(v.Metadata)
	delete(data, "in_bytes")
	delete(data, "out_bytes")
	byts, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(byts)

}

type IVideoRepo interface {
	GetByID(ctx context.Context, ID string) (*Video, error)
	Create(ctx context.Context, video *Video) (*Video, error)
	Update(ctx context.Context, video *Video) (*Video, error)
	Delete(ctx context.Context, ID string) error
}

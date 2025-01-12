package v1_handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"time"
	"videoverse/av"
	"videoverse/pkg/config"
	"videoverse/pkg/models"
	"videoverse/pkg/utils"

	"videoverse/response"
)

type IVideoHandlerV1 interface {
	GetVideo(ctx context.Context, videoID int64) (any, error)
	PostVideo(ctx context.Context, payload *models.ReqSaveVideo) (any, error)
	PostTrimVideo(ctx context.Context, payload *models.ReqTrimVideo) (any, error)
	PostMergeVideo(ctx context.Context, payload *models.ReqMergeVideo) (any, error)
}

func (h *HandlerV1) GetVideo(ctx context.Context, videoID int64) (any, error) {
	video, err := h.repo.Video().GetByID(ctx, videoID)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (h *HandlerV1) PostVideo(ctx context.Context, payload *models.ReqSaveVideo) (any, error) {

	filename := fmt.Sprintf("O_%s_%s", utils.GenerateUUID(), payload.File.Filename)
	filepath := fmt.Sprintf("%s/%s", config.FILE_UPLOAD_PATH, filename)
	if _, err := h.repo.Storage().Upload(io.NopCloser(bytes.NewReader(payload.AVFile.InBytes)), filename, config.FILE_UPLOAD_PATH); err != nil {
		return nil, err
	}
	var video = models.Video{
		Title:       payload.Title,
		Description: payload.Description,
		UserID:      payload.UserID,
		Type:        models.ORIGINAL,
		FilePath:    filepath,
		FileName:    payload.File.Filename,
		SizeInBytes: int64(len(payload.AVFile.InBytes)),
		Duration:    payload.AVFile.Duration,
		Metadata:    payload.AVFile,
		CreatedAt:   time.Now().UTC(),
	}
	outcome, err := h.repo.Video().Create(ctx, &video)
	if err != nil {
		return nil, err
	}
	outcome.Metadata.InBytes = nil
	outcome.Metadata.OutBytes = nil
	data := utils.ToMap(outcome)
	return data, nil
}

func (h *HandlerV1) PostTrimVideo(ctx context.Context, payload *models.ReqTrimVideo) (any, error) {

	// get the video
	video, err := h.repo.Video().GetByID(ctx, payload.VideoID)
	if err != nil {
		return nil, err
	}
	if !video.IsFileAvailable() {
		return nil, response.InternalServerError(errors.New("file not found"))
	}
	avfile := av.NewAVFile(av.WithFile(video.FilePath))
	avfile.FetchMetaInfo()
	avfile.Name = video.FileName
	trimmed, err := avfile.Trim(payload.StartTime, payload.EndTime)
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	var trimmedVideo = models.Video{
		Title:         payload.Title,
		Description:   payload.Description,
		UserID:        payload.UserID,
		Type:          models.TRIMMED,
		FilePath:      trimmed.Path,
		FileName:      trimmed.Name,
		SizeInBytes:   int64(len(trimmed.InBytes)),
		Duration:      trimmed.Duration,
		Metadata:      trimmed,
		SourceVideoID: &video.ID,
		StartTime:     payload.StartTime,
		EndTime:       payload.EndTime,
		CreatedAt:     time.Now().UTC(),
	}
	outcome, err := h.repo.Video().Create(ctx, &trimmedVideo)
	if err != nil {
		return nil, err
	}
	outcome.Metadata.InBytes = nil
	outcome.Metadata.OutBytes = nil
	data := utils.ToMap(outcome)
	return data, nil
}

func (h *HandlerV1) PostMergeVideo(ctx context.Context, payload *models.ReqMergeVideo) (any, error) {

	// get the videos
	var videos []*av.AVFile
	for _, vid := range payload.VideoIDs {
		video, err := h.repo.Video().GetByID(ctx, vid)
		if err != nil {
			return nil, err
		}
		if !video.IsFileAvailable() {
			return nil, response.InternalServerError(errors.New("file not found"))
		}
		avfile := av.NewAVFile(av.WithFile(video.FilePath))
		videos = append(videos, avfile)
	}

	// merge the videos
	filename := fmt.Sprintf("M_%s_%s.mp4", utils.GenerateUUID(), utils.ToFlatCase(payload.Title))
	filepath := fmt.Sprintf("%s/%s", config.FILE_UPLOAD_PATH, filename)
	merged := av.Merge(filename, filepath, videos)
	if merged == nil {
		return nil, response.InternalServerError(errors.New("failed to merge videos"))
	}

	// save the merged video
	var mergedVideo = models.Video{
		Title:       payload.Title,
		Description: payload.Description,
		UserID:      payload.UserID,
		Type:        models.MERGED,
		FilePath:    merged.Path,
		FileName:    merged.Name,
		SizeInBytes: int64(len(merged.InBytes)),
		Duration:    merged.Duration,
		Metadata:    merged,
		CreatedAt:   time.Now().UTC(),
	}
	outcome, err := h.repo.Video().Create(ctx, &mergedVideo)
	if err != nil {
		return nil, err
	}
	outcome.Metadata.InBytes = nil
	outcome.Metadata.OutBytes = nil
	data := utils.ToMap(outcome)
	return data, nil
}

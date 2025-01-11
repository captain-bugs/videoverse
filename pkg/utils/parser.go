package utils

import (
	"io"
	"mime/multipart"
	"videoverse/response"
)

func ReadMultipartFileHeader(fh *multipart.FileHeader) ([]byte, error) {
	f, err := fh.Open()
	if err != nil {
		return nil, response.BadRequest(err)
	}
	defer f.Close()
	fb, err := io.ReadAll(f)
	if err != nil {
		return nil, response.BadRequest(err)
	}
	return fb, nil
}

func SupportedFileTypes(fileType string) bool {
	// Check if the file type is supported (images,csv and ttf)
	switch fileType {
	case "video/mp4",
		"video/ogg",
		"application/octet-stream",
		"video/webm":
		return true
	}
	return false
}

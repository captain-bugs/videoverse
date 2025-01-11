package storage

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	"videoverse/pkg/logbox"
)

type IFileStore interface {
	Upload(file io.Reader, key string, bucket string) (any, error)
	Download(key string, bucket string) ([]byte, error)
	Exist(key, bucket string) (any, bool)
}

type Disk struct{}

func NewDisk() IFileStore {
	return &Disk{}
}

func (storage Disk) Upload(file io.Reader, key string, path string) (any, error) {
	t1 := time.Now()
	directory := filepath.Join(path)
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		logbox.NewLogBox().Error().
			Err(err).
			Str("event", "UPLOAD_TO_DISK").
			Str("key", key).
			Str("bucket", path).
			Msg("failed to create directory")
		return nil, err
	}

	filePath := filepath.Join(directory, key)
	outFile, err := os.Create(filePath)
	if err != nil {
		logbox.NewLogBox().Error().
			Err(err).
			Str("event", "UPLOAD_TO_DISK").
			Str("key", key).
			Str("bucket", path).
			Msg("failed to create file")
		return nil, err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, file); err != nil {
		logbox.NewLogBox().Error().
			Err(err).
			Str("event", "UPLOAD_TO_DISK").
			Str("key", key).
			Str("bucket", path).
			Msg("failed to write to file")
		return nil, err
	}

	logbox.NewLogBox().Debug().
		Str("event", "UPLOAD_TO_DISK").
		Str("key", key).
		Str("bucket", path).
		Str("ms", time.Since(t1).String()).
		Msg("file uploaded successfully")

	return filePath, nil
}

func (storage Disk) Download(key string, path string) ([]byte, error) {
	filePath := filepath.Join(path, key)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logbox.NewLogBox().Error().
			Err(err).
			Str("event", "DOWNLOAD_FROM_DISK").
			Str("key", key).
			Str("path", path).
			Msg("failed to read file")
		return nil, err
	}

	if len(data) == 0 {
		logbox.NewLogBox().Error().
			Str("event", "ZERO_BYTES_DATA").
			Str("key", key).
			Str("path", path).
			Msg("empty file")
		return nil, errors.New("zero bytes written to memory")
	}

	return data, nil
}

func (storage Disk) Exist(key, path string) (any, bool) {
	filePath := filepath.Join(path, key)
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false
		}
		return nil, false
	}
	return info, true
}

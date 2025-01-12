package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	v1_controller "videoverse/cmd/api/v1.controller"
	v1_handler "videoverse/cmd/api/v1.handler"
	"videoverse/repository"
)

func NewTestRequest(t *testing.T, method, path string, payload any, isMultipart bool) *http.Request {
	t.Helper()

	var req *http.Request
	var err error

	u, err := url.Parse(path)
	if err != nil {
		t.Fatalf("Failed to parse URL: %v", err)
	}

	if method == http.MethodGet || payload == nil {
		req, err = http.NewRequest(method, u.String(), nil)
	} else {
		var body bytes.Buffer
		if isMultipart {
			writer := multipart.NewWriter(&body)
			for key, val := range payload.(map[string]string) {
				if key == "file" {
					currentDir, err := os.Getwd()
					if err != nil {
						t.Fatalf("Failed to get current directory: %v", err)
					}
					filePath := currentDir + "/" + val
					file, err := os.Open(filePath)
					if err != nil {
						t.Fatalf("Failed to open file: %v", err)
					}
					defer file.Close()

					part, err := writer.CreateFormFile(key, val)
					if err != nil {
						t.Fatalf("Failed to create form file: %v", err)
					}
					_, err = io.Copy(part, file)
					if err != nil {
						t.Fatalf("Failed to copy form file: %v", err)
					}
				} else {
					err = writer.WriteField(key, val)
					if err != nil {
						t.Fatalf("Failed to write field: %v", err)
					}
				}
			}
			err = writer.Close()
			if err != nil {
				t.Fatalf("Failed to close writer: %v", err)
			}
			req, err = http.NewRequest(method, u.String(), &body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
		} else {
			jsonBody, err := json.Marshal(payload)
			if err != nil {
				t.Fatalf("Failed to marshal payload: %v", err)
			}
			req, err = http.NewRequest(method, u.String(), bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
		}
	}

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	return req
}

func TestSaveVideo(t *testing.T) {

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	var data = map[string]string{
		"title":       "Test Video",
		"description": "This is a test video.",
		"file":        "60s.mp4",
	}
	ctx.Request = NewTestRequest(t, http.MethodPost, "/api/v1/video/", data, true)
	ctx.Set("user_id", int64(1))

	// Mock repository, handler, and controller.
	c := v1_controller.NewControllerV1()
	h := v1_handler.NewHandlerV1(repository.NewRepository())

	err := c.PostVideo(ctx, h.PostVideo)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.NoError(t, err)

	// Check the response status code.
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse and validate the response body.
	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

}

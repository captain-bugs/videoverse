package response

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"videoverse/pkg/logbox"
)

type APIController[T any] func(ctx *gin.Context, _ T) error

type CONTENT_TYPE string

const (
	JSON CONTENT_TYPE = "application/json"
)

func write(ctx *gin.Context, status int, payload any) error {
	ctx.Header("Content-Type", string(JSON))
	ctx.Status(status)
	return json.NewEncoder(ctx.Writer).Encode(payload)
}

func GinWrapper[T any](controller APIController[T], handler T) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := controller(ctx, handler); err != nil {
			logbox.NewLogBox().Error().Err(err).Str("path", ctx.Request.URL.Path).Msg("API_ERROR")
			var apiError APIError
			if errors.As(err, &apiError) {
				_ = write(ctx, apiError.StatusCode, apiError)
				ctx.Abort()
				return
			}
			reply := map[string]any{
				"status_code": http.StatusInternalServerError,
				"message":     "Internal Server Error",
				"error":       err.Error(),
			}
			_ = write(ctx, http.StatusInternalServerError, reply)
			ctx.Abort()
			return
		}
	}
}

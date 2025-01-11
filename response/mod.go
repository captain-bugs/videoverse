package response

import (
	"fmt"
	"net/http"
)

type APIError struct {
	ErrorCode  string `json:"error_code,omitempty"`
	StatusCode int    `json:"status_code"`
	Message    any    `json:"message"`
}

func (a APIError) Error() string {
	return fmt.Sprintf("%v", a.Message)
}

func NewAPIError(status int, err error) APIError {
	return APIError{
		StatusCode: status,
		Message:    err.Error(),
	}
}

func UnAuthorized(msg string) APIError {
	return APIError{
		StatusCode: 401,
		Message:    msg,
	}
}

func InvalidRequestBody() APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Message:    "invalid request data",
	}
}

func ErrorsInRequestBody(errors map[string]any) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errors,
	}
}

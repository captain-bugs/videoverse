package response

import (
	"fmt"
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

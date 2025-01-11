package response

import (
	"fmt"
)

type APIError struct {
	ErrorCode       string `json:"error_code,omitempty"`
	StatusCode      int    `json:"status_code"`
	Message         any    `json:"message"`
	InternalMessage error  `json:"internal_message,omitempty"`
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

package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	ApiStatus() int
	ApiMessage() string
	ApiError() string
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func (e *apiError) ApiStatus() int {
	return e.Status
}

func (e *apiError) ApiMessage() string {
	return e.Message
}

func (e *apiError) ApiError() string {
	return e.Error
}

func NewApiErrFromBytes(body []byte) (ApiError, error) {
	var result apiError

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		Status:  statusCode,
		Message: message,
	}
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}
func NewBadRequestError(message string) ApiError {
	return &apiError{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}
func NewInternalServerError(message string) ApiError {
	return &apiError{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

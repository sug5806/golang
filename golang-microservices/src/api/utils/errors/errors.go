package errors

import "net/http"

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

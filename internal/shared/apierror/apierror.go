package apierror

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

const (
	CodeInternal     = "internal_error"
	CodeBadRequest   = "bad_request"
	CodeUnauthorized = "unauthorized"
	CodeForbidden    = "forbidden"
	CodeNotFound     = "not_found"
	CodeConflict     = "conflict"
)

func New(code, message string, status int) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

func Internal(err error) *APIError {
	return New(CodeInternal, "Something went wrong", http.StatusInternalServerError)
}

func BadRequest(msg string) *APIError {
	return New(CodeBadRequest, msg, http.StatusBadRequest)
}

func Unauthorized(msg string) *APIError {
    return New(CodeUnauthorized, msg, http.StatusUnauthorized)
}

func Forbidden(msg string) *APIError {
    return New(CodeForbidden, msg, http.StatusForbidden)
}

func NotFound(msg string) *APIError {
    return New(CodeNotFound, msg, http.StatusNotFound)
}

func Conflict(msg string) *APIError {
    return New(CodeConflict, msg, http.StatusConflict)
}

func From(err error) *APIError {
    if err == nil {
        return nil
    }

    if apiErr, ok := err.(*APIError); ok {
        return apiErr
    }

    return Internal(err)
}

func Write(w http.ResponseWriter, err error) {
    apiErr := From(err)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(apiErr.Status)

    _ = json.NewEncoder(w).Encode(apiErr)
}

func WriteGin(c *gin.Context, err error) {
    apiErr := From(err)
    c.AbortWithStatusJSON(apiErr.Status, apiErr)
}
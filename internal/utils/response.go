package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	ErrCodeBadRequest     ErrorCode = "BAD_REQUEST"
	ErrCodeNotFound       ErrorCode = "NOT_FOUND"
	ErrCodeInternal       ErrorCode = "INTERNAL_ERROR"
	ErrCodeConflict       ErrorCode = "CONFLICT"
	ErrCodeInvalidInput   ErrorCode = "INVALID_INPUT"
	ErrCodeEmailExists    ErrorCode = "EMAIL_ALREADY_EXISTS"
	ErrCodeUnauthorized   ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden      ErrorCode = "FORBIDDEN"
	ErrCodeInternalServer ErrorCode = "INTERNAL_SERVER_ERROR"
)

type AppError struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
	Err     string    `json:"error"`
}

func (e *AppError) Error() string {
	return ""
}

func NewError(message string, code ErrorCode) error {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func WrapError(err error, message string, code ErrorCode) error {
	return &AppError{
		Message: message,
		Code:    code,
		Err:     err.Error(),
	}
}

func ReponseErr(ctx *gin.Context, err error) {
	if appError, ok := err.(*AppError); ok {
		httpStatus := httpStatusFromCode(ErrorCode(appError.Code))

		response := gin.H{

			"code":  appError.Code,
			"error": appError.Message,
		}

		if appError.Err != "" {
			response["detail"] = appError.Err
		}

		ctx.JSON(httpStatus, response)
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
		"code":    string(ErrCodeInternalServer),
		"error":   err.Error(),
	})
}

func ReponseSuccess(ctx *gin.Context, status int, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"data":   data,
		"status": "success",
	})
}

func ResponseStatus(ctx *gin.Context, status int) {
	ctx.JSON(status, gin.H{})
}

func httpStatusFromCode(code ErrorCode) int {
	switch code {
	case ErrCodeBadRequest:
		return http.StatusBadRequest
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeInternal:
		return http.StatusInternalServerError
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeInvalidInput:
		return http.StatusBadRequest
	case ErrCodeEmailExists:
		return http.StatusConflict
	case ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrCodeForbidden:
		return http.StatusForbidden
	// Add more cases as needed
	default:
		return http.StatusInternalServerError
	}

}

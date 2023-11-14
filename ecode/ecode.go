package ecode

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func New(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrRefreshDB = New(500, "failed to refresh DB")
	ErrWriteDB   = New(500, "failed to write data in DB")

	ErrRefreshRedis = New(500, "failed to refresh redis")
	ErrWriteRedis   = New(500, "failed to write data in REDS")

	ErrUnauthorized = New(401, "unauthorized")
)

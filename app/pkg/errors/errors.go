package errors

import (
	"github.com/postmy/app/constant"
)

// IError is errors interface
type IError interface {
	GetCustomCode() int
	GetHTTPCode() int
	GetCustomMessage() string
}

// Error is error model
type Error struct {
	CustomCode    int
	HTTPCode      int
	CustomMessage string
}

// New is errors constructor
func New(customCode int) IError {
	errorCode := constant.ErrorCodes.FindErrorCodeByCustomCode(customCode)

	return &Error{
		CustomCode:    customCode,
		HTTPCode:      errorCode.HTTPCode,
		CustomMessage: errorCode.CustomMessage,
	}
}

// GetCustomCode is function to get custom code
func (e *Error) GetCustomCode() int {
	return e.CustomCode
}

// GetHTTPCode is function to get http code
func (e *Error) GetHTTPCode() int {
	return e.HTTPCode
}

// GetCustomMessage is function to get custom message
func (e *Error) GetCustomMessage() string {
	return e.CustomMessage
}

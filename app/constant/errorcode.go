package constant

import (
	"github.com/postmy/app/pkg/enum"
	"net/http"
)

// error code client error response (4xxxx)
const (
	ErrorBindingParameter  = 40001
	ErrorValidatingPayload = 40002
	ErrorPasswordNotMatch  = 40003
	ErrorDataNotFoundDB    = 40004
	ErrorDataNotFoundCache = 40005
	ErrorUserIsExist       = 40006
	ErrorUserDoesntExist   = 40007
	ErrorIncorrectOTP      = 40008
	ErrorIncorrectPassword = 40009
	ErrorUnauthorized      = 40010
	ErrorLoginNeedOTP      = 40011
	ErrorNoPermission      = 40012
	ErrorDataConversion    = 40013
)

// error code server error response (5xxxx)
const (
	ErrorInternaly    = 50001
	ErrorSendingEmail = 50002
)

var (
	// ErrorCodes is list enumeration of custom error code
	ErrorCodes = enum.Enum{
		ErrorCode: []enum.ErrorCode{
			{CustomCode: ErrorBindingParameter, HTTPCode: http.StatusBadRequest, CustomMessage: "error binding parameter"},
			{CustomCode: ErrorValidatingPayload, HTTPCode: http.StatusBadRequest, CustomMessage: "error validating payload"},
			{CustomCode: ErrorPasswordNotMatch, HTTPCode: http.StatusBadRequest, CustomMessage: "password and confirmation password did not match"},
			{CustomCode: ErrorDataNotFoundDB, HTTPCode: http.StatusBadRequest, CustomMessage: "data not found on database"},
			{CustomCode: ErrorDataNotFoundCache, HTTPCode: http.StatusBadRequest, CustomMessage: "data not found on cache"},
			{CustomCode: ErrorUserIsExist, HTTPCode: http.StatusBadRequest, CustomMessage: "user is exist on database"},
			{CustomCode: ErrorUserDoesntExist, HTTPCode: http.StatusBadRequest, CustomMessage: "user does not exist on database"},
			{CustomCode: ErrorIncorrectOTP, HTTPCode: http.StatusBadRequest, CustomMessage: "otp is incorrect"},
			{CustomCode: ErrorIncorrectPassword, HTTPCode: http.StatusBadRequest, CustomMessage: "password is incorrect"},
			{CustomCode: ErrorUnauthorized, HTTPCode: http.StatusForbidden, CustomMessage: "unauthorized, please login first"},
			{CustomCode: ErrorLoginNeedOTP, HTTPCode: http.StatusForbidden, CustomMessage: "login need otp"},
			{CustomCode: ErrorNoPermission, HTTPCode: http.StatusForbidden, CustomMessage: "no access permission"},
			{CustomCode: ErrorDataConversion, HTTPCode: http.StatusBadRequest, CustomMessage: "invalid data conversion"},

			{CustomCode: ErrorInternaly, HTTPCode: http.StatusInternalServerError, CustomMessage: "internal server error, please try again later"},
			{CustomCode: ErrorSendingEmail, HTTPCode: http.StatusInternalServerError, CustomMessage: "error while sending email"},
		},
	}
)

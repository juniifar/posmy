package enum

// Enum is enumeration struct
type Enum struct {
	ErrorCode []ErrorCode
}

// ErrorCode is mapping of custom error code
type ErrorCode struct {
	CustomCode    int
	HTTPCode      int
	CustomMessage string
}

// FindErrorCodeByCustomCode is funtion to find error code by custom code on error code enumeration
func (enum Enum) FindErrorCodeByCustomCode(customCode int) ErrorCode {
	errorCode := ErrorCode{}

	for _, val := range enum.ErrorCode {
		if customCode == val.CustomCode {
			return val
		}
	}

	return errorCode
}

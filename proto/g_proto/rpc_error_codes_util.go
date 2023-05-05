package g_proto

import (
	"fmt"
	"strings"
)

// normal code NewXXX
func NewRpcError(code int32, message string) (err *GTVRpcError) {
	httpCode := int32(code) / 1000
	if httpCode >= 1000 {
		httpCode = httpCode / 10
	}

	message = strings.TrimSpace(message)
	if name, ok := GTVRpcErrorCodes_name[int32(code)]; ok {
		if message == "" {
			message = name
		}

		if httpCode <= 500 {
			err = &GTVRpcError{
				Code:      httpCode,
				ErrorCode: code,
				Message:   fmt.Sprintf("%s: %s", name, message),
				WaitFor:   10,
			}
		} else {
			err = &GTVRpcError{
				Code:      httpCode,
				ErrorCode: int32(code),
				Message:   fmt.Sprintf("%s: %s", name, name),
				WaitFor:   30,
			}
		}
	} else {
		code = int32(GTVRpcErrorCodes_ERROR_INTERNAL)
		httpCode := int32(code) / 1000
		if httpCode >= 1000 {
			httpCode = httpCode / 10
		}

		err = &GTVRpcError{
			Code:      httpCode,
			ErrorCode: code,
			Message:   fmt.Sprintf("INTERNAL_SERVER_ERROR: code = %d, message = %s", code, message),
			WaitFor:   60,
		}
	}

	return
}

// normal code NewXXX
func NewRpcError2(code GTVRpcErrorCodes) (err *GTVRpcError) {
	httpCode := int32(code) / 1000
	if httpCode >= 1000 {
		httpCode = int32(code) / 10
	}

	waitingFor := int32(10)
	errName := ""
	if name, ok := GTVRpcErrorCodes_name[int32(code)]; ok {
		errName = name
		if code <= GTVRpcErrorCodes_ERROR_OTHER {
			waitingFor = int32(10)
		} else {
			waitingFor = int32(60)
		}
	} else {
		code = GTVRpcErrorCodes_ERROR_INTERNAL
		waitingFor = int32(120)
		errName = "INTERNAL_SERVER_ERROR"
	}

	err = &GTVRpcError{
		Code:      httpCode,
		ErrorCode: int32(code),
		Message:   fmt.Sprintf("%s: %s", errName, errName),
		WaitFor:   waitingFor,
	}

	return
}

// Impl error interface
func (e *GTVRpcError) IsOK() bool {
	if e == nil {
		return true
	}
	return e.GetErrorCode() == int32(GTVRpcErrorCodes_ERROR_CODE_OK)
}

// Impl error interface
func (e *GTVRpcError) Error() string {
	return fmt.Sprintf("rpc error: code = %d desc = %s", e.GetErrorCode(), e.GetMessage())
}

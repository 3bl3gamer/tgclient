package mtproto

import (
	"github.com/ansel1/merry"
)

// https://core.telegram.org/api/errors
func IsError(obj TL, message string) bool {
	err, ok := obj.(TL_rpc_error)
	return ok && err.error_message == message
}

func IsErrorType(obj TL, code int32) bool {
	err, ok := obj.(TL_rpc_error)
	return ok && err.error_code == code
}

// func String(obj TL) string {

// }

func WrongRespError(obj TL) error {
	_type := "response"
	if _, ok := obj.(TL_rpc_error); ok {
		_type = "error"
	}
	return merry.Errorf("got unexpected %s: %T%+v", _type, obj, obj).WithStackSkipping(1)
}

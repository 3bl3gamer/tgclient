package mtproto

import (
	"fmt"

	"github.com/ansel1/merry"
)

// https://core.telegram.org/api/errors
func IsError(obj TL, message string) bool {
	err, ok := obj.(TL_rpc_error)
	return ok && err.ErrorMessage == message
}

func IsErrorType(obj TL, code int32) bool {
	err, ok := obj.(TL_rpc_error)
	return ok && err.ErrorCode == code
}

func Sprint(obj TL) string {
	return fmt.Sprintf("%T%+v", obj, obj)
}

func UnexpectedTL(name string, obj TL) string {
	return fmt.Sprintf("unexpected " + name + ": " + Sprint(obj))
}

func WrongRespError(obj TL) error {
	_type := "response"
	if _, ok := obj.(TL_rpc_error); ok {
		_type = "error"
	}
	return merry.Errorf(UnexpectedTL(_type, obj)).WithStackSkipping(1)
}

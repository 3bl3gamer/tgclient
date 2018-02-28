package mtproto

import (
	"fmt"
	"strings"

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

func IsClosedConnErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "use of closed network connection")
}

func Sprint(obj TL) string {
	return fmt.Sprintf("%#v", obj)
}

func UnexpectedTL(name string, obj TL) string {
	return fmt.Sprint("unexpected " + name + ": " + Sprint(obj))
}

func WrongRespError(obj TL) error {
	_type := "response"
	if _, ok := obj.(TL_rpc_error); ok {
		_type = "error"
	}
	return merry.Errorf(UnexpectedTL(_type, obj)).WithStackSkipping(1)
}

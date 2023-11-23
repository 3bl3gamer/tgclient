package mtproto

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ansel1/merry/v2"
)

// https://core.telegram.org/api/errors
func IsError(obj TL, message string) bool {
	err, ok := obj.(TL_rpcError)
	return ok && err.ErrorMessage == message
}

func IsErrorType(obj TL, code int32) bool {
	err, ok := obj.(TL_rpcError)
	return ok && err.ErrorCode == code
}

const FloodWaitErrPerfix = "FLOOD_WAIT_"

func IsFloodError(obj TL) (time.Duration, bool) {
	if err, ok := obj.(TL_rpcError); ok && strings.HasPrefix(err.ErrorMessage, FloodWaitErrPerfix) {
		secs, _ := strconv.ParseInt(err.ErrorMessage[len(FloodWaitErrPerfix):], 10, 64)
		if secs <= 0 {
			secs = 1
		}
		return time.Duration(secs) * time.Second, true
	}
	return 0, false
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
	if _, ok := obj.(TL_rpcError); ok {
		_type = "error"
	}
	return merry.WrapSkipping(errors.New(UnexpectedTL(_type, obj)), 1)
}

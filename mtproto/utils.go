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

func IsErrorMessage(obj any, message string) bool {
	if val, ok := unwrapUnexpectedTypeErrValue(obj); ok {
		obj = val
	}
	err, ok := obj.(TL_rpcError)
	return ok && err.ErrorMessage == message
}

const FloodWaitErrPerfix = "FLOOD_WAIT_"
const FloodPremiumWaitErrPerfix = "FLOOD_PREMIUM_WAIT_"

func IsFloodError(tlOrErr any) (time.Duration, bool) {
	var secs int64

	if val, ok := unwrapUnexpectedTypeErrValue(tlOrErr); ok {
		tlOrErr = val
	}

	err, ok := tlOrErr.(TL_rpcError)
	if !ok {
		return 0, false
	}

	if strings.HasPrefix(err.ErrorMessage, FloodWaitErrPerfix) {
		secs, _ = strconv.ParseInt(err.ErrorMessage[len(FloodWaitErrPerfix):], 10, 64)
	} else if strings.HasPrefix(err.ErrorMessage, FloodPremiumWaitErrPerfix) {
		secs, _ = strconv.ParseInt(err.ErrorMessage[len(FloodPremiumWaitErrPerfix):], 10, 64)
	} else {
		return 0, false
	}

	if secs <= 0 {
		secs = 1
	}

	return time.Duration(secs) * time.Second, true
}

// https://core.telegram.org/mtproto/service_messages_about_messages#notice-of-ignored-error-message
func IsWrongClientTimeError(tlOrErr any) bool {
	if val, ok := unwrapUnexpectedTypeErrValue(tlOrErr); ok {
		tlOrErr = val
	}

	if msg, ok := tlOrErr.(TL_badMsgNotification); ok {
		return msg.ErrorCode == 16 || msg.ErrorCode == 17
	}
	return false
}

func IsMsgSeqnoTooLowError(tlOrErr any) bool {
	if val, ok := unwrapUnexpectedTypeErrValue(tlOrErr); ok {
		tlOrErr = val
	}

	if msg, ok := tlOrErr.(TL_badMsgNotification); ok {
		return msg.ErrorCode == 32
	}
	return false
}

func IsMsgSeqnoTooHighError(tlOrErr any) bool {
	if val, ok := unwrapUnexpectedTypeErrValue(tlOrErr); ok {
		tlOrErr = val
	}

	if msg, ok := tlOrErr.(TL_badMsgNotification); ok {
		return msg.ErrorCode == 33
	}
	return false
}

func unwrapUnexpectedTypeErrValue(obj any) (TL, bool) {
	if err, ok := obj.(error); ok {
		err = errors.Unwrap(err)
		if typeErr, ok := err.(UnexpectedTypeError); ok {
			return typeErr.Value, true
		}
	}
	return nil, false
}

func IsClosedConnErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "use of closed network connection")
}

func Sprint(obj TL) string {
	return fmt.Sprintf("%#v", obj)
}

func UnexpectedTL(name string, obj TL) string {
	return "unexpected " + name + ": " + Sprint(obj)
}

type UnexpectedTypeError struct {
	Value TL
}

func (r UnexpectedTypeError) Error() string {
	_type := "response"
	if _, ok := r.Value.(TL_rpcError); ok {
		_type = "error"
	}
	return UnexpectedTL(_type, r.Value)
}

func WrongRespError(obj TL) error {
	return merry.WrapSkipping(UnexpectedTypeError{Value: obj}, 1)
}

func UnwrapWrongRespError[T TL](err error) (T, bool) {
	err = errors.Unwrap(err)
	if typeErr, ok := err.(UnexpectedTypeError); ok {
		resp, ok := typeErr.Value.(T)
		return resp, ok
	}
	var zero T
	return zero, false
}

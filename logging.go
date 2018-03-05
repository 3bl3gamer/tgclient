package mtproto

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/ansel1/merry"
)

type LogLevel int

const (
	ERROR LogLevel = iota
	WARN
	INFO
	DEBUG
)

type LogHandler interface {
	Log(LogLevel, error, string, ...interface{})
	Message(bool, TL, int64)
}

type SimpleLogHandler struct{}

func (h SimpleLogHandler) Log(level LogLevel, err error, msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	if err != nil {
		msg += ":\n" + merry.Details(err)
	}
	switch level {
	case ERROR:
		msg = "[ERROR] " + msg
	case WARN:
		msg = "[WARN] " + msg
	case INFO:
		msg = "[INFO] " + msg
	case DEBUG:
		msg = "\033[90m[DEBUG] " + msg + "\033[0m"
	}
	log.Print(msg)
}

func tlName(obj interface{}) string {
	return reflect.TypeOf(obj).Name()
}

func (h SimpleLogHandler) Message(isIncoming bool, msg TL, id int64) {
	var text string
	switch x := msg.(type) {
	case TL_msg_container:
		names := make([]string, len(x.Items))
		for i, item := range x.Items {
			names[i] = tlName(item)
		}
		text = tlName(x) + " -> [" + strings.Join(names, ", ") + "]"
	case TL_rpc_result:
		text = tlName(x) + " -> " + tlName(x.obj)
	default:
		text = tlName(x)
	}
	if isIncoming {
		text = ">>> " + text
	} else {
		text = "<<< " + text + fmt.Sprintf(" (#%d)", id)
	}
	h.Log(DEBUG, nil, text)
}

type Logger struct {
	hnd LogHandler
}

func (l Logger) Error(err error, msg string, args ...interface{}) {
	l.hnd.Log(ERROR, err, msg, args...)
}

func (l Logger) Warn(msg string, args ...interface{}) {
	l.hnd.Log(WARN, nil, msg, args...)
}

func (l Logger) Info(msg string, args ...interface{}) {
	l.hnd.Log(INFO, nil, msg, args...)
}

func (l Logger) Debug(msg string, args ...interface{}) {
	l.hnd.Log(DEBUG, nil, msg, args...)
}

func (l Logger) Message(isIncoming bool, message TL, id int64) {
	l.hnd.Message(isIncoming, message, id)
}

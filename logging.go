package mtproto

import (
	"fmt"
	"log"

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

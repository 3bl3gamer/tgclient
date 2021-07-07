package mtproto

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/ansel1/merry"
	"github.com/fatih/color"
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

func (h SimpleLogHandler) TLName(obj interface{}) string {
	return reflect.TypeOf(obj).Name()
}

func (h SimpleLogHandler) StringifyLog(level LogLevel, err error, msg string, args ...interface{}) string {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	if err != nil {
		msg += ":\n" + merry.Details(err)
	}
	return msg
}

func (h SimpleLogHandler) AddLevelPrevix(level LogLevel, text string) string {
	switch level {
	case ERROR:
		text = "[ERROR] " + text
	case WARN:
		text = "[WARN] " + text
	case INFO:
		text = "[INFO] " + text
	case DEBUG:
		text = "[DEBUG] " + text
	}
	return text
}

func (h SimpleLogHandler) StringifyMessage(isIncoming bool, msg TL, id int64) string {
	var text string
	switch x := msg.(type) {
	case TL_msg_container:
		names := make([]string, len(x.Items))
		for i, item := range x.Items {
			names[i] = h.TLName(item)
		}
		text = h.TLName(x) + " -> [" + strings.Join(names, ", ") + "]"
	case TL_rpc_result:
		text = h.TLName(x) + " -> " + h.TLName(x.obj)
	default:
		text = h.TLName(x)
	}
	if isIncoming {
		text = ">>> " + text
	} else {
		text = "<<< " + text + fmt.Sprintf(" (#%d)", id)
	}
	return text
}

func (h SimpleLogHandler) Log(level LogLevel, err error, msg string, args ...interface{}) {
	text := h.StringifyLog(level, err, msg, args...)
	text = h.AddLevelPrevix(level, text)
	log.Print(text)
}

func (h SimpleLogHandler) Message(isIncoming bool, msg TL, id int64) {
	h.Log(DEBUG, nil, h.StringifyMessage(isIncoming, msg, id))
}

// ColorLogHandler adds log line colors depending on severity level
type ColorLogHandler struct {
	SimpleLogHandler
	StdLogger *log.Logger
}

func NewColorLogHandler() *ColorLogHandler {
	// color.Error pases text via go-colorable which converts color escape sequences into ansi color on windows
	return &ColorLogHandler{StdLogger: log.New(color.Error, "", log.LstdFlags)}
}

var debugLogColor = color.New(color.FgHiBlack).SprintFunc()
var warnLogColor = color.New(color.FgHiYellow).SprintFunc()
var errorLogColor = color.New(color.FgRed).SprintFunc()

func (h ColorLogHandler) AddLevelColor(level LogLevel, text string) string {
	switch level {
	case DEBUG:
		return debugLogColor(text)
	case WARN:
		return warnLogColor(text)
	case ERROR:
		return errorLogColor(text)
	}
	return text
}

func (h ColorLogHandler) Log(level LogLevel, err error, msg string, args ...interface{}) {
	text := h.StringifyLog(level, err, msg, args...)
	text = h.AddLevelPrevix(level, text)
	text = h.AddLevelColor(level, text)
	h.StdLogger.Print(text)
}

func (h ColorLogHandler) Message(isIncoming bool, msg TL, id int64) {
	h.Log(DEBUG, nil, h.StringifyMessage(isIncoming, msg, id))
}

type Logger struct {
	Hnd LogHandler
}

func (l Logger) Error(err error, msg string, args ...interface{}) {
	l.Hnd.Log(ERROR, err, msg, args...)
}

func (l Logger) Warn(msg string, args ...interface{}) {
	l.Hnd.Log(WARN, nil, msg, args...)
}

func (l Logger) Info(msg string, args ...interface{}) {
	l.Hnd.Log(INFO, nil, msg, args...)
}

func (l Logger) Debug(msg string, args ...interface{}) {
	l.Hnd.Log(DEBUG, nil, msg, args...)
}

func (l Logger) Message(isIncoming bool, message TL, id int64) {
	l.Hnd.Message(isIncoming, message, id)
}

package mylogger

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type logLevel uint16

const (
	DEFAULT logLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// mylog 结构体
type Mylog interface {
	Debug(string)
	Info(string)
	Trace(string)
	Warning(string)
	Error(string)
	Fatal(string)
}

func parseString(l string) (logLevel, error) {
	switch strings.ToLower(l) {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "trace":
		return TRACE, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		return DEFAULT, errors.New("指定的日志界别无效")
	}

}

// parseLevelToString ...
func parseLevelToString(lv logLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case TRACE:
		return "TRACE"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		panic("无效的日志级别")
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, fileName, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("get runtime caller info failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	return
}

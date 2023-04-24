package mylogger

import (
	"fmt"
	"strings"
)

func NewconsoleLogger(s string) consoleLogger {
	temps := strings.ToLower(s)
	n := transLevel(temps)

	return consoleLogger{
		level: n,
	}
}

func (c consoleLogger) log(level logLevel, format string, a ...interface{}) {
	if c.enable(level) {
		msg := fmt.Sprintf(format, a...)
		fileName, funcName, lineNo := getCaller()
		fmt.Printf("[%s] [%s] %s:%s:%d %s\n", unTransLevel(level), getTime(), fileName, funcName, lineNo, msg)
	}
}

func (c consoleLogger) enable(mylevel logLevel) bool {
	return c.level <= mylevel
}

func (c consoleLogger) Debug(s string, a ...interface{}) {
	c.log(DebugLevel, s, a...)
}

func (c consoleLogger) Info(s string, a ...interface{}) {
	c.log(InfoLevel, s, a...)
}

func (c consoleLogger) Trace(s string, a ...interface{}) {
	c.log(TraceLevel, s, a...)
}

func (c consoleLogger) Warning(s string, a ...interface{}) {
	c.log(WarningLevel, s, a...)
}

func (c consoleLogger) Error(s string, a ...interface{}) {
	c.log(ErrorLevel, s, a...)
}

func (c consoleLogger) Fatal(s string, a ...interface{}) {
	c.log(FatalLevel, s, a...)
}

package mylogger

//向终端上写日志

import (
	"fmt"
	"time"
)

// mylog 结构体
type MyConsolelog struct {
	level logLevel
}

// NewLogger 构造函数
func NewConsolelog(s string) MyConsolelog {
	l, err := parseString(s)
	if err != nil {
		panic(err)
	}
	return MyConsolelog{
		level: l,
	}
}

// a ...interface{} 切片传入
func (l MyConsolelog) log(loglevel logLevel, format string, a ...interface{}) {
	if l.enable(loglevel) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05.000")
		fmt.Printf("[%s] [%s] 这是一条日志\n", now, msg)
	}
}

// Debug 函数
func (l MyConsolelog) Debug(s string) {
	l.log(DEBUG, "debug")
}

// Info 函数
func (l MyConsolelog) Info(s string) {
	l.log(INFO, "debug")
}

// Trace 函数
func (l MyConsolelog) Trace(s string) {
	l.log(TRACE, "debug")
}

// Warning 函数
func (l MyConsolelog) Warning(s string) {
	l.log(WARNING, "debug")
}

func (l MyConsolelog) Error(s string) {
	l.log(ERROR, "debug")
}

// Fatal 函数
func (l MyConsolelog) Fatal(s string) {
	l.log(FATAL, "debug")
}

func (l MyConsolelog) enable(L logLevel) bool {
	return l.level <= L
}


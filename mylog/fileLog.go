package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//向文件中写日志
type FileLogger struct {
	level       string
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

func NewFileLoger(levelString, fp, fn string, maxSize int64) *FileLogger {
	f := &FileLogger{
		level:       levelString,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err := f.initFile()
	if err != nil {
		panic(err)
	}
	return f
}

func (f *FileLogger) initFile() error {
	fullFilename := path.Join(f.filePath, f.fileName)
	fp, err := os.OpenFile(fullFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err=%v\n", err)
		return err
	}
	f.fileObj = fp
	ferrp, err := os.OpenFile(fullFilename+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err=%v\n", err)
		return err
	}
	f.errFileObj = ferrp
	return nil
}

func (f *FileLogger) checkSize(fp *os.File) bool {
	fileInfo, err := fp.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err=%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) log(lv logLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15.04.05.000")
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, parseLevelToString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, parseLevelToString(lv), fileName, funcName, lineNo, msg)
		}
		if f.checkSize(f.fileObj) {
			f.fileObj.Close()
			fullFilename := path.Join(f.filePath, f.fileName)
			err := os.Rename(fullFilename, fullFilename+"."+now+".bak")
			if err != nil {
				fmt.Printf("rename log file to bak file failed, err=%v\n", err)
				return
			}
			fp, err := os.OpenFile(fullFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("open log file failed, err=%v\n", err)
				return
			}
			f.fileObj = fp

			if f.checkSize(f.errFileObj) {
				f.errFileObj.Close()
				err = os.Rename(fullFilename+".err", fullFilename+".err"+now+".bak")
				if err != nil {
					fmt.Printf("rename err log file to bak file failed, err=%v\n", err)
					return
				}
				ferrp, err := os.OpenFile(fullFilename+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Printf("open err log file failed, err=%v\n", err)
					return
				}
				f.errFileObj = ferrp
			}
		}
	}
}

func (f *FileLogger) enable(lv logLevel) bool {
	fl, err := parseString(f.level)
	if err != nil {
		fmt.Printf("parse file level from string to int64 failed, err=%v\n", err)
		return false
	}
	return fl <= lv
}

// Debug 函数
func (f *FileLogger) Debug(s string) {
	f.log(DEBUG, "debug")
}

// Info 函数
func (f *FileLogger) Info(s string) {
	f.log(INFO, "debug")
}

// Trace 函数
func (f *FileLogger) Trace(s string) {
	f.log(TRACE, "debug")
}

// Warning 函数
func (f *FileLogger) Warning(s string) {
	f.log(WARNING, "debug")
}

// Error 函数
func (f *FileLogger) Error(s string) {
	f.log(ERROR, "debug")
}

// Fatal 函数
func (f *FileLogger) Fatal(s string) {
	f.log(FATAL, "debug")
}

package mylogger

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

func NewFileLogger(levelStr, fpath, fname string, maxSize int64) *fileLogger {
	tempStr := strings.ToLower(levelStr)
	n := transLevel(tempStr)

	if !strings.HasSuffix(fname, ".log") {
		fname = fname + ".log"
	}

	filepath := path.Join(fpath, fname)
	fileObj, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open log file faild, err=%v\n", err)
		panic(err)
	}
	errfileObj, err := os.OpenFile(filepath+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open err log file faild, err=%v\n", err)
		panic(err)
	}

	return &fileLogger{
		level:       n,
		filePath:    fpath,
		fileName:    fname,
		fileObj:     fileObj,
		errFileObj:  errfileObj,
		maxFileSize: maxSize,
	}
}

func (f *fileLogger) enable(mylevel logLevel) bool {
	return f.level <= mylevel
}

func (f *fileLogger) checkSize(fp *os.File) bool {
	fileInfo, err := fp.Stat()
	if err != nil {
		fmt.Printf("Obtain log file state faild, err=%v\n", err)
		panic(err)
	}
	if fileInfo.Size() > f.maxFileSize {
		return true
	}
	return false
}

func (f *fileLogger) splitFile(fp *os.File) (nFp *os.File) {
	fileInfo, _ := fp.Stat()
	fileName := fileInfo.Name()
	oldFileName := path.Join(f.filePath, fileName)
	nowStr := time.Now().Format("20060102150405.000")
	fp.Close()
	os.Rename(oldFileName, oldFileName+"-"+nowStr+".bak")
	nFp, err := os.OpenFile(oldFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Split - Create new file faild, err=%v\n", err)
		panic(err)
	}
	return nFp
}

func (f *fileLogger) log(level logLevel, format string, a ...interface{}) {
	if f.enable(level) {
		msg := fmt.Sprintf(format, a...)
		sw_split := f.checkSize(f.fileObj)
		if sw_split {
			f.fileObj = f.splitFile(f.fileObj)
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] %s %s\n", unTransLevel(level), getTime(), getCaller(), msg)
		if level >= ErrorLevel {
			sw_errSplit := f.checkSize(f.errFileObj)
			if sw_errSplit {
				f.errFileObj = f.splitFile(f.errFileObj)
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] %s %s\n", unTransLevel(level), getTime(), getCaller(), msg)
		}
	}
}

func (f *fileLogger) Debug(s string, a ...interface{}) {
	f.log(DebugLevel, s, a...)
}

func (f *fileLogger) Info(s string, a ...interface{}) {
	f.log(InfoLevel, s, a...)
}

func (f *fileLogger) Trace(s string, a ...interface{}) {
	f.log(TraceLevel, s, a...)
}

func (f *fileLogger) Warning(s string, a ...interface{}) {
	f.log(WarningLevel, s, a...)
}

func (f *fileLogger) Error(s string, a ...interface{}) {
	f.log(ErrorLevel, s, a...)
}

func (f *fileLogger) Fatal(s string, a ...interface{}) {
	f.log(FatalLevel, s, a...)
}

func (f *fileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

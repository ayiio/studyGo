package taillog

// 用于收集日志

import (
	"time"

	"github.com/hpcloud/tail"
)

var (
	tails *tail.Tail
)

func InitTail(path string) (err error) {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Follow:    true,
		Poll:      true,
	}
	tails, err = tail.TailFile(path, config)
	return
}

func TailFile() (line *tail.Line) {
	select {
	case line = <-tails.Lines:
		return
	default:
		time.Sleep(time.Second)
	}
	return
}


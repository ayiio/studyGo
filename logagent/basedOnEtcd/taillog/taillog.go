package taillog

// 用于收集日志

import (
	"context"
	"fmt"
	"logagent/kafka"
	"time"

	"github.com/hpcloud/tail"
)

//一个日志收集任务的结构体
type TailTask struct {
	path      string     //该task收集的日志路径
	topic     string     //该task收集的日志将要放到的kafka目标
	instances *tail.Tail //tailf打开的文件对象
	//使用context实现控制子goroutine退出
	taskCtx    context.Context    //上下文context
	taskCancel context.CancelFunc //cancel func
}

//任务的构造函数
func newTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		taskCtx:    ctx,
		taskCancel: cancel,
	}
	tailObj.init_task() //实例化日志文件对象
	return
}

//使用tail工具打开日志文件
func (t *TailTask) init_task() {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Follow:    true,
		Poll:      true,
	}
	var err error
	t.instances, err = tail.TailFile(t.path, config) //赋值任务结构体中的instances
	if err != nil {
		fmt.Printf("tail file failed, err=%v\n", err)
	}
	//goroutine执行的函数退出后，goroutine退出
	go t.readLog() //后台从instances中采集日志发送到kafka
}

//采集日志
func (t *TailTask) readLog() {
	for {
		select {
		case <-t.taskCtx.Done(): //子goroutine上下文收到Done信号
			fmt.Printf("tail task:%s_%s exist.\n", t.path, t.topic)
			return
		case line := <-t.instances.Lines:
			//将信息发送到chan，在其他包中处理chan，实现异步
			kafka.PutChan(t.topic, line.Text)
		default:
			time.Sleep(time.Millisecond * 15)
		}
	}
}

package taillog

import (
	"fmt"
	"logagent/etcd"
)

type tailLogMgr struct {
	logEntry []*etcd.LogEntry     //存放获取到的所有的logEntry
	taskMap  map[string]*TailTask //热加载管理预留
}

var taskMgr *tailLogMgr //全局task管理者，存放当前日志收集项的配置信息

func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntry: logEntryConf,
	}
	for _, logEntry := range logEntryConf {
		newTailTask(logEntry.Path, logEntry.Topic)
	}
}

// task管理预留方法
func TaskManager() {
	fmt.Println(taskMgr.taskMap)
}

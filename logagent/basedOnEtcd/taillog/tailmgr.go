package taillog

import (
	"fmt"
	"logagent/etcd"
	"time"
)

//tailTask管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry      //存放获取到的所有的logEntry
	taskMap     map[string]*TailTask  //key-value方式存放日志任务，便于管理
	newConfChan chan []*etcd.LogEntry //新的配置
}

var taskMgr *tailLogMgr //全局task管理者，存放当前日志收集项的配置信息

func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntry:    logEntryConf,
		taskMap:     make(map[string]*TailTask, 32),
		newConfChan: make(chan []*etcd.LogEntry), //无缓冲区通道，有最新配置就更新到chan
	}
	for _, logEntry := range logEntryConf {
		taskObj := newTailTask(logEntry.Path, logEntry.Topic)
		//记录原始task配置用于后期管理，使用log路径为key
		orikey := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		taskMgr.taskMap[orikey] = taskObj
	}
	go taskMgr.watchNewConfChan() //异步等待新的配置
}

// 监听newConfChan，对应以下操作
//1.配置新增/配置更新
//2.配置删除
func (t *tailLogMgr) watchNewConfChan() {
	for {
		select {
		case newConf := <-t.newConfChan: //通道无值时走default sleep
			//1.配置新增/配置更新
			for _, nConf := range newConf {
				newkey := fmt.Sprintf("%s_%s", nConf.Path, nConf.Topic)
				//1.新增配置
				if _, exist := t.taskMap[newkey]; !exist { //出现新配置项
					newTailObj := newTailTask(nConf.Path, nConf.Topic)
					//保存新的配置项
					t.taskMap[newkey] = newTailObj
				}
			}
			//2.删除配置
			//设置辅助map，用于判断删除项，map不能复制，不可以使用tmpDelMap:=t.taskMap
			var tmpDelMap = make(map[string]*TailTask, len(t.taskMap))
			for tmpK, tmpV := range t.taskMap {
				var tmpVV = *tmpV
				tmpDelMap[tmpK] = &tmpVV
			}
			for _, nConf := range newConf {
				newkey := fmt.Sprintf("%s_%s", nConf.Path, nConf.Topic)
				delete(tmpDelMap, newkey)
			}
			for delTaskKey := range tmpDelMap { //辅助map中剩余项为需要删除的配置
				t.taskMap[delTaskKey].taskCancel() //使用context的cancelFunc关掉删除的配置项
				delete(t.taskMap, delTaskKey)      //更新已有的配置项
			}
		default:
			time.Sleep(time.Microsecond * 10)
		}
	}
}

//向外部暴漏本包私有字段
func GetNewConfToChan() chan<- []*etcd.LogEntry {
	return taskMgr.newConfChan //初始化在init中，避免未初始化时的空指针问题
}

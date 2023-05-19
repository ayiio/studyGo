package main

import (
	"fmt"
	"logagent/logagent/config"
	"logagent/logagent/kafka"
	"logagent/logagent/taillog"
)

// loagent入口

func main() {
	// 0.load config.ini
	param, err := config.InitParam()
	if err != nil {
		fmt.Printf("init file load failed, err=%v\n", err)
		return
	}
	fmt.Println("# load init file success.")
	// 1.初始化kafka连接
	err = kafka.InitKafka([]string{param.Address})
	if err != nil {
		fmt.Printf("init Kafka failed, err=%v\n", err)
		return
	}
	fmt.Println("# init kafka success.")
	// 2.打开日志文件准备收集日志
	err = taillog.InitTail(param.Path)
	if err != nil {
		fmt.Printf("init Tail failed, err=%v\n", err)
	}
	fmt.Println("# init tail success.")
	line := taillog.TailFile()
	pid, offset, err := kafka.SendMsg(param.Topic, line.Text)
	if err != nil {
		fmt.Printf("send msg to kafka failed, err=%v\n", err)
		return
	}
	fmt.Printf("pid=%v, offset=%v\n", pid, offset)
}

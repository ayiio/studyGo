package main

import (
	"fmt"
	"logagent/logagent/config"
	"logagent/logagent/kafka"
	"logagent/logagent/taillog"
	"time"
)

// loagent入口

var (
	cfg *config.AppConfig
)

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.TailFile():
			// 2.发送到kafka
			kafka.SendMsg(cfg.KafkaConfig.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 0.load config.ini
	cfg = config.InitConfig()
	if cfg == nil {
		fmt.Println("init file load failed")
		return
	}
	fmt.Println("# load init file success.")
	// 1.初始化kafka连接
	err := kafka.InitKafka([]string{cfg.KafkaConfig.Address})
	if err != nil {
		fmt.Printf("init Kafka failed, err=%v\n", err)
		return
	}
	fmt.Println("# init kafka success.")
	// 2.打开日志文件准备收集日志
	err = taillog.InitTail(cfg.TailConfig.FilePath)
	if err != nil {
		fmt.Printf("init Tail failed, err=%v\n", err)
	}
	fmt.Println("# init tail success.")
	// 处理日志
	run()
}

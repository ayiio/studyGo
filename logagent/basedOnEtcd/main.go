package main

import (
	"fmt"
	config "logagent/conf"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"time"
)

// loagent入口

var (
	cfg *config.AppConfig
)

func main() {
	// 0.load conf.ini
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

	// 2.初始化etcd
	err = etcd.InitEtcd([]string{cfg.EtcdConfig.Address}, time.Duration(cfg.EtcdConfig.TimeOut)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err=%v\n", err)
		return
	}
	fmt.Println("# init etcd success.")

	// 2-1.从etcd中拉取日志收集项信息
	confs, err := etcd.GetConfByKey(cfg.EtcdConfig.LogKey)
	if err != nil {
		fmt.Printf("get conf from etcd failed, err=%v\n", err)
		return
	}
	fmt.Println("# get conf from etcd success.")

	// 2-2.分配哨兵检测收集项信息的变更，以便随时通知app实现热加载

	// 3.收集日志，发往kafka
	// 3-1.遍历log配置项，创建tailobj
	for _, logconf := range confs {
		err = taillog.InitTail(logconf.Path)
		if err != nil {
			fmt.Printf("init tail failed, err=%v\n", err)
			return
		}

	}
	// 3-2.发往Kafka
}

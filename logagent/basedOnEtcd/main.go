package main

import (
	"fmt"
	config "logagent/conf"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"logagent/utils"
	"sync"
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
	err := kafka.InitKafka([]string{cfg.KafkaConfig.Address}, cfg.KafkaConfig.MaxChanSize)
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

	// 2.从etcd中拉取日志收集项信息
	//实现根据[ip]拉取对应的日志配置
	ipstr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfkey := fmt.Sprint(cfg.EtcdConfig.LogKey, ipstr)
	logEntries, err := etcd.GetConfByKey(etcdConfkey)
	if err != nil {
		fmt.Printf("get conf from etcd failed, err=%v\n", err)
		return
	}
	fmt.Println("# get conf from etcd success.")

	// 3.tailf收集日志，发往kafka
	taillog.Init(logEntries)

	var wg sync.WaitGroup
	wg.Add(1)
	newLogChan := taillog.GetNewConfToChan()
	// 4.后台分配哨兵检测收集项信息的变更，以便随时通知app
	go etcd.Watcher(etcdConfkey, newLogChan)
	wg.Wait()

}

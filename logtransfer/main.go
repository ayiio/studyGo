package main

import (
	"fmt"
	"logtransfer/conf"
	"logtransfer/es"
	"logtransfer/kafka"
)

//log transfer
//将日志数据从kafka中取出来发往es

func main() {
	//a.加载配置文件
	conf := conf.InitConf()
	if conf == nil {
		fmt.Println("get config failed")
		return
	}
	fmt.Println("load config success.")
	//0.初始化
	err := es.InitES(conf.ESConf.Address, conf.ESConf.User, conf.ESConf.Passwd, conf.ESConf.ChanSize, conf.ESConf.GoNum)
	if err != nil {
		fmt.Printf("init ES failed, err=%v\n", err)
		return
	}
	fmt.Println("init ES success.")
	err = kafka.InitKafka([]string{conf.KafkaConf.Address}, conf.KafkaConf.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer failed, err=%v\n", err)
		return
	}
	fmt.Println("init kafka success.")
	select {}
}



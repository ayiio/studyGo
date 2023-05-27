package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type KafkaConfig struct {
	Address     string `ini:"address"`     // kafka地址
	MaxChanSize int    `ini:"maxchansize"` // 消息缓冲区大小
}

// kafka的topic以及tailf文件地址将从etcd中拉取
type EtcdConfig struct {
	Address string `ini:"address"` // etcd地址
	TimeOut int    `ini:"timeout"` // etcd超时时间
	LogKey  string `ini:"logkey"`  // 存储key值
}

type AppConfig struct {
	KafkaConfig `ini:"kafka"`
	EtcdConfig  `ini:"etcd"`
}

var (
	config = new(AppConfig)
)

func InitConfig() *AppConfig {
	// iniFile, err := ini.Load("config/config.ini")
	// if err != nil {
	// 	err = fmt.Errorf("load ini file failed, err=%v", err)
	// 	return nil, err
	// }
	// c = &AppConfig{
	// 	TailParam{Path: iniFile.Section("tail").Key("filePath").String()},
	// 	KafkaParam{
	// 		Address: iniFile.Section("kafka").Key("address").String(),
	// 		Topic:   iniFile.Section("kafka").Key("topic").String(),
	// 	},
	// }
	// 结构体方式映射
	err := ini.MapTo(config, "conf/conf.ini")
	if err != nil {
		fmt.Printf("load ini file failed, err=%v\n", err)
		return nil
	}
	return config
}


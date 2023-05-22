package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type TailConfig struct {
	FilePath string `ini:"filePath"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type AppConfig struct {
	KafkaConfig `ini:"kafka"`
	TailConfig  `ini:"tail"`
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
	err := ini.MapTo(config, "config/config.ini")
	if err != nil {
		fmt.Printf("load ini file failed, err=%v\n", err)
		return nil
	}
	return config
}


package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type TailParam struct {
	Path string
}

type KafkaParam struct {
	Address string
	Topic   string
}

type CofigParam struct {
	TailParam
	KafkaParam
}

func InitParam() (c *CofigParam, err error) {
	iniFile, err := ini.Load("config/config.ini")
	if err != nil {
		err = fmt.Errorf("load ini file failed, err=%v", err)
		return nil, err
	}
	c = &CofigParam{
		TailParam{Path: iniFile.Section("tail").Key("path").String()},
		KafkaParam{
			Address: iniFile.Section("kafka").Key("address").String(),
			Topic:   iniFile.Section("kafka").Key("topic").String(),
		},
	}
	return
}

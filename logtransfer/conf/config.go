package conf

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESConf struct {
	Address  string `ini:"address"`
	ChanSize int    `ini:"chansize"`
	GoNum    int    `ini:"gonum"`
	User     string `ini:"user"`
	Passwd   string `ini:"passwd"`
}

type LogTransfer struct {
	KafkaConf `ini:"kafka"`
	ESConf    `ini:"es"`
}

func InitConf() *LogTransfer {
	var lt LogTransfer
	err := ini.MapTo(&lt, "./conf/config.ini")
	if err != nil {
		fmt.Printf("ini map failed, err=%v\n", err)
		return nil
	}
	return &lt
}


package kafka

// 用于向kafka写日志的模块

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer // 声明全局连接kafka的生产者client
)

func Init(address []string) (err error) {
	config := sarama.NewConfig()
	// tail包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follower的答复
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选取partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Println("producer closed, err=", err)
		return
	}
	return
}

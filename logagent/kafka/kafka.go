package kafka

// 用于向kafka写日志的模块

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer // 声明全局连接kafka的生产者client
)

func InitKafka(address []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follower的答复
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选取partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(address, config)
	return
}

func SendMsg(topic, msg string) {
	msgObj := &sarama.ProducerMessage{}
	msgObj.Topic = topic
	msgObj.Value = sarama.StringEncoder(msg)
	pid, offset, err := client.SendMessage(msgObj)
	if err != nil {
		fmt.Printf("send kafka failed, err=%v\n", err)
		return
	}
	fmt.Printf("send kafka success, pid=%v, offset=%v\n", pid, offset)
}


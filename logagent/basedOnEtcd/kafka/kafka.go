package kafka

// 用于向kafka写日志的模块

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type logData struct {
	topic string
	data  string
}

var (
	client  sarama.SyncProducer // 声明全局连接kafka的生产者client
	logChan chan *logData       //kafka全局log chan
)

func InitKafka(address []string, maxChanSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follower的答复
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选取partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(address, config)
	//初始化log通道
	logChan = make(chan *logData, maxChanSize)
	//初始化完成后就开启后台goroutine，从log通道中取数据并发送到kafka
	go sendToKafka()
	return
}

//向外部暴漏发送数据到log通道的方法，用于将日志数据发送到内部chan
func PutChan(tp, msg string) {
	logs := &logData{
		topic: tp,
		data:  msg,
	}
	select {
	case logChan <- logs: //向log通道中发送信息
	default:
		time.Sleep(time.Microsecond * 10)
	}
}

//内部方法，从内部chan中取数据发送到kafka
func sendToKafka() {
	for {
		select {
		case log := <-logChan:
			msg := log.data
			msgObj := &sarama.ProducerMessage{}
			msgObj.Topic = log.topic
			msgObj.Value = sarama.StringEncoder(msg)
			pid, offset, err := client.SendMessage(msgObj) // 发送到kafka
			if err != nil {
				fmt.Printf("send kafka failed, err=%v\n", err)
				return
			}
			fmt.Printf("send kafka success, pid=%v, offset=%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

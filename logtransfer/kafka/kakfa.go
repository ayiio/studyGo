package kafka

import (
	"fmt"
	"logtransfer/es"

	"github.com/Shopify/sarama"
)

//初始化kafka消费者，从kafka取数据并发送到es
func InitKafka(address []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("new consumer failed, err=%v\n", err)
		return
	}
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("conn topic:%v failed, err=%v", topic, err)
		return
	}
	for partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("new partition consumer failed, err=%v\n", err)
			return err
		}
		// defer pc.AsyncClose()
		//消费kafka并发送数据到es
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Println(msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				//发送es
				es.PutEsChan(msg)
			}
		}(pc)
	}
	return
}

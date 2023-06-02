package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	//创建消费者
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("new consumer failed,", err)
		return
	}
	fmt.Println("new consumer success.")
	//根据topic取到所有分区
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Println("fail to get list of partition", err)
		return
	}
	fmt.Println("consumer conn the partition success.")
	fmt.Println(partitionList)
	var wg sync.WaitGroup
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d, err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		//异步从每个分区消费
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d, key:%v, value:%v\n", msg.Partition, msg.Key, string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
}

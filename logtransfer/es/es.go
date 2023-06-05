package es

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/olivere/elastic/v7"
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	esChan chan *sarama.ConsumerMessage
)

//ES模块

//初始化ES
func InitES(address, user, passwd string, chanSize, gonum int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(
		elastic.SetURL(address),
		elastic.SetBasicAuth(user, passwd),
	)
	if err != nil {
		fmt.Printf("new elastic client failed, err=%v\n", err)
		return
	}
	esChan = make(chan *sarama.ConsumerMessage, chanSize)

	//多个goroutine处理
	for i := 0; i < gonum; i++ {
		go sendToES(esChan)
	}
	return
}

func PutEsChan(esch *sarama.ConsumerMessage) {
	esChan <- esch
}

//发送数据到ES
func sendToES(ec <-chan *sarama.ConsumerMessage) {
	for {
		select {
		case msg := <-esChan:
			ld := LogData{
				Topic: msg.Topic,
				Data:  string(msg.Value),
			}
			put1, err := client.Index().Index(msg.Topic).Id(strconv.Itoa(int(msg.Offset))).BodyJson(ld).Do(context.Background())
			if err != nil {
				fmt.Printf("send ES failed, err=%v\n", err)
			}
			fmt.Printf("Indexed user:%s to index:%v.\n", put1.Id, put1.Index)
		default:
			time.Sleep(time.Second)
		}
	}

}


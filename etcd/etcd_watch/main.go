package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// watch获取更改的通知
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("etcd connect failed, err=%v\n", err)
		return
	}
	fmt.Println("etcd connect success")
	defer cli.Close()

	// 设定一个哨兵， 一直监视kfc 这个值发生变化(新增，修改，删除)
	ch := cli.Watch(context.Background(), "kfc") // <-chan watchResponse
	for wresp := range ch {
		for _, ev := range wresp.Events {
			fmt.Printf("type:%v, key:%v, val:%v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}
}

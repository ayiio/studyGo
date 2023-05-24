package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

type ConfEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化etcd
func InitEtcd(address []string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("etcd connect failed, err=%v\n", err)
		return
	}
	return
}

// 从etcd中根据key获取配置信息
func GetConfByKey(key string) (confs []*ConfEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get %v from etcd failed, err=%v\n", key, err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &confs)
		if err != nil {
			fmt.Printf("unmarshal conf from etcd.value failed, err=%v\n", err)
			return
		}
	}
	return
}

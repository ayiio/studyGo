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

type LogEntry struct {
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
func GetConfByKey(key string) (confs []*LogEntry, err error) {
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

//监视etcd中对应key的变化，并通知有使用到配置项的地方--tailMgr
func Watcher(key string, ch chan<- []*LogEntry) {
	watchresp := cli.Watch(context.Background(), key)
	for wr := range watchresp {
		for _, wrv := range wr.Events {
			//通知外部taillog/tailMgr
			//通过类型判断具体的值
			var newConf []*LogEntry
			//删除操作不需要使用unmarshal
			if wrv.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(wrv.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal new conf failed, err=%v\n", err)
					continue
				}
			}
			fmt.Println("conf has been updated")
			ch <- newConf
		}
	}

}

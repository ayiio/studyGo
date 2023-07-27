package main

// SET ETCDCTL_API=3

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	//设置续期时间
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	//将k-v设置到etcd
	_, err = cli.Put(context.TODO(), "root", "admin", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	//若想一直有效，需要一直续期
	ch, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Fatal(err)
	}

	for {
		c := <-ch
		fmt.Println("c: ", c)
	}
}

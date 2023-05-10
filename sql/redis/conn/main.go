package main

/*
GO111MODULE=on模式下：
go mod init github.com/my/repo
go get github.com/redis/go-redis/v9

使用go run验证
*/

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisdb *redis.Client
var ctx = context.Background()

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping(ctx).Result()
	return
}

func main() {

	fmt.Println("Go Redis Connection Test")
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err=%v\n", err)
		return
	}
	fmt.Println("连接redis成功")

	//Set / Get
	err = redisdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisdb.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)
}

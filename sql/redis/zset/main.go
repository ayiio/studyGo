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
	// _ "gopkg.in/redis.v4"
)

var redisdb *redis.Client
var ctx = context.Background()

func initRedis() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("redis ping failed, err=%v\n", err)
		return
	}
	fmt.Println("连接redis成功")
}

// 存储 Score和member： ZSet
// zrange key 0 n : 0到n从小到大排
// zrevrange key 0 n : 0到n从大到小排
// zscore key member : 获取Member分数
func zsetDemo() {
	// ZSet
	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 90, Member: "zz"},
		redis.Z{Score: 91, Member: "aa"},
		redis.Z{Score: 92, Member: "ff"},
		redis.Z{Score: 93, Member: "dd"},
		redis.Z{Score: 95, Member: "tt"},
	}
	n, _ := redisdb.ZAdd(ctx, key, items...).Result()
	fmt.Printf("保存了%d个元素\n", n)
	// 增加分数
	newScore, err := redisdb.ZIncrBy(ctx, key, 10, "zz").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err=%v\n", err)
		return
	}
	fmt.Printf("zz's new score is %f now\n", newScore)
	// 取分数最高的3个
	ret, err := redisdb.ZRevRangeWithScores(ctx, key, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err=%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println("分数前3个：", z.Member, z.Score)
	}
	// 取95到100分的元素
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(ctx, key, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscorewithscores failed, err=%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println("95-100: ", z.Member, z.Score)
	}
}

func main() {
	initRedis()
	zsetDemo()
}

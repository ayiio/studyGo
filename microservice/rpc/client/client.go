package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//请求参数
type RequestParam struct {
	A, B int
}


func main() {
	//连接远程rpc服务
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	//调用方法
	resC := 0  //用于接收返回值
	//计算周长
	err = client.Call("Area.MathZ", RequestParam{10, 20}, &resC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("周长: ", resC)
	//计算面积
	err = client.Call("Area.MathM", RequestParam{10, 20}, &resC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("面积： ", resC)
}

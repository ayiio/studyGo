package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "gRPC/proto"
)

//1.连接服务端
//2.实例化gRPC客户端
//3.调用


func main() {
	//1.连接
	conn, e := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if e != nil {
		fmt.Printf("连接异常:%s\n", e)
	}
	defer conn.Close()
	//2.实例化gRPC客户端
	client := pb.NewUserInfoServiceClient(conn)
	//3.组装一个请求参数
	req := new(pb.UserRequest)
	req.Name = "zs"
	//4.调用接口
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常：%s\n", err)
	}
	fmt.Printf("响应结果: %v\n", response)
}

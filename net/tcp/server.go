package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp server
func process(conn net.Conn) {
	var tmp [128]byte
	var msg string
	reader := bufio.NewReader(os.Stdin)
	for {
		//接受信息
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Printf("read from conn failed, err=%v\n", err)
			return
		}
		fmt.Println("收到信息", string(tmp[:n]))

		//回复信息
		fmt.Print("please reply:")
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}

func main() {
	//1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("start tcp server on localhost:20000 failed, err=%v\n", err)
		return
	}
	//2.等待其他链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed, err=%v\n", err)
			return
		}
		//3.与客户端通信
		go process(conn)
	}
}

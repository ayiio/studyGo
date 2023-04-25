package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	//1.与server端建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial localhost:20000 failed, err=%v\n", err)
		return
	}
	//2.发送数据
	var tmp [128]byte
	var msg string
	reader := bufio.NewReader(os.Stdin)
	for {
		//发送信息
		fmt.Print("please say:")
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))

		//接受信息
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Printf("read from conn failed, err=%v\n", err)
			return
		}
		fmt.Println("收到信息", string(tmp[:n]))
	}
	conn.Close()
}

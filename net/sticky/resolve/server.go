package main

import (
	"bufio"
	"fmt"
	proto "resolve/proto"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("decode failed, err=%v\n", err)
		}
		fmt.Println("收到的消息：", msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("listen failed, err=%v\n", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed, err=%v\n", err)
			return
		}
		go process(conn)
	}
}

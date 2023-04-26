package main

import (
	"fmt"
	proto "resolve/proto"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("dial failed, err=%v\n", err)
		return
	}
	msg := "hello, hello net and tcp"
	btmsg, err := proto.Encode(msg)
	if err != nil {
		fmt.Printf("encoding message failed, err=%v", err)
		return
	}
	for i := 0; i < 20; i++ {
		conn.Write(btmsg)
	}
}

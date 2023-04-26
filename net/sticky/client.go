package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("dial failed, err=%v\n", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "hello, hello net and tcp"
		conn.Write([]byte(msg))
	}

}

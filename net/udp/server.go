package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	//net.Listen("udp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("listen UDP failed, err=%v\n", err)
		return
	}
	defer listen.Close()

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read upd failed, err=%v\n", err)
			continue
		}
		fmt.Printf("receive from addr:%v, msg:%v\n", addr, string(data[:n]))
		reply := strings.ToUpper(string(data[:n]))
		listen.WriteToUDP([]byte(reply), addr)
	}
}

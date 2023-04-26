package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("dial udp failed, err=%v\n", err)
		return
	}
	defer socket.Close()
	var reader = bufio.NewReader(os.Stdin)
	var reply [1024]byte
	for {
		fmt.Print("please enter msg:")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		//收数据
		n, addr, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Printf("read from addr:%v failed, err=%v\n", addr, err)
			return
		}
		fmt.Printf("receive from addr:%v, reply:%v\n", addr, string(reply[:n]))
	}
}

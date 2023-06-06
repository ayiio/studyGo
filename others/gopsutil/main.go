package main

import (
	"fmt"
	"time"

	builtin_net "net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

//CPU info
func getCPUInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err=%v\n", err)
		return
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	//CPU使用率
	for i := 0; i < 5; i++ {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

//CPU 负载
func getCPULoad() {
	info, _ := load.Avg()
	fmt.Printf("cpu load:%v\n", info)
}

//mem info
func getMemoryInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}

//host info
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

//desk info
func getDeskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

//net info
func getNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

//get IP, method 1
func getLocalIP() {
	addrs, err := builtin_net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*builtin_net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		fmt.Printf("ip address:%v\n", ipAddr.IP.String())
	}
}

//get IP, method 2
func getLocalIP2() {
	conn, err := builtin_net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Printf("net dial failed, err=%v\n", err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*builtin_net.UDPAddr)
	fmt.Printf("local addr:%v, ip:%v\n", localAddr.String(), localAddr.IP.String())
}

func main() {
	//get cpu info
	getCPUInfo()
	//get cpu load
	getCPULoad()
	//get mem info
	getMemoryInfo()
	//host info
	getHostInfo()
	//desk info
	getDeskInfo()
	//net info
	getNetInfo()
	//get local IP method1
	getLocalIP()
	//get local IP method2
	getLocalIP2()
}

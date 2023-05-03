package main

// flag支持在命令行控制开启CUP或Mem的性能分析
// 编译成可执行文件后在执行时加上 -cpu=true 命令行参数可以生成cpu.pprof文件

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 存在问题的代码
func loginCode() {
	var c chan int // nil，一直阻塞
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			// continue  // CPU将被for-select一直占用
			time.Sleep(time.Millisecond * 500) // sleep 0.5s，让出cpu
		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		f1, err := os.Create("./cpu.pporf")
		if err != nil {
			fmt.Printf("create cpu pporf failed, err=%v\n", err)
			return
		}
		pprof.StartCPUProfile(f1)
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()
	}

	for i := 0; i < 4; i++ {
		go loginCode()
	}

	time.Sleep(10 * time.Second)

	if isMemPprof {
		f2, err := os.Create("./mem.pporf")
		if err != nil {
			fmt.Printf("create mem pprof failed, err=%v\n", err)
			return
		}
		pprof.WriteHeapProfile(f2)
		f2.Close()
	}
}

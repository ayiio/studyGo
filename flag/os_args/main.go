package main

import (
	"fmt"
	"os"
)

// 获取命令行参数
func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%T\n", os.Args)
	fmt.Printf("第一个参数: %v, 第二个参数: %v\n", os.Args[0], os.Args[1])
}

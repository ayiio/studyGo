package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
退出goroutine
方式1：全局变量
方式2：全局channel
方式3：context
*/

var wg sync.WaitGroup
var exit1 bool
var exit2 = make(chan struct{}, 1)

// 使用全局标志
func f1() {
	defer wg.Done()
	for {
		fmt.Println("test")
		time.Sleep(time.Millisecond * 500)
		if exit1 {
			break
		}
	}
}

// 使用全局channel
func f2() {
	defer wg.Done()
exit2Loop:
	for {
		fmt.Println("test")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exit2:
			break exit2Loop
		default:
		}
	}
}

// 使用context
func f3(ctx context.Context) {
	defer wg.Done()
	go f3_1(ctx)
exit3loop:
	for {
		fmt.Println("test")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break exit3loop
		default:
		}
	}
}

func f3_1(ctx context.Context) {
	defer wg.Done()
exitf3_1loop:
	for {
		fmt.Println("test")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break exitf3_1loop
		default:
		}
	}
}

func main() {
	wg.Add(1)
	fmt.Println("method 1...")
	go f1()
	time.Sleep(time.Second * 5)
	exit1 = true
	wg.Wait()

	wg.Add(1)
	fmt.Println("method 2...")
	go f2()
	time.Sleep(time.Second * 5)
	exit2 <- struct{}{}
	wg.Done()

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("method 3...")
	go f3(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()

}

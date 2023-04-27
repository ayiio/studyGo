package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作

var x int64
var y int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	lock.Lock()
	x++
	lock.Unlock()
	wg.Done()
}

func add2() {
	atomic.AddInt64(&y, 1)
	wg.Done()
}

func compare_swap() {
	var a int64 = 10
	// 旧的值是否是20，是则改为5，否则为原值
	ok := atomic.CompareAndSwapInt64(&a, 10, 5)
	fmt.Println(ok, a)
}

func main() {
	wg.Add(1000)
	// Metux方式
	for i := 0; i < 1000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)

	wg.Add(1000)
	// atomic方式
	for j := 0; j < 1000; j++ {
		go add2()
	}
	wg.Wait()
	fmt.Println(y)

	// 比较和交换
	compare_swap()
}

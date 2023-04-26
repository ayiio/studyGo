package main

import (
	"fmt"
	"sync"
)

var (
	wg   sync.WaitGroup
	once sync.Once
)

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * 2  
	}
	//如果计算复杂，或写入ch2前有其他操作，则可能出现问题
	//第一个goroutine-go f2(a,b)可能已经将ch1取完，在进行到其他操作时第二个goroutine-go f2(a,b)执行到<-ch1，此时ok将为false，第二个goroutine将继续执行close(ch2)
	//第二个goroutine关闭后退出f2，第一个goroutine执行完其他操作后将进行ch2<-，此时chan已关闭，将报panic。可不使用Once，将close拿到wg.Wait()后执行
	once.Do(func() { close(ch2) })
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for i := range b {
		fmt.Println(i)
	}
}

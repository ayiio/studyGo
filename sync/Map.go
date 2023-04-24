package main

import (
	"fmt"
	"strconv"
	"sync"
)

var mp = make(map[string]int)

// #1.使用lock
var lock sync.Mutex

// #2.使用sync.Map
var lm sync.Map

func set(k string, v int) {
	mp[k] = v
}

func get(k string) int {
	return mp[k]
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			lock.Lock()   // #1.
			set(key, n)   // 不加锁，fatal error: concurrent map writes
			lock.Unlock() // #1.
			fmt.Printf("key=%v, val=%v\n", key, get(key))
			wg.Done()
		}(i)
	}

	// #2 开箱即用的sync.Map，不需要使用make初始化
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			lm.Store(key, n)     //sync.Map -> Store
			v, _ := lm.Load(key) //sync.Map -> Load
			fmt.Printf("使用sync.Map -->  key=%v, val=%v\n", key, v)
			wg.Done()
		}(j)
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	total = 1024 * 1024
)

func main() {
	test1()
	test2()
}

// sync.Mutex
func test1() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	count := int64(0)
	t := time.Now()
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			count++
			wg.Done()
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Printf("test1 花费时间：%d, count的值为：%d \n", time.Now().Sub(t), count)
}

// sync.atomic
func test2() {
	var wg sync.WaitGroup
	count := int64(0)
	t := time.Now()
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func(i int) {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("test2 花费时间：%d, count的值为：%d \n", time.Now().Sub(t), count)
}
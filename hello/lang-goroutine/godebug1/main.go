package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg)
	}
	wg.Wait()
	// Wait to see the global run queue deplete.
	time.Sleep(3 * time.Second)
}

func work(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}
	wg.Done()
}


// go build main.go

// schedtrace参数告诉运行时打印一行调度器的摘要信息到标准err输出中，时间间隔可以指定，单位毫秒，如下所示：
// GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./example
// 如果在windows下可以运行 set GOMAXPROCS=1 && set GODEBUG=schedtrace=1000 && example
// 程序开始后每个一秒就会打印一行调度器的概要信息

// SCHED 0ms: gomaxprocs=1 idleprocs=0 threads=2 spinningthreads=0 idlethreads=0 runqueue=0 [1]
// SCHED 1009ms: gomaxprocs=1 idleprocs=0 threads=3 spinningthreads=0 idlethreads=1 runqueue=0 [9]


// 输出项	意义
// 1009ms		自从程序开始的毫秒数
// gomaxprocs=1	配置的处理器数(逻辑的processor，也就是Go模型中的P,会通过操作系统的线程绑定到一个物理处理器上)
// threads=3		运行期管理的线程数，目前三个线程
// idlethreads=1	空闲的线程数,当前一个线程空闲，两个忙
// idleprocs=0	空闲的处理器数,当前0个空闲
// runqueue=0	在全局的run队列中的goroutine数，目前所有的goroutine都被移动到本地run队列
// [9]			本地run队列中的goroutine数，目前9个goroutine在本地run队列中等待
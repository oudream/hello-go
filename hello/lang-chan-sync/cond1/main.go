package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	c := sync.NewCond(&m)
	n := 200
	running := make(chan int, n)
	awake := make(chan int, n)
	exit := false
	for i := 0; i < n; i++ {
		go func(g int) {
			m.Lock()
			for !exit {
				running <- g
				c.Wait()
				awake <- g
			}
			m.Unlock()
		}(i)
	}
	for i := 0; i < n; i++ {
		for i := 0; i < n; i++ {
			<-running // Will deadlock unless n are running.
		}
		if i == n-1 {
			m.Lock()
			exit = true
			m.Unlock()
		}
		select {
		case <-awake:
			fmt.Errorf("goroutine not asleep")
		default:
		}
		m.Lock()
		c.Broadcast()
		m.Unlock()
		seen := make([]bool, n)
		for i := 0; i < n; i++ {
			g := <-awake
			if seen[g] {
				fmt.Errorf("goroutine woke up twice")
			}
			seen[g] = true
		}
	}
	select {
	case <-running:
		fmt.Errorf("goroutine did not exit")
	default:
	}
	c.Broadcast()
}

package main

import (
	"fmt"
	"time"
)

func HelloTicker1() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

func main() {
	fmt.Println("Language begin:")
	HelloTicker1()
	fmt.Println("Language end.")
}

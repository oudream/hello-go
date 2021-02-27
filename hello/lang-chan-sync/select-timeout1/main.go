package main

import "time"

import "fmt"

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// select设置为1秒超时
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}
package main

import (
	"fmt"
	"time"
)

func chan21() {
	var c chan int
	if c == nil {
		fmt.Println("channel a is nil, going to define it")
		c = make(chan int)
		fmt.Printf("Type of a is %T", c)
	}
}

func hello22() {
	fmt.Println("Hello world goroutine")
}

func chan22() {
	go hello22()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}

func hello23(i int) {
	fmt.Println("go %d", i)
}

func chan23() {
	for i:=0; i<10; i++ {
		go hello23(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}

func main() {
	chan22()
}
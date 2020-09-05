package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Begin")

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	fmt.Println("Step1")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	fmt.Println("End1")

	time.Sleep(2 * time.Second)

	fmt.Println("End2")
}
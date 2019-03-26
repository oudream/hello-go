package main

import (
	"fmt"
	"hello-go/hello/language1/panic1"
)

func main() {
	fmt.Println("Language begin:")
	//loop1.HelloFor1()
	//chan1.HelloChanTimeOut1()
	panic1.HelloPanic1()
	fmt.Println("Language end.")
}

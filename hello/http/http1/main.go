package main

import (
	"fmt"
	"hello-go/hello/http/http1/multi-server1"
)

func main() {
	fmt.Println("Language begin:")
	//server1.HelloServer1()
	multi_server1.HelloMultiServer1()
	fmt.Println("Language end.")
}

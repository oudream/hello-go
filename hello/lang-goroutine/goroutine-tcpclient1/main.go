package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
)

var host = flag.String("host", "localhost", "The hostname or IP to connect to; defaults to \"localhost\".")
var port = flag.Int("port", 5555, "The port to connect to; defaults to 5555.")
var address string

func helloTcpClient1(i int) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
		os.Exit(1)
	}
	for {
		scanner := bufio.NewScanner(conn)
		for {
			ok := scanner.Scan()
			text := scanner.Text()
			if !ok {
				fmt.Println("Reached EOF on server connection.")
				break
			}
			fmt.Printf("\b\b** connection%d: %s\n> ", i, text)
		}
	}
}

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumCgoCall())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.GOMAXPROCS(100))
	fmt.Println(runtime.GOMAXPROCS(0))

	flag.Parse()

	address = *host + ":" + strconv.Itoa(*port)
	fmt.Printf("Connecting to %s...\n", address)

	for i := 0; i < 24; i++ {
		go helloTcpClient1(i)
	}
	helloTcpClient1(999)
}

package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"time"
)

func see(s string) {
	//runtime.LockOSThread()

	a := time.Now()
	for i := 1; i < 20000; i++ {
		//runtime.Gosched()
		var sum int
		sum = 1
		sum = sum * i
		fmt.Println("sum:", sum)
		fmt.Println("s:", s)
		getid("see " + s)
		//fmt.Println("cpus:", runtime.StackRecord{})
		//fmt.Println("cpus:", runtime.LockOSThread)
	}
	fmt.Println(time.Since(a))
}
func getid(s string) {
	pid := unix.Getpid()
	fmt.Println(s+" Getpid", pid)
	fmt.Println(s+" Getppid", unix.Getppid())

	pgid, _ := unix.Getpgid(pid)
	fmt.Println(s+" Getpgid", pgid)

	fmt.Println(s+" Gettid", unix.Gettid())

	sid, _ := unix.Getsid(pid)
	fmt.Println(s+" Getsid", sid)

	fmt.Println(s+" Getegid", unix.Getegid())
	fmt.Println(s+" Geteuid", unix.Geteuid())
	fmt.Println(s+" Getgid", unix.Getgid())
	fmt.Println(s+" Getuid", unix.Getuid())
}

func main() {
	go see("hello")
	go see("world")
	getid("main")
	time.Sleep(20 * time.Second)
}
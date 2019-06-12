package main

// See main.go first.
// Sometimes you still need a hand-written wrapper.

import (
	"hello-go/hello/cgo"
)

func main() {
	cgo.CallF5WithF()
	cgo.CallHello1(11)
}

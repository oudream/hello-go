package main

// See main.go first.
// Sometimes you still need a hand-written wrapper.

import (
	"hello-go/hello/cgo2/wrapper"
)

func main() {
	wrapper.CallF5WithF()
	wrapper.CallHello1(11)
}

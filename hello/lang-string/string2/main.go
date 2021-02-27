package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

type BinlogState struct {
	name string
	position int
	time int64
}

func main() {
	helloScantf3()
	//fmt.Println(fmt.Sprintf("%s:%d", "0.0.0.0", 8002))
}

func helloScantf3() {
	// Declaring two variables
	var name string
	var alphabet_count int
	var count int64

	// Calling the Sscanf() function which
	// returns the number of elements
	// successfully parsed and error if
	// it persists
	n, err := fmt.Sscanf("mysql-bin.000007 766840 1610945977225",
		"%s %d %d", &name, &alphabet_count, &count)

	// Below statements get
	// executed if there is any error
	if err != nil {
		panic(err)
	}

	// Printing the number of
	// elements and each elements also
	fmt.Printf("%d: %s, %d\n", n, name, alphabet_count)
}

func helloScantf2() {
	// Declaring two variables
	var name string
	var alphabet_count int

	// Calling the Sscanf() function which
	// returns the number of elements
	// successfully parsed and error if
	// it persists
	n, err := fmt.Sscanf("GFG is having 3 alphabets.",
		"%s is having %d alphabets.", &name, &alphabet_count)

	// Below statements get
	// executed if there is any error
	if err != nil {
		panic(err)
	}

	// Printing the number of
	// elements and each elements also
	fmt.Printf("%d: %s, %d\n", n, name, alphabet_count)
}

func helloScantf1() {
	var state BinlogState
	lineText := "mysql-bin.000007,766840,1610945977225"
	if _, err := fmt.Sscanf(lineText, "%s,%d,%d", &state.name, &state.position, &state.time); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(state)
	}
}

// 强转换: https://segmentfault.com/a/1190000037679588
// 通过unsafe和reflect包，可以实现另外一种转换方式，我们将之称为强转换（也常常被人称作黑魔法）。
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func helloByte1() {
	var b1 strings.Builder
	b1.WriteString("ABC")
	for i := 0; i < 1024*1024*10; i++ {
		b1.WriteByte(30 + byte(i%10))
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	fmt.Println(b1.Len())
}

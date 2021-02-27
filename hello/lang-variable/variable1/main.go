package main

import (
	"fmt"
)

// *** prove variable init sort *** begin: ***
func getName() (r string) {
	println("i am getName[main]")
	r = "a"
	return
}

var fiIndex = 1

func getCount() (r int) {
	fiIndex++
	r = fiIndex
	return r
}

var fsName = getName()

var (
	a = getCount()
	//b int = getCount()
	c int
	d int
)
// *** prove variable init sort *** end: ***



func testVariable1(v int) (r int) {
	// *** prove const variable iota scope in page
	fmt.Println(fsName)
	fmt.Println(a, " - ", "\n")
	fmt.Println(c, " - ", d, "\n")

	// *** prove const variable iota scope in page
	//fmt.Println("main Variable1 :", pgida, " - ", pgidb)
	return v
}

func RegCallBack(fn interface{}) {

}

func main() {
	fmt.Println("Language begin:")

	const(
		MaxUint32 = ^uint32(0)
		MinUint32 = uint32(0)
		MaxInt32 = int32(MaxUint32 >> 1)
		MinInt32 = -MaxInt32 - 1
		MaxUint64 = ^uint64(0)
		MinUint64 = uint64(0)
		MaxInt64 = int(MaxUint64 >> 1)
		MinInt64 = -MaxInt64 - 1
	)
	fmt.Println(MaxInt32)
	fmt.Println(MinInt32)
	fmt.Println(MaxInt64)
	fmt.Println(MinInt64)

	fmt.Println("Language end.")
}

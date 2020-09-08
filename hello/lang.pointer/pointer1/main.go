package main

import "fmt"

var str *string
var i *int
var arr [3]int

func helloPointer1() {
	str1 := "abcdef"
	*str = str1
}

func helloPointer2() {
	str1 := "b"
	str = &str1
	fmt.Println(&str1)
}

func helloPointer3() {
	fmt.Println(*str)
	fmt.Println(str)
}

func main() {
	//arr[2] = 11
	//j := &arr[2]
	//fmt.Println(j)
	//fmt.Printf("%T\n",j)
	//fmt.Printf("%T\n",*j)
	//fmt.Println(*j)
	//*j = 12
	//fmt.Println(j)
	//fmt.Println(*j)
	//fmt.Println(str)
	//fmt.Println(&str)
	//fmt.Println(i)
	//fmt.Println(&i)
	helloPointer2()
	helloPointer1()
	helloPointer3()
}

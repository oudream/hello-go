package main

import "fmt"

var str *string
var i *int
var arr [3]int

func helloPointer0() {
	arr[2] = 11
	j := &arr[2]
	fmt.Println(j)
	fmt.Printf("%T\n",j)
	fmt.Printf("%T\n",*j)
	fmt.Println(*j)
	*j = 12
	fmt.Println(j)
	fmt.Println(*j)
	fmt.Println(str)
	fmt.Println(&str)
	fmt.Println(i)
	fmt.Println(&i)
}

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

func helloPointer4() {
	var p1 *string
	str1 := "xxxxxxxxxx"
	var str2 string
	p1 = &str2
	*p1 = str1
	fmt.Println(*p1)
}

func main() {
	//helloPointer2()
	//helloPointer1()
	//helloPointer3()
	helloPointer4()
}

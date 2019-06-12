package array1

import (
	"fmt"
)


func getName() (r string) {
	println("i am getName[array1]")
	r = "a"
	return
}

var fsName = getName()

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func HelloArray11() {
	arr1 := make([]int, 3, 10)
	printSlice(arr1)
	for i := 0; i < len(arr1); i++  {
		arr1[i] = i * 10
	}
	printSlice(arr1)
	for i := len(arr1); i < cap(arr1); i++  {
		arr1 = append(arr1, i * 11)
	}
	printSlice(arr1)
}

func HelloArray12() {
	arr1 := make([]int, 3, 10)
	printSlice(arr1)
	for i := 0; i < len(arr1); i++  {
		arr1[i] = i * 10
	}
	printSlice(arr1)
	for i := len(arr1); i < cap(arr1); i++  {
		arr1 = append(arr1, i * 11)
	}
	printSlice(arr1)
}

func testVariable1(v int) (r int) {
	HelloArray11()
	r = 0
	return
}
//var iTemp = utils.RegCallBack(testVariable1)


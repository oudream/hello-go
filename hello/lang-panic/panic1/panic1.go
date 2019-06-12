package panic1

import (
	"fmt"
)

func HelloPanic11() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()
	//g()
}

func f() {
	fmt.Println("a")
	fmt.Println("b")
	panic(55)
	fmt.Println("f")
}

func g() {
	fmt.Println("a")
	fmt.Println("b")
	fmt.Println("f")
}

func init() {
	fmt.Println("i am panic1, init 2!")
}

func init() {
	fmt.Println("i am panic1, init 1!")
}

func testVariable1(v int) (r int) {
	HelloPanic11()
	r = 0
	return
}
//var iTemp = utils.RegCallBack(testVariable1)


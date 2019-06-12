package sub1

import (
	"fmt"
	"hello-go/hello/lang-utils/utils1"
)

const (
	pgida = iota
	pgidb = iota
	pgidc = iota
)

var Var1String1 = "abc1"
var Var1String2 = "abc2"
var Var1Int1 = 1
var Var1Int2 = 2


func HelloVariable1() {
	fmt.Println("Hello Variable1 :", pgida, " - ", pgidb)
}

func changeRunes(s []rune) {
	s = s[1:]
}

func changeRunes3(s []rune) {
	s[0] = 3
	s[1] = 5
}

func changeRunes2(s *[]rune) {
	x := []rune{7,9}
	*s = x
}

func changeRunes4(s * int) {
	*s = 11
}

func HelloVariable2 () {
	x := []rune{1,2}
	changeRunes(x)
	fmt.Println(x)
	changeRunes2(&x)
	fmt.Println(x)
	changeRunes3(x)
	fmt.Println(x)
	var s = [5]int{1, 2, 3, 4, 5}
	i := int(x[1])
	changeRunes4(&i)
	changeRunes4(&s[1])
	fmt.Println(s)
}

func testVariable1(v int) (r int) {
	HelloVariable1()
	HelloVariable2()
	return v
}
var iTemp = utils1.RegCallBack(testVariable1)

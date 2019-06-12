package variable1

import (
	"fmt"
	"hello-go/hello/lang-utils/utils1"
	"hello-go/hello/lang-variable/variable1/sub1"
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


// *** prove const variable iota scope in page
const (
	pgida = iota
	pgidb = iota
	pgidc = iota
)

func testVariable1(v int) (r int) {
	// *** prove const variable iota scope in page
	fmt.Println(fsName)
	fmt.Println(a, " - ", "\n")
	fmt.Println(c, " - ", d, "\n")

	// *** prove const variable iota scope in page
	fmt.Println("main Variable1 :", pgida, " - ", pgidb)

	// *** prove global variable can modify by other page
	fmt.Println(sub1.Var1Int1)
	sub1.Var1Int2++
	fmt.Println(sub1.Var1Int2)
	return v
}
var iTemp1 = utils1.RegCallBack(testVariable1)

func main() {
	fmt.Println("Language begin:")

	utils1.DoCallBack(0)

	fmt.Println("Language end.")
}

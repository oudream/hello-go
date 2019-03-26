package wrapper

/*
#include "wrapper.h"
*/
import "C"
import "fmt"

func CallF5WithF() {
	C.call_f5_with_F()
}

func CallHello1(i int) (r int) {
	cR := C.c_hello1(C.int(i))
	goR := int(cR)
	fmt.Println("CallHello1.Result=",goR)
	return int(cR)
}

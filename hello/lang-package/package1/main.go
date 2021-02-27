package main

import (
	"fmt"
)

import userinfo "./userinfo"
import common "../common"

func init() {
	WhatIsThe = 0
}

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func main() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}

	var u Utils
	fmt.Println(u.ShowName())
	fmt.Println(u.ShowName())
	fmt.Println(u.ShowName())
	fmt.Println(u.ShowName())
	fmt.Println(UtilsName)

	userinfo.UserLogin() // 包名.函数名() 函数名首字母必须大写

	common.ShowDemo()

	fmt.Println(Redis)
}
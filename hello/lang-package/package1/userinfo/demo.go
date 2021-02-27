package userinfo // 包名与所在目录名相同 (同一目录下的所有.go文件的包名要相同)

import "fmt"

// 包名.函数名()的方式调用时，函数名首字母必须大写
func UserLogin() {
	fmt.Println("用户登录成功！")
}

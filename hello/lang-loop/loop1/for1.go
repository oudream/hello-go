package loop1

import (
	"fmt"
	"unicode/utf8"
)

/*
Go 语言循环语句Go 语言循环语句

for循环是一个循环控制结构，可以执行指定次数的循环。

语法 ：3种形式
Go语言的For循环有3种形式，只有其中的一种使用分号。

和 C 语言的 for 一样：

for init; condition; post { }
和 C 的 while 一样：

for condition { }
和 C 的 for(;;) 一样：

for { }
init： 一般为赋值表达式，给控制变量赋初值；
condition： 关系表达式或逻辑表达式，循环控制条件；
post： 一般为赋值表达式，给控制变量增量或减量。
for语句执行过程如下：

①先对表达式1赋初值；
②判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，
	进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

for key, value := range oldMap {
newMap[key] = value
}
 */

func HelloFor12() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
}

func HelloFor11() (n int, err error) {
	s := "猪不是笨的！"       //含有中文字符 ​
	fmt.Println(len(s)) // output: 18, len获得的是字节数，一个中文3个字节
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))
	// output: Rune count:6，所以可以使用 utf8.RuneCountInString(s)获得字符数量

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) //UTF-8,英文1个字节，中文3个字节
	}
	// output: E7 8C AA E7 8C AA E6 98 AF E4 B8 AA E7 AC A8 EF BC 81

	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	// output: (0 732A) (3 732A) (6 662F) (9 4E2A) (12 7B28) (15 FF01)

	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %c) ", i, ch)
	} // output: (0 猪) (3 不) (6 是) (9 笨) (12 的) (15 ！) 直接使用range，返回的下标不是连续的

	fmt.Println()
	fmt.Println([]byte(s)) // []byte可以获得所有的字节
	// output: [231 140 170 231 140 170 230 152 175 228 184 170 231 172 168 239 188 129]
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	// output: 猪 不 是 笨 的 ！

	fmt.Println()
	// []rune会自己进行转换，把准换好的东西放在一个数组中，再开一个rune slice出来
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	// output: (0 猪) (1 不) (2 是) (3 笨) (4 的) (5 ！)
	fmt.Println()

	return 1, nil
}

func testVariable1(v int) (r int) {
	HelloFor11()
	r = 0
	return
}
//var iTemp = utils.RegCallBack(testVariable1)

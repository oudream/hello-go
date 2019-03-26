package loop1

import (
	"fmt"
	"unicode/utf8"
)

func HelloFor1() (n int, err error) {
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

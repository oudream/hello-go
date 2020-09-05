package main

import "fmt"

func helloMap1() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

func helloMap2() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	// m2 := map[string]int{"a":1, "b":2, "c":3}

	m2 := make(map[string]int)

	for key, v := range m1 {
		// fmt.Println(key)
		// fmt.Println(v)
		m2[v] = key

	}

	fmt.Println(m2)
}

func main() {
	helloMap2()
}

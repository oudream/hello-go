package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func helloList1() {
	// 切片 slice  比较常用, 很灵活
	list1 := [6]int{1, 2, 3, 4} //6是总数, 后面是值, 如果不够会自动补0
	list6 := []string{"a", "b"}
	fmt.Println(reflect.TypeOf(list1))
	var list2 [3]int
	// 声明列表, 下面两种为初始化, 生成内存地址, 双链表 ----容器
	var list3 = list.List{}
	list4 := list.New()
	//多维数组
	list5 := [3][2]string{
		{"飞", "小"},
		{"祥", "泰"},
		{"德", "丙"},
	}
	//列表插入方法
	a1 := list3.PushFront(2)         //从左插入
	a2 := list3.PushBack(1)          //从右插入
	list3.InsertAfter("after", a2)   //在 a2之后
	list3.InsertBefore("before", a1) //在 a1之前
	fmt.Println(reflect.TypeOf(list3))
	fmt.Println(reflect.TypeOf(list4))

	//列表删除
	list3.Remove(a2)

	//列表(容器)遍历
	for x := list3.Front(); x != nil; x = x.Next() {
		if x.Value == "after" {
			fmt.Println(x.Value)
		}
		fmt.Print(x.Value, " , ")
	}
	//切片遍历
	for _, x := range list1 {
		if x == 1 {
			fmt.Println(x)
		}
	}
	//切片追加
	list6 = append(list6, "c")
	//也可以和python一样根据索引覆盖值
	list1[5] = 9
	list2[2] = 9

	fmt.Println(list1)
	fmt.Println(list2)
	printlist(list3)
	fmt.Println(list4)
	fmt.Println(list5)
	fmt.Println(list6)
}

func printlist(lists list.List) {
	for x := lists.Front(); x != nil; x = x.Next() {
		fmt.Println(x.Value)
	}
}

type StructA struct {

}

func getList1() *list.List {
	return nil
}

func main() {
	var list6 []int
	list3 := list.List{}
	list6 = nil
	fmt.Println(reflect.TypeOf(list6))
	fmt.Println(len(list6))
	for i, i2 := range list6 {
		fmt.Println(i, i2)
	}
	//helloList1()
	fmt.Println(reflect.TypeOf(list3))
	fmt.Println(list3.Len())

	list4 := getList1()
	list3 = *list4
	fmt.Println(list3.Len())
}

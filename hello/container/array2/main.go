package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func helloExist1() {
	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strArray, "Canada"))
	fmt.Println(itemExists(strArray, "Africa"))
}

func helloRemove1() {
	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	arr1 := strArray[:2]
	println(arr1)
}

type Student struct {
	Name string
	Age int
}

func helloModify1(students []Student)  {
	for i := 0; i < len(students); i++ {
		student := &(students)[i]
		student.Name = fmt.Sprintf("M1-%d", i)
		student.Age = 10 * i
	}
}

func helloModify2(students []Student)  {
	for i := 0; i < len(students); i++ {
		students[i].Name = fmt.Sprintf("M2-%d", i)
		students[i].Age = 100 * i
	}
}

func helloModify3(students *[]Student)  {
	for i := 0; i < len(*students); i++ {
		student := &(*students)[i]
		student.Name = fmt.Sprintf("name%d", i)
		student.Age = i * i
	}
}

func helloModify4a(student *Student)  {
	student.Name = fmt.Sprintf("M1-%d", 666)
	student.Age = 10 * 666
}

func helloModify4(students []Student)  {
	for i := 0; i < len(students); i++ {
		student := &students[i]
		helloModify4a(student)
	}
}

func helloModify0() {
	var students []Student
	for i := 0; i < 10; i++ {
		students = append(students, Student{"", i})
	}
	fmt.Println(students)
	helloModify1(students)
	fmt.Println(students)
	helloModify2(students)
	fmt.Println(students)
	helloModify4(students)
	fmt.Println(students)
}

func helloModifyString1(s string)  {
	s = "helloModifyString1"
}

func helloModifyString3(s *string)  {
	*s = "helloModifyString3"
}

func helloModifyString0()  {
	s := "a"
	fmt.Println(s)
	helloModifyString1(s)
	fmt.Println(s)
	helloModifyString3(&s)
	fmt.Println(s)
}

func helloModifyStruct1(s Student)  {
	s.Name = "helloModifyStruct1"
}

func helloModifyStruct3(s *Student)  {
	s.Name = "helloModifyStruct3"
}

func helloModifyStruct0()  {
	//var s Student
	s := new(Student)
	fmt.Println(s)
	helloModifyStruct1(*s)
	fmt.Println(s)
	helloModifyStruct3(s)
	fmt.Println(s)
}

type Teacher struct {
	Name string
	Level string
	Students []Student
}

func helloTeacher1(teachers []Teacher)  {
	for i := 0; i < len(teachers); i++ {
		teacher := &(teachers)[i]
		teacher.Name = fmt.Sprintf("M1-%d", i)
		teacher.Level = fmt.Sprintf("L1-%d", i)
		students := []Student{ {Name: "a", Age: 1}, {Name: "b", Age: 2} }
		teacher.Students = append(teacher.Students, students...)
	}
}

func helloTeacher2(teachers []Teacher)  {
	for i := 0; i < len(teachers); i++ {
		teachers[i].Name = fmt.Sprintf("M2-%d", i)
		teachers[i].Level = fmt.Sprintf("L2-%d", i)
		students := []Student{ {Name: "c", Age: 3}, {Name: "d", Age: 4} }
		teachers[i].Students = append(teachers[i].Students, students...)
	}
}

func helloTeacher3(teachers *[]Teacher)  {
	for i := 0; i < len(*teachers); i++ {
		teacher := &(*teachers)[i]
		teacher.Name = fmt.Sprintf("name%d", i)
		teacher.Level = fmt.Sprintf("level%d", i)
		students := []Student{ {Name: "e", Age: 7}, {Name: "f", Age: 8} }
		teacher.Students = append(teacher.Students, students...)
	}
}

func helloTeacher0() {
	teachers := make([]Teacher, 10, 100)
	//var teachers []Teacher
	//students := []Student{ {Name: "o", Age: 17}, {Name: "p", Age: 18} }
	//for i := 0; i < 10; i++ {
	//	teachers = append(teachers, Teacher{"", "", students})
	//}
	fmt.Println(teachers)
	helloTeacher1(teachers)
	fmt.Println(teachers)
	helloTeacher2(teachers)
	fmt.Println(teachers)
	helloTeacher3(&teachers)
	fmt.Println(teachers)
}

func helloCap1() {
	iis := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		iis = append(iis, i)
	}
	fmt.Println(len(iis))
	fmt.Println(cap(iis))
}

func main() {
	//helloExist1()
	helloModify0()
	//helloModifyString0()
	//helloModifyStruct0()
	//helloTeacher0()
	//helloCap1()
	//helloDeepEqual1()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func helloDeepEqual1() {
	s1 := RandStringRunes(10240)
	s2 := RandStringRunes(10240)

	fmt.Println(reflect.DeepEqual(1.1, 1.1))
	fmt.Println(reflect.DeepEqual(1.1, 1))
	fmt.Println(reflect.DeepEqual("aaa", "aaa"))
	fmt.Println(reflect.DeepEqual(time.Now(), time.Now()))

	var a1 []int
	var a2 []int
	for i := 0; i < 10000; i++ {
		a1 = append(a1, i)
		a2 = append(a2, i)
	}
	fmt.Println(reflect.DeepEqual(a1, a2))
	fmt.Println(reflect.DeepEqual(s1, s2))

	now1 := time.Now().Nanosecond() / 1e6
	for i := 0; i < 100000; i++ {
		reflect.DeepEqual([]byte(s1), []byte(s2))
	}
	now2 := time.Now().Nanosecond() / 1e6
	fmt.Println(now2 - now1)
	for i := 0; i < 1000; i++ {
		reflect.DeepEqual(a1, a2)
	}
	now3 := time.Now().Nanosecond() / 1e6
	fmt.Println(now3 - now2)
}

package main

import (
	"encoding/json"
	"time"
)

type Student2 struct {
	Name string `json:"name"`
	// 一定要将json的tag设置忽略掉不解析出来
	Brith time.Time `json:"-"`
}

// 实现它的json序列化方法
func (this Student2) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasStu Student2
	// 定义一个新的结构体
	tmpStudent := struct {
		AliasStu
		Brith string `json:"brith"`
	}{
		AliasStu: (AliasStu)(this),
		Brith:    this.Brith.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(tmpStudent)
}
func main() {
	stu2 := Student2{
		Name:  "qiangmzsx",
		Brith: time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local),
	}

	b2, err := json.Marshal(stu2)
	if err != nil {
		println(err)
	}
	println(string(b2)) //{"name":"qiangmzsx","brith":"1993-01-01 20:08:23"}

	stus := []Student2{{
		Name:  "qiangmzsx1",
		Brith: time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local),
	},{
		Name:  "qiangmzsx2",
		Brith: time.Date(1994, 2, 2, 20, 8, 22, 22, time.Local),
	}}

	b3, err := json.Marshal(stus)
	if err != nil {
		println(err)
	}

	println(string(b3)) //{"name":"qiangmzsx","brith":"1993-01-01 20:08:23"}
}

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Inner struct {
	Data time.Time `json:"date"`
}

type Outer struct {
	I   Inner `json:"i"`
	Num int   `json:"num"`
}

func helloJson1() {
	b, _ := json.Marshal(Outer{I: Inner{
		Data: time.Now()}, Num: 11})
	fmt.Println(string(b))
}

type Level struct {
	A int	 `json:"a"`
	B string `json:"b"`
}

type Student struct {
	I int
	S string
	O []Level
}

func helloJson2() {
	b, _ := json.Marshal(Student{I: 11, S: "abc", O: []Level{
		{A: 12, B: "eft"},
		{A: 13, B: "gik"},
	}})
	fmt.Println(string(b))
}

func helloJson3()  {
	const s = "{\n    \"error\": {\n        \"root_cause\": [\n            {\n                \"type\": \"index_already_exists_exception\",\n                \"reason\": \"index [goeft_properties1/0UAO5IIDSzCR1ThvJJ20cQ] already exists\",\n                \"index_uuid\": \"0UAO5IIDSzCR1ThvJJ20cQ\",\n                \"index\": \"goeft_properties1\"\n            }\n        ],\n        \"type\": \"index_already_exists_exception\",\n        \"reason\": \"index [goeft_properties1/0UAO5IIDSzCR1ThvJJ20cQ] already exists\",\n        \"index_uuid\": \"0UAO5IIDSzCR1ThvJJ20cQ\",\n        \"index\": \"goeft_properties1\"\n    },\n    \"status\": 400\n}"
	result := map[string]interface{}{}
	_ = json.Unmarshal([]byte(s), &result)
	fmt.Println(result)
}

func main() {
	helloJson3()
}

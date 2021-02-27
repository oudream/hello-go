package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type AS struct {
	Description string `json:"description"`
	Processors []ASP `json:"processors"`
}

type ASP struct {
	Attachment ASPA `json:"attachment"`
}

type ASPA struct {
	Field string `json:"field"`
	IndexedChars int64 `json:"indexed_chars"`
	IndexedWords int64 `json:"indexed_words"`
	T1 time.Time `json:"t1"`
	T2 *time.Time `json:"t2"`
}

func getTime() *time.Time {
	t := time.Now()
	return &t
}

func helloEncode1() {
	attachment := &AS{
		Description: "Process documents",
		Processors: []ASP{
			ASP{
				Attachment: ASPA{
					Field: "thedata",
					IndexedChars: -1,
					T1: time.Now(),
					T2: getTime(),
				},
			},
		},
	}
	b, _ := json.Marshal(attachment)
	fmt.Println(string(b))
}

func helloDecode1() {
	attachment := &AS{
		Description: "",
		Processors: []ASP{ASP{Attachment: ASPA{}}},
	}
	s := "{\"description\":\"Process documents\",\"processors\":[{\"attachment\":{\"field\":\"thedata1\",\"indexed_chars\":-9}},{\"attachment\":{\"field\":\"thedata2\",\"col\":\"coldata2\",\"indexed_chars\":-11}}]}"
	_ = json.Unmarshal([]byte(s), attachment)
	fmt.Println(*attachment)
}

func helloDecode2() {
	attachment := &AS{
		Description: "",
		Processors: []ASP{ASP{Attachment: ASPA{}}},
	}
	//s := "{\"description\":\"Process documents\",\"processors\":[" +
	//	"{\"attachment\":{\"field\":\"thedata\",\"indexed_chars\":-11,\"indexed_words\":2,\"t1\":\"2020-11-07T10:48:37.6667399+08:00\",\"t2\":\"2020-11-07T10:48:37.6667399+08:00\"}}," +
	//	"{\"attachment\":{\"field\":\"thedata\",\"indexed_chars\":-12,\"indexed_words\":3,\"t1\":\"2020-11-08 11:11:11\"}}" +
	//	"]}"
	s1 := "\"_shards\": {\n        \"total\": 1,\n        \"successful\": 1,\n        \"skipped\": 0,\n        \"failed\": 0\n    }"
	_ = json.Unmarshal([]byte(s1), attachment)
	fmt.Println(*attachment)
}

type RpResult struct {
	Code string `json:"code"`
	Datas interface{} `json:"datas"`
}

type RpError struct {
	Error string `json:"error"`
}

func encodeRespError(c, e string) string {
	attachment := &RpResult{
		Code: c,
		Datas: RpError{
			Error: e,
		},
	}
	b, _ := json.Marshal(attachment)
	return string(b)
}

func encodeRespSuccess(d map[string]interface{}) string {
	attachment := &RpResult{
		Code: "200",
		Datas: d,
	}
	b, _ := json.Marshal(attachment)
	return string(b)
}

func helloEncode2() {
	fmt.Println(encodeRespError("400", "xxx"))
}

func helloEncode3() {
	attachment := AS{
		Description: "Process documents",
		Processors: []ASP{
			ASP{
				Attachment: ASPA{
					Field: "thedata",
					IndexedChars: -1,
					T1: time.Now(),
					T2: getTime(),
				},
			},
		},
	}
	d := map[string]interface{}{
		"query": attachment,
	}
	fmt.Println(encodeRespSuccess(d))
}

type Configuration struct {
	Val string 		`json:"val"`
	Proxy struct {
		Address string `json:"addr"`
		Port    int `json:"port"`
	} `json:"pro"`
}

func helloDecode3()  {
	attachment := Configuration{}
	s := "xxxxxxxxx, yyyyyyyy"
	//s := "{\"val\": \"val1\",\"pro\": {\"addr\": \"addr1\",\"port\": 1212}}"
	_ = json.Unmarshal([]byte(s), &attachment)
	fmt.Println(attachment)
}

func main() {
	//helloEncode1()
	//helloDecode1()
	//helloDecode2()
	//helloEncode2()
	//helloEncode3()
	helloDecode3()
}

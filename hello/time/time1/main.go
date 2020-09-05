package main

import (
	"fmt"
	"time"
)

func main() {
	//返回现在时间 Time 时间类型
	timeNow := time.Now() //2012-10-31 15:50:13.793654 +0000 UTC

	//Time 时间转化为string
	timeString := timeNow.Format("2006-01-02 15:04:05:001") //2015-06-15 08:52:32
	fmt.Println(timeString)

	//获取时间戳
	timestamp := time.Now().Unix() //1504079553

	//时间戳转Time 再转 string
	timeNow = time.Unix(timestamp, 0) //2017-08-30 16:19:19 +0800 CST
	timeString = timeNow.Format("2006-01-02 15:04:05") //2015-06-15 08:52:32
	fmt.Println(timeString)

	//string 转 时间戳
	stringTime := "2017-08-30 16:40:41"
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	if err == nil {
		unixTime := theTime.Unix() //1504082441
		fmt.Println(unixTime)
	}
}
package main

import (
	"fmt"
	"time"
)

// layout 数值必须为： "2006-01-02 15:04:05"
// layout 数值必须为： "2006-01-02 15:04:05"
// layout : RFC3339
// layout : RFC3339Nano

func tmTimestamp2Time(tm int64) time.Time {
	return time.Unix(tm/1000, tm%1000*1e6)
}

func tmTime2Timestamp(tm time.Time) int64 {
	return tm.UnixNano() / 1e6
}

func helloTime1() {
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

func helloTime2() {
	dtNow := time.Now()
	fmt.Println(dtNow.Unix())
	fmt.Println(dtNow.UnixNano() / 1e6)
	fmt.Println(dtNow.UnixNano())
}

func tmGetDayTimestamp(tm time.Time) (int64, int64) {
	tm1 := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	return tm1.UnixNano() / 1e6, tm2.UnixNano() / 1e6
}

func tmGetCurrentDayTimestamp() (int64, int64) {
	return tmGetDayTimestamp(time.Now())
}

func helloTime3() {
	dtBegin, dtEnd := tmGetCurrentDayTimestamp()
	fmt.Println(dtBegin, dtEnd, dtEnd - dtBegin)
}

func helloTime4() {
	tmBegin,_ := tmGetCurrentDayTimestamp()
	for i := 10; i >= 0; i-- {
		fmt.Println(tmBegin - int64(86400000*i))
	}
}

func helloTime5() {
	s1 := "2020-08-21T13:16:49.997"
	s2 := "2020-08-21 13:16:50"
	tm1, _ := time.Parse("2006-01-02T15:04:05.999", s1)
	tm2, _ := time.Parse("2006-01-02 15:04:05", s2)
	fmt.Println(tm1)
	fmt.Println(tm2)
	fmt.Println(tm2.Unix()-tm1.Unix())
}

func helloTime6()  {
	tm := time.Now()
	fmt.Println(tm.AddDate(0, 0, 1))
	fmt.Println(tm.AddDate(0, 0, -1))
	fmt.Println(tm.AddDate(0, 0, -2))
	fmt.Println(tm.AddDate(0, 0, 15))
	fmt.Println(tm.AddDate(0, 0, 16))
	fmt.Println(tm.AddDate(0, 0, 17))
	fmt.Println(tm.AddDate(0, 0, 46))
	fmt.Println(tm.AddDate(0, 0, 47))
	fmt.Println(tm.AddDate(0, 0, 48))
}

func helloTime7() {
	tmBegin, tmEnd  := tmGetDayTimestamp(time.Now().AddDate(0, 0, -20))
	tm1 := tmTimestamp2Time(tmBegin)
	tm2 := tmTimestamp2Time(tmEnd)
	fmt.Println(tm1)
	fmt.Println(tm2)
	fmt.Println(tm1.Format("2006-01-02 15:04:05"))
	fmt.Println(tm2.Format("2006-01-02 15:04:05"))
}

func main() {
	// helloTime6()
	//for {
	//	time.Sleep(2*time.Second)
	//	fmt.Println(time.Now())
	//}
	//helloTime7()
	for {
		fmt.Println(time.Now().Hour())
		time.Sleep(1*time.Second)
	}
}

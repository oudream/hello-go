package main

import (
	"fmt"
	"strings"
)

// 1代表已連接 0代表未連接 ，從右到左順序為 ES Redis DB
// 111 (十進制7)代表都正常
// 110(十進制6)代表ES異常
// 101(十進制5)代表Redis異常
// 011(十進制3)代表DB異常
// 100(十進制4)代表ES和Redis異常
// 010(十進制2)代表ES和DB異常
// 001(十進制1)代表Redis和DB異常
// 000(十進制0)代表都全部連接異常

func getStateText(stateCode int) string {
	var alerts []string
	if stateCode & 0x01 != 0x01 {
		alerts = append(alerts, "ES異常")
	}
	if stateCode & 0x02 != 0x02 {
		alerts = append(alerts, "Redis異常")
	}
	if stateCode & 0x04 != 0x04 {
		alerts = append(alerts, "DB異常")
	}
	return strings.Join(alerts, ", ")
}

func main() {
	fmt.Println(getStateText(0xFA6))
	fmt.Println(getStateText(0xFA7))
	fmt.Println(getStateText(0XFA5))
	fmt.Println(getStateText(0XFA4))
}

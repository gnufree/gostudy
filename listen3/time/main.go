package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	send := now.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)

	timestamp := now.Unix()
	fmt.Printf("timestamp is :%d\n", timestamp)
}

func main() {
	testTime()
}

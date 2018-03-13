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

func testTimestamp(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	send := timeObj.Second()
	fmt.Printf("current timestamp:%d\n", timestamp)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)
}

func testTicker() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
	}
}

func testConst() {
	fmt.Printf("nano second:%d\n", time.Nanosecond)  // 纳秒
	fmt.Printf("mico second:%d\n", time.Microsecond) // 微妙
	fmt.Printf("mill second:%d\n", time.Millisecond) // 毫秒
	fmt.Printf("second second:%d\n", time.Second)    // 秒
}

func testFormat() {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	fmt.Printf("time:%s\n", timeStr)
}

func testFormat1() {
	now := time.Now()
	timeStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("currec time:%s\n", timeStr)
}

func testCost() {
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	cos := (end - start) / 1000
	fmt.Printf("cost time: %d us\n", cos)
}
func main() {
	// testTime()
	// timestamp := time.Now().Unix()
	// testTimestamp(timestamp)
	// testTicker()
	// testConst()
	// testFormat()
	// testFormat1()
	testCost()
}

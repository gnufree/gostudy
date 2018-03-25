package main

import (
	"fmt"
)

func justify(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func example1() {
	for i := 2; i < 100; i++ {
		if justify(i) == true {
			fmt.Printf("[%d] is prime\n", i)
		}
	}
}

func is_shuaixianhua(n int) bool {
	first := n % 10
	second := (n / 10) % 10
	third := (n / 100) % 10
	// fmt.Printf("n: %d first:%d second: %d third:%d\n",n, first, second, third)
	sum := first*first*first + second*second*second + third*third*third
	if sum == n {
		return true
	}
	return false
}

func cale(str string) (chartCount int, numCount int, spaceCount int, otherCount int) {
	utfCharts := []rune(str)
	for i := 0; i < len(utfCharts); i++ {
		if utfCharts[i] > 'a' && utfCharts[i] < 'z' || utfCharts[i] > 'A' && utfCharts[i] < 'Z' {
			chartCount++
			continue
		}
		if utfCharts[i] >= '0' && utfCharts[i] <= '9' {
			numCount++
			continue
		}
		if utfCharts[i] == ' ' {
			spaceCount++
			continue
		}
		otherCount++
	}
	return
}

func example2() {
	for i := 100; i < 1000; i++ {
		if is_shuaixianhua(i) == true {
			fmt.Printf("shui xian hua num: %d\n", i)
		}
	}
}

func example3() {
	var str string = "dfgs     我爱天安门 22324234"
	chartCount, numCount, spaceCount, otherCount := cale(str)
	fmt.Printf("chartCount:%d numCount:%d spaceCount:%d otherCount:%d \n", chartCount, numCount, spaceCount, otherCount)

}
func main() {
	// example1()
	// example2()
	example3()
}

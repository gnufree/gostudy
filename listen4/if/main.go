package main

import (
	"fmt"
)

func testIf1() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("num is 偶数 %d\n", num)
	} else {
		fmt.Printf("num is 奇数 %d\n", num)
	}
}

func testIf2()  {
	if num := 11; num % 2 == 0 {
		fmt.Printf("num is 偶数 %d\n", num)
	} else  {
		fmt.Printf("num is 奇数 %d\n", num)
	}
}

func getNum() int {
	return 23
}

func testIf3()  {
	if num := getNum(); num % 2 == 0 {
		fmt.Printf("num is 偶数 %d\n", num)
	} else  {
		fmt.Printf("num is 奇数 %d\n", num)
	}
}

func testIf4()  {
	num := 99
	if num % 40 >= 20 {
		fmt.Printf("num is 偶数 %d\n", num)
	} else if num % 40 >= 2 {
		fmt.Printf("num is 奇数 %d\n", num)
	} else {
		fmt.Printf("num is 111 %d\n", num)
	}
}
func main() {
	// testIf1()
	// testIf2()
	// testIf3()
	testIf4()
}

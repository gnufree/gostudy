package main

import (
	"fmt"
)

func testFor(n int) {
	var i int
	for i = 1; 1 <= n; i++ {
		fmt.Printf("i = %d\n", i)
		if i > n {
			break
		}
	}
	// return n
}

func testD(n int) {
	if n > 0 {
		fmt.Println(n)
		fmt.Printf("n = %d\n",n)
	}
}

func main() {
	var n int 
	n = 10
	// testFor(n)
	testD(n)
}
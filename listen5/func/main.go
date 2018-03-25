package main

import (
	"fmt"
)

func testFunc(a int, b ...int) int {

	a = 0
	for i := 0; i < len(b); i++ {
		a = a + b[i]
	}
	return a
}
func main() {
	sum := testFunc(1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 32)
	fmt.Printf("sum is:%d\n", sum)
}

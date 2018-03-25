package main

import (
	"fmt"
)

func sunArray(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}
	return sum
}
func main()  {
	var a [10]int = [10]int{10,2,4,5,9,22,12,23,32,21}
	b := sunArray(a)
	fmt.Printf("b=%d\n",b)
}
package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func testFunc() {
	f1 := add
	fmt.Printf("type of f1=%T\n", f1)
	sum := f1(2, 5)
	fmt.Printf("sum is %d\n", sum)
}

func testFunc2() {
	f1 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("type of f1=%T\n", f1)
	sum := f1(3, 5)
	fmt.Printf("sum is %d\n", sum)
}

func testFunc3() {
	var i int = 100
	defer fmt.Printf("defer i = %d\n", i)
	i = 200
	fmt.Printf("i = %d\n", i)
}

func testFunc4() {
	var i int = 100
	defer func() {
		fmt.Printf("defer i = %d\n", i)
	}()
	i = 300
	fmt.Printf("i = %d\n", i)
}

var f = func(a, b int) int {
	return a + b
}

func f1(i int) func() int {
	return func() int {
		i++
		return i
	}
}
func main() {
	// testFunc()
	// testFunc2()
	// testFunc3()
	// testFunc4()
	// fmt.Printf("f=%d\n", f(3, 5))
	c1 := f1(0)
	c2 := f1(0)
	fmt.Println(c1)
	fmt.Println(c2)

}

package main

import (
	"fmt"
)

func testDefer1() {
	defer fmt.Println("aaaa")
	defer fmt.Println("bbb")
	defer fmt.Println("cccc")
	fmt.Println("fff")
	fmt.Println("ggg")
}

func testDefer2()  {
	for i :=0; i < 10; i++ {
		defer fmt.Printf("i=%d\n", i)
	}

	fmt.Printf("running\n")
}

func main() {
	// testDefer1()
	testDefer2()
}

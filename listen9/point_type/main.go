package main

import (
	"fmt"
)

func testPoint1() {
	var a int32
	a = 1000
	fmt.Printf("the addr of a:%p, a:%d\n ", &a, a)

	var b *int
	if b == nil {
		fmt.Println("b is nil addr")
	}
	fmt.Printf("the addr of	b: %p b:%p\n", &b, b)
	// *b = 100
}

func testPoint2() {
	var a int = 200
	var b *int = &a

	fmt.Printf("the addr b:%d\n ", *b)
	*b = 1000
	fmt.Printf("b 指向的地址存储的值为:%b\n", *b)
	fmt.Printf("a = %d\n", a)
}
func modify(val *int) {
	*val = 100
}

func modify_arr(val *[3]int) {
	(*val)[0] = 100
}

func testPoint3() {
	var b int = 10
	p := &b
	modify(p)
	fmt.Printf("b:%d\n", b)

}

func testPoint4() {
	var b [3]int = [3]int{1, 2, 3}
	modify_arr(&b)
	fmt.Printf("b is :%v\n", b)
}
func main() {
	// testPoint1()
	// testPoint2()
	// testPoint3()
	testPoint4()
}

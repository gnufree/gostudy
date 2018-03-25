package main

import (
	"fmt"
	"strings"
	"time"
)

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return d
	}
}

func add(base int) func(int) int {
	return func(b int) int {
		base += b
		return base
	}
}

func testClosure() {
	f := Adder()
	ret := f(1)
	fmt.Printf("f(1) ret = %d\n", ret)
	ret = f(21)
	fmt.Printf("f(21) ret = %d\n", ret)
	ret = f(100)
	fmt.Printf("f(100) ret = %d\n", ret)

}

func testadd() {
	f := add(20)
	ret := f(1)
	fmt.Printf("ret = %d\n", ret)
	ret = f(211)
	fmt.Printf("f(211) ret = %d\n", ret)
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func testmakeSuffixFunc() {
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	func3 := makeSuffixFunc("test.gif")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
	fmt.Println(func3("test.gif"))

}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func testcalc() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(3))
	fmt.Println(f1(5), f2(8))
	fmt.Println(f1(8), f2(9))
}

func func5() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}
	time.Sleep(time.Second)
}
func main() {
	// testClosure()
	// testadd()
	// testmakeSuffixFunc()
	// testcalc()
	func5()
}

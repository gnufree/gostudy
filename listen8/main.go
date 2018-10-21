package main

import (
	"fmt"
)

func testAppend() {
	//定义切片并赋值
	var a []int = []int{10, 20, 30}
	fmt.Printf("a = %v\n", a)
	b := []int{11, 22, 33}
	fmt.Printf("b = %v\n", b)
	var c []int
	//使用make函数初始化切片
	c = make([]int, 5, 10)
	fmt.Println(c)
	c[0] = 100
	fmt.Println(c)
	b = append(b,c...)
	fmt.Println(b)
	var dd [][]int
	dd = make([][]int, 2, 2)
	dd[0] = make([]int, 2, 2)
	dd[0][0] = 1
	dd[1] = make([]int,2,2)
	dd[1][0] = 2
	parms := []int{10,20,30,40,50}
	// 使用切片创建新切片
	var gg []int = parms[1:3]

	fmt.Println(dd)
	fmt.Println(parms)
	fmt.Println(gg)

}

func printSlice(s []int)  {
	fmt.Printf("s = %v len=%d cap=%d\n", s, len(s),cap(s))
}

func main() {
	// testAppend()
	s := []int{10,20,30,40,50,60}
	printSlice(s)
	s = s[0:4]
	printSlice(s)
	s = s[1:3]
	printSlice(s)
}


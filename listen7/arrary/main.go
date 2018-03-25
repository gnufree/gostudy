package main

import (
	"fmt"
)

func testArrary1() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
}

func testArrary2() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
}

func testArrary3() {
	a := [5]int{1, 2, 3}
	fmt.Println(a)
}

func testArrary4() {
	var a [3]int = [3]int{2, 3, 4}
	fmt.Println(a)
}

func testArrary5() {
	a := [5]int{3: 100, 2: 100, 4: 213}
	fmt.Println(a)
}

func testArrary6() {
	a := [5]int{3: 100, 2: 100, 4: 213}
	fmt.Println(a)
	var b [5]int
	b = a
	fmt.Println(b)
}

func testArrary7() {
	a := [5]int{3: 100, 2: 100, 4: 213}
	fmt.Println(a)
	var b [5]int
	b = a
	fmt.Println(b)
	fmt.Printf("b len(%d)\n", len(b))
}

func testArrary8() {
	a := [5]int{3: 100, 2: 100, 4: 213}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}

func testArrary9() {
	a := [5]int{3: 100, 2: 100, 4: 213}
	for _, val := range a {
		// fmt.Printf("index=%d val=%d \n", index, val)
		fmt.Printf("value:%d\n", val)
	}
}

func testArrary10() {
	var a [3][2]int
	a[0][0] = 10
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 40
	a[2][0] = 50
	a[2][1] = 60

	// fmt.Println(a)
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Printf("%d ", a[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// other mother
	for i, val := range a {
		fmt.Printf("%d %d\n", i, val)
		for j, val2 := range val {
			fmt.Printf("(%d,%d)=%d\n", i, j, val2)
		}
		fmt.Println()
	}
}

func testArrary11() {
	var a [3]int = [3]int{10, 20, 30}
	b := a
	b[0] = 1000
	fmt.Printf("a=%d", a)
	fmt.Printf("b=%d", b)
}

func modif(b [3]int) {
	b[1] = 2000
	fmt.Printf("b=%v\n", b)
}

func testArrary12() {
	var a [3]int = [3]int{10, 20, 30}
	modif(a)
	fmt.Printf("a=%v\n", a) // 数组是值类型， 修改数据内容不会改变原数组
}

func testSort(b [5]int) [5]int {
	for i := 1; i < len(b); i++ {
		for j := i; j > 0; j-- {
			if b[j] < b[j-1] {
				b[j], b[j-1] = b[j-1], b[j]
			} else {
				break
			}
		}
	}
	return b
}
func testArrary13() {
	var a [5]int = [5]int{1, 3, 5, 8, 7}
	c := testSort(a)
	fmt.Println(c)
	// for index, val := range c {
		
	// }
	if c[0]+c[4] = 8 && c[1] + c[2] = 8 {
		fmt.Printf("%d + %d = 8\n", c[0],c[4])
		fmt.Printf("%d + %d = 8\n", c[1],c[2])
	}
}
func main() {
	// testArrary1()
	// testArrary2()
	// testArrary3()
	// testArrary4()
	// testArrary5()
	// testArrary6()
	// testArrary7()
	// testArrary8()
	// testArrary9()
	// testArrary10()
	// testArrary11()
	// testArrary12()
	testArrary13()
}

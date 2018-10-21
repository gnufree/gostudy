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
	// if c[0]+c[4] = 8 && c[1] + c[2] = 8 {
	// 	fmt.Printf("%d + %d = 8\n", c[0],c[4])
	// 	fmt.Printf("%d + %d = 8\n", c[1],c[2])
	// }
}

func testArrary14()  {
	array := [...]int{1:3330,2:20,30,40,50}
	bb := [5]int{0:22,1:22,4:232}
	fmt.Println(len(array))
	fmt.Println(array)
	fmt.Println(bb)
	cc := [5]int{10,20,30,40,50}
	cc[2] = 35
	fmt.Println(cc)
	// 声明包含5个元素的指向整数的数组
	// 用整型指针初始化索引为0和1的数组元素
	dd := [5]*int{0:new(int), 1:new(int)}

	// 为索引为0 和1的元素赋值
	*dd[0] = 10
	*dd[1] = 20
	fmt.Println(dd)
}

func testArray1()  {
	var array1 [5]string

	// 声明第二个包含5个元素的字符串数组
	// 用颜色初始化数组
	array2 := [5]string{"red","blue","green","Yellow","Pink"}
	array1 = array2
	fmt.Printf("array1=%v, array2=%v", array1, array2)
}

func testArray2()  {
	// 声明第一个包含3个元素的指向字符串的指针的数组
	var array1 [3]*string

	//声明第二个包含3个元素的指向字符串的指针数组
	//使用字符串指针初始化这个数组
	array2 := [3]*string{new(string),new(string),new(string)}

	//使用颜色为每个元素赋值
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	array1 = array2

	fmt.Println(array1,array2)

}

func testArray3()  {
	// 声明一个二维整型数组,两个维度分别存储4个元素和2个元素
	var array [4][2]int
	
	// 使用数组字面量来声明并初始化一个二维整型数组
	array2 := [4][2]int{{10,11},{20,21},{30,31},{40,41}}

	//声明并初始化外层数组中索引为1和3的元素
	array1 := [4][2]int{1:{20,21}, 3:{40,41}}
	fmt.Printf("array1=%v\n array2=%v\n array1=%v",array,array2,array1)

	// 声明两个不同的二维整数数组
	var a1 [2][2]int
	var a2 [2][2]int

	// 为每个元素赋值
	a2[0][0] = 10
	a2[0][1] = 20
	a2[1][0] = 30
	a2[1][1] = 40

	//将a2的值复制给a1
	a1 = a2
	//因为每个数组都是一个值，所以可以独立复制某个维度，如代码清单4-13所示。

	//将a1 的索引为1的维度复制管道一个同类型的新数组里
	var a3 [2]int = a1[1]

	//将外层数组的索引为1,内层数组的索引为0的整数值复制到新的整型变量里
	var value int = a1[1][0]
	fmt.Println()
	fmt.Printf("a1=%v\n a2=%v\n a3=%v\n value=%v\n",a1,a2,a3,value)
	// 在函数间传递数组 如果传递非常大的值是会占用很多的内存空间，最好使用指针数组传递，不过使用指针数组传递，如果改变指针指向的值，会改变共享的内存。

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
	// testArrary13()
	// testArrary14()
	// testArray1()
	// testArray2()
	testArray3()
}

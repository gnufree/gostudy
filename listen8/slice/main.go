package main

import (
	"fmt"
)

func testSlice1()  {
	//Go语言中有多种创建切片和初始化切片的方法，是否提前知道切片需要的容量通常会决定要使用哪种方法创建切片

	//make函数创建方法
	//创建一个字符串切片
	// 其长度和容量都是5个元素
	slice := make([]string, 5)
	// 创建一个整型切片
	// 其长度为3个元素，容量为5
	slice1 := make([]int, 3, 5)

	//通过切片字面量来声明切片
	//创建字符串切片
	// 其长度和容量都是5个元素
	slice2 := []string{"Red","Blue","Green","Yellow","Pink"}
	// 创建一个整型切片
	// 其长度和容量都是3个元素
	slice3 := []int{10,20,30}
	// 创建字符串切片
	// 使用空字符串初始化第100个元素
	slice4 := []string{99: " ff"}

	fmt.Printf("slice=%v slice1=%v\n", slice, slice1)

	fmt.Println(slice2,slice3,slice4)
	// 创建nil整型切片
	var slice5 []int
	//使用make创建空的整型切片
	slice6 := make([]int, 0)

	//使用切片字面量创建空的整型切片
	slice7 := []int{}

	fmt.Printf("slice5=%v \nslice6=%v slice7=%v\n", slice5, slice6, slice7)


}

func testSlice2()  {
	// 创建一个整型切片
	// 其容量和长度都是5个元素
	slice := []int{10,20,30,40,50}
	// 改变索引为1的元素的值
	slice[1] = 25
	fmt.Printf("slice=%v len=%v cap=%v\n", slice, len(slice), cap(slice))
	// 使用切片创建切片
	// 创建一个新切片
	// 其长度为2个元素,容量为4个元素
	newSlice := slice[1:3]
	fmt.Printf("newSlic=%v len=%v cap=%v\n", newSlice, len(newSlice), cap(newSlice))
	// 如何计算长度和容量
	// 底层数据容量是k 的切片slice[i:j]来说
	//长度: j - i
	//容量: k - i
	//例如：slice[1:3]  其长度为：3-1 = 2 其容量为：5 - 1 = 4 因为5是其底层数据的容量

	//修改切片内容可能导致的后果
	// 修改newSlice索引为1的元素
	// 同时修改了原来的slice的索引为2的元素
	newSlice[1] = 35
	// 这里两个切片共享同一个底层数组，如果一个切片修改了底层数组的共享部分，另一个切片也能感知到。
	fmt.Printf("slice=%v, newSlice=%v\n", slice, newSlice)
	// 表示索引越界的语言运行时的错误
	// newSlice[3] = 45
	// fmt.Println(newSlice)

}

// 使用append向切片增加元素
func testSlice3()  {
	//创建一个整型切片
	// 其长度和容量都是5个元素
	slice := []int{10, 20, 30, 40, 50}

	//创建一个新切片
	//其长度为2，容量为4个元素
	newSlice := slice[1:3]
	fmt.Printf("slice=%v newSlice=%v len=%v cap=%v \n", slice, newSlice, len(newSlice),cap(newSlice))
	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为60
	newSlice = append(newSlice, 60)
	fmt.Printf("newSlice=%v slice=%v len=%v cap=%v \n", newSlice, slice, len(newSlice), cap(newSlice))

}

func testSlice4()  {
	//使用append 同时增加切片的长度和容量
	//创建一个整型切片
	// 其长度和容量都是4个元素
	slice := []int{10, 20, 30, 40}
	//向切片追加一个新元素
	// 将新元素赋值为50
	newSlice := append(slice, 50)
	// 当这个append 完成后，newSlice拥有一个全新的底层数组，这个数组的容量是原来的两倍
	fmt.Printf("slice=%v len=%v cap=%v \n", slice, len(slice), cap(slice))
	fmt.Printf("newSlice=%v len=%v cap=%v\n", newSlice, len(newSlice), cap(newSlice))
}

func testSlice5()  {
	//创建切片时的第三个索引 ，目的不是要增加容量，而是要限制容量
	//slice[1:3:4]
	// 使用切片字面量声明一个字符串切片
	// 创建字符串切片
	// 其长度和容量都是5个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	//查看切片的值
	fmt.Printf("souce=%v\n", source)
	// 现在试着用第三个索引选项来完成切片操作.
	// 将第三个元素切片,并限制容量
	// 其长度为1个元素，容量为2个元素
	slice := source[2:3:4]
	// 如何计算长度和容量
	//对于slice[i:j:k]或[2:3:4]
	//计算公式如下:
	// 长度: j-i 或 3 - 2 = 1
	// 容量: k-i 或 4 - 2 = 2
	fmt.Printf("slice=%v len=%v cap=%v \n", slice, len(slice), cap(slice))

	//设置容量大于已有容量的语言运行时错误
	// 这个切片操作试图设置容量为4
	// 这比可用的容量大
	// blice := source[2:3:6]

	// fmt.Printf("blice =%v\n", blice)
	//设置长度和容量一样的好处
	//创建字符串切片
	// 其长度和容量都是5个元素
	source2 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Printf("source2=%v len=%v cap=%v \n", source2, len(source2), cap(source2))
	
	// 对第三个元素做切片，并限制容量
	// 其长度和容量都是1个元素
	slice2 := source2[2:3:3]
	fmt.Printf("slice2=%v len=%v cap=%v \n", slice2, len(slice2), cap(slice2))
	//向slice2 追加新字符串
	slice2 = append(slice2, "Kiwi")
	fmt.Printf("slice2=%v len=%v cap=%v \n", slice2, len(slice2), cap(slice2))
	fmt.Printf("source2=%v len=%v cap=%v \n", source2, len(source2), cap(source2))

	//展示append 可变参数函数特性 如果使用... 运算符，可以将一个切片的所有元素追加到另一个切片里
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))
	//Output [1 2 3 4] 通过输出可以看出，切片s2 里的所有值追加到了切片s1的后面

}

func testForSlice() {
	// 使用for range 迭代切片
	// 创建一个整型切片
	// 其长度和容量偶数4个元素
	slice := []int{10, 20, 30, 40}

	// 迭代每一个元素,并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	// Output:
	// Index: 0 Value: 10
	// Index: 1 Value: 20
	// Index: 2 Value: 30
	// Index: 3 Value: 40
	// 当迭代切片时，关键值range 会返回两个值。第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本
	// range 创建了每个元素的副本，而不是直接返回对该元素的引用，如下面代码，如果使用该值变量的地址作为指向每个元素的指针，就会造成报错。
	for key, val := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", val, &val, &slice[key])
	}

	// 因为迭代返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以value的地址总是相同的，要想获取每个元素的地址，可以使用切片变量和索引值slice[index]
	// 如果不想使用索引值，可以使用占位符胡洛 _
	for _, v := range slice {
		fmt.Printf("value: %d\n", v)
	}
	// 关键字range 总是从切片头部开始迭代，如果想对迭代做更多的控制，可以使用传统的for 循环
	// 使用传统的for 循环对切片进行迭代
	// 创建一个整型切片
	// 其长度和容量都是4个元素
	s1 := []int{10, 20, 30, 40}
	for index := 2; index < len(s1); index++ {
		fmt.Printf("Index: %d Valuse: %d\n", index, slice[index])
	}
}

func testMustSlice()  {
	// 创建一个整型切片的切片
	slice := [][]int{{10}, {100,200}}
	fmt.Printf("slice=%v\n", slice)
	// 为第一个切片追加值为20的元素
	slice[0] = append(slice[0], 20)
	fmt.Printf("slice=%v\n", slice)
	slice[0] = append(slice[0],30)
	fmt.Printf("slice=%v\n", slice)
	slice[1] = append(slice[1],30)
	fmt.Printf("slice=%v\n", slice)

}


func testSliceFunc()  {
	//分配包含100万个整型值的切片
	slice := make([]int, 1e6)
	//将slice 传递给函数foo
	slice = foo(slice)
	fmt.Printf("slice=%v\n", slice)
}

func foo(slice []int) []int {
	return slice
}

func testMake1()  {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 10
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}
	// 观察切片的扩容操作，扩容的策略是翻倍扩容
	a = append(a, 1000)
	fmt.Printf("扩容之后的地址： a= %v addr:%p len:%d cap:%d \n", a, a, len(a),cap(a))
}

func testMake2() {
	var a []int
	a = make([]int, 5, 10)

	a = append(a, 10)
	fmt.Printf("a=%v \n", a)
	// 初始化一个切片， 其长度为0 容量为10
	b := make([]int, 0, 10)
	fmt.Printf("b=%v len:=%d cap:%d \n", b, len(b), cap(b))
	// 给切片添加一个元素，并打印其元素和长度 容量
	b = append(b, 100)
	fmt.Printf("b=%v len:%d cap:%d \n", b, len(b), cap(b))

	b = append(b, 1, 101)
	fmt.Printf("b=%v len:%d cap:%d \n", b, len(b), cap(b))

}

func testCap()  {
	a := [...]string {"a", "b", "c", "d", "e", "f","g", "h"}
	b := a[:4]
	fmt.Printf("b=%v len=%d cap=%d\n", b, len(b), cap(b))
}

func testCap2()  {
	a := [...]string {"a", "b", "c", "d", "e", "f","g", "h"}
	b := a[1:3]
	fmt.Printf("b=%v len=%d cap=%d\n", b, len(b), cap(b))
	b = b[:cap(b)]
	fmt.Printf("b=%v len=%d cap=%d\n", b, len(b), cap(b))
	
}

func testSlice()  {
	var a [] int
	fmt.Printf("addr:%p, len:%d cap:%d \n", a, len(a), cap(a))
	if a == nil {
		fmt.Printf("a is nil\n")
	}
	// 使用append 追加元素， 切片的容量成本增长
	a = append(a, 100)
	fmt.Printf("a=%v addr:%p, len:%d cap:%d \n", a, a, len(a), cap(a))
	
	a = append(a, 200)
	fmt.Printf("a=%v addr:%p, len:%d cap:%d \n", a, a, len(a), cap(a))
	a = append(a, 300)
	fmt.Printf("a=%v addr:%p, len:%d cap:%d \n", a, a, len(a), cap(a))
	a = append(a, 400)
	fmt.Printf("a=%v addr:%p, len:%d cap:%d \n", a, a, len(a), cap(a))
	a = append(a, 500)
	fmt.Printf("a=%v addr:%p, len:%d cap:%d \n", a, a, len(a), cap(a))
	
}

func testAppend()  {
	var a []int = []int{1,2,3}
	var b []int = []int{4,5,6}

	a = append(a, 23, 24, 25)
	fmt.Println("a=%v",a)
	// 切片最近切片 append(a, b...)
	a = append(a, b...)
	fmt.Println("a=%v",a)
	
}

func sunArray(a []int) int {
	var sum int = 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func testSumArray()  {
	var a [10]int = [10]int{1, 3, 3, 4, 5, 5, 8}
	sum := sunArray(a[:])
	fmt.Printf("sum:%v", sum)
}

func main()  {
	// testSlice1()
	// testSlice2()
	// testSlice3()
	// testSlice4()
	// testSlice5()
	// testForSlice()
	// testMustSlice()
	// testSliceFunc()
	// testMake1()
	// testMake2()
	// testCap()
	// testCap2()
	// testSlice()
	// testAppend()
	testSumArray()


}
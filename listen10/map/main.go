package main

import (
	"fmt"
	"math/rand"
)

// 根据ID查询学生信息
// 学生信息包括ID 姓名 分数 年龄
func testMap() {
	var stdMap map[int]map[string]interface{}
	stdMap = make(map[int]map[string]interface{})
	// 插入学生信息 id=1 name="std01" score=97.5 age=18
	for i := 0; i < 10; i++ {
		value, ok := stdMap[i]
		if !ok {
			value = make(map[string]interface{})
			value["name"] = fmt.Sprintf("std%d", i)
			value["id"] = i
			value["score"] = rand.Float64() * 100.0
			value["age"] = rand.Intn(30)
			stdMap[i] = value
		}
	}
	for k, v := range stdMap {
		fmt.Printf("id = %d, stu info=%#v\n", k, v)
	}

}
func testInterface() {
	var a interface{}
	var b int = 10
	var c string = "hello"
	var d float32 = 3.23
	a = b
	fmt.Printf("a = %v\n", a)
	a = c
	fmt.Printf("a = %v\n", a)
	a = d
	fmt.Printf("a = %v\n", a)
}

func main() {
	testMap()
	// testInterface()
}

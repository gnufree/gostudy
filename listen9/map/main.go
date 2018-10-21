package main

import (
	"strings"
	"fmt"
)

func chartCount(str string) map[string]int {
	var m map[string]int = make(map[string]int)
	ss := strings.Split(str, " ")
	for _, val := range ss {
		count, ok := m[val]
		if !ok {
			m[val] = 1
		}
		count = m[val] + 1
		m[val] = count
	}
    // fmt.Printf("m =%v \n",m)
	return m
}

func testMap()  {
	var m1 map[string]int = make(map[string]int)
	keys := "how"
	m1[keys] = 1
	_, ok := m1[keys] 
	if ok == true {
		fmt.Printf("m1 de keys is exist %v\n",m1[keys])
	} else {
		fmt.Printf("m1 key is not exist! %v\n",m1[keys])
	}
	fmt.Printf("m1 = %v \n", m1)

}


func studentInfo(id int) {
	var sInfo map[int](map[string]int) = make(map[int](map[string]int))
	sInfo[id]map[string]int = []int{1,3}
	fmt.Println(sInfo)
}

func main()  {
	// var s string = "how do you do you a shi shui aa wo hh do"
	// testMap()
	// fmt.Println(chartCount(s))
	studentInfo()

}
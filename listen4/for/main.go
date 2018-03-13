package main

import (
	"fmt"
)

func  testFor1()  {
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Printf("i=%d\n",i)
	}

	fmt.Printf("i=%d\n",i)
}

func  testFor2()  {
	var i int
	for i = 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("i=%d\n",i)
	}

	fmt.Printf("i=%d\n",i)
}


func  testFor3()  {
	var i int
	for i = 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Printf("i=%d\n",i)
	}

	fmt.Printf("i=%d\n",i)
}
func main()  {
	// testFor1()
	// testFor2()
	testFor3()
}
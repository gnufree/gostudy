package main

import (
	"fmt"
)

var a int = 100

func testGlobleVarible() {
	fmt.Printf("a=%d b=%d\n", a,b)
}

func testLocalVarible()  {
	var b int = 200
	fmt.Printf("a=%d b=%d\n",a,b)
}
func main() {
	// testGlobleVarible()
	testLocalVarible()

}

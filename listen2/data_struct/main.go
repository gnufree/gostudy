package main

import (
	"fmt"
	"strings"

	"github.com/gnufree/gostudy/listen2/access"
)

func testBool() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	a = !a
	fmt.Println(a)

	var b bool = true
	// && 与操作
	if a == true && b == true {
		fmt.Println("right")
	} else {
		fmt.Println("not right")
	}
	if a == true || b == true {
		fmt.Println("|| right")
	} else {
		fmt.Println("|| not right")
	}
	fmt.Printf("%t %t\n", a, b)
}

func testInt() {
	var a int8
	a = 18
	fmt.Println("a= ", a)
	a = -12
	fmt.Println("a=", a)
	var b int
	b = 1323232323
	fmt.Println("b= ", b)

	var c = 5.6
	fmt.Println(c)

	fmt.Printf("a=%d\n b=%x\n c=%f", a, a, c)
}

func testStr() {
	var a = "a:hello"
	fmt.Println(a)

	b := a
	fmt.Println(b)

	c := "c:\nhello"
	fmt.Println(c)
	c = `
	床前明月光，
	疑是地上霜，

	`
	fmt.Printf("a=%v b=%v c=%v", a, b, c)
	clen := len(c)
	fmt.Printf("len of c =%d\n", clen)

	c = c + c
	fmt.Printf("c = %s", c)

	c = fmt.Sprintf("%s%s", c, c)
	fmt.Printf("c=%s\n", c)

	ips := "10.3.2.1;10.3.2.53"
	ipArray := strings.Split(ips, ";")
	fmt.Printf("first ip :%s\n", ipArray[0])
	fmt.Printf("second ip :%s\n", ipArray[1])

	result := strings.Contains(ips, "10.3.2.53")
	fmt.Println(result)

	urls := "http://baidu.com"
	if strings.HasPrefix(urls, "http") {
		fmt.Println("urls is http url")
	} else {
		fmt.Println("urls is not http url")
	}

	if strings.HasSuffix(urls, "baidu.com") {
		fmt.Println("urls is baidu web")
	} else {
		fmt.Println("urls is not baidu web")
	}

	index := strings.Index(urls, "baidu")
	fmt.Printf("baidu is index:%d\n", index)

	index = strings.LastIndex(urls, "baidu.com")
	fmt.Printf("baidu is index:%d\n", index)

	var str []string = []string{"10.3.2.1", "10.3.2.2", "10.3.2.22"}
	resultStr := strings.Join(str, ";")
	fmt.Printf("result=%s\n", resultStr)

}

func testOperator() {
	var a int = 2
	if a != 2 {
		fmt.Printf("is right\n")
	} else {
		fmt.Printf("is not reght\n")
	}
	a = a + 100
	fmt.Printf("a=%d\n", a)
}

func testAccess() {
	fmt.Printf("access.a=%d\n", access.A)
}
func main() {
	// testBool()
	// testInt()
	// testStr()
	// testOperator()
	testAccess()
}

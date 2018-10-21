package main

import (
	"fmt"
)

// struct 定义
type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	// struct 初始化
	var user User
	user.Username = "zhangsan"
	user.Sex = "男"
	user.Age = 20
	user.AvatarUrl = "http://www.baidu.com/image/xxx.jpg"
	fmt.Printf("user.username=%s age=%d sex=%s avatarUrl=%s\n", user.Username, user.Age, user.Sex, user.AvatarUrl)
	// struct 第二种初始化
	var user2 User = User{
		Username: "zhiwenjun",
		// Sex: "男",
		Age:       29,
		AvatarUrl: "http://www.baidu.com/image/zhiwenjun.jpg",
	}
	fmt.Printf("user2.username=%s age=%d avatarUrl=%s\n", user2.Username, user2.Age, user2.AvatarUrl)
	// 更简单的初始化方法
	user3 := User{
		Username: "zhiwenjun",
		Sex:      "男",
		// Age: 29,
		AvatarUrl: "http://www.baidu.com/image/zhiwenjun.jpg",
	}
	fmt.Printf("user3.username=%s sex=%s avatarUrl=%s\n", user3.Username, user3.Sex, user3.AvatarUrl)
}

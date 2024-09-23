package main

import (
	"fmt"
	"strings"
)

func test_test1() {
	//字符串比较
	// a := "hello"
	// b := "bello"
	// fmt.Println(a == b)
	// fmt.Println(a != b)
	// fmt.Println(a > b)

	//是否包含
	name := "golang"
	//fmt.Println(strings.Contains(name, "go"))

	//出现次数
	// fmt.Println(strings.Count(name, "o"))

	//分割字符串
	// fmt.Println(strings.Split(name, "o"))

	//字符串前缀和后缀
	fmt.Println(strings.HasPrefix(name, "go"))
	fmt.Println(strings.HasSuffix(name, "lang"))

	//查询子串出现的位置
	fmt.Println(strings.Index(name, "go"))
}
func main() {
	test_test1()
}

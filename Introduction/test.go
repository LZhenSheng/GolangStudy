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
	name := "golang工程师"
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
	//字节数
	fmt.Println(strings.Index(name, "师"))

	fmt.Println(strings.IndexRune(name, []rune(name)[4]))
	//n<0，所有都替换掉，n为替换个数
	fmt.Println(strings.Replace(name, "go", "java", -1))

	//大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("java"))

	//去掉左右两边特殊符号
	fmt.Println(strings.Trim(" hello go ", " "))
	fmt.Println(strings.Trim("#hello go#", "#"))
	fmt.Println(strings.Trim("#hello go$", "#$"))
	fmt.Println(strings.TrimLeft("#hello go$", "#$"))
	fmt.Println(strings.TrimRight("#hello go$", "#$"))

}
func test_test2() {

}
func main() {
	test_test1()
}

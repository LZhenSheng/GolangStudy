package main

import "fmt"

func add(a int, b int) (sum int) {
	return a + b
}
func main() {
	//普通函数、匿名函数、闭包
	//go函数是一等公民、1、函数本身可以当作变量使用2、匿名函数和闭包3、函数可以满足接口
	//go语言都是值传递
	temp := add(2, 3)
	fmt.Println(temp)
}

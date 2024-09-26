package main

import (
	"fmt"
	"strconv"
)

type MyInt2 int

func (mi MyInt2) string() string {
	return strconv.Itoa(int(mi))
}
func main() {
	//type 1定义结构体2定义类型别名3定义接口4类型定义
	//别名是为了更好的理解代码
	type MyInt = int //类型别名
	var i MyInt = 9
	var j int = 9
	fmt.Println(i + j)
	fmt.Printf("%T\n", i)
	// type MyInt2 int //类型定义,int不能增加方法，但Myint2可以，扩展基础能力
	var ii MyInt2 = 2
	fmt.Sprintf("%T\n", ii)
	fmt.Println(int(ii) + i)
	fmt.Println(ii.string())

	var a interface{} = "abc"
	//断言
	switch a.(type) {
	case string:
		fmt.Println("string")
	}
	m := a.(string)
	fmt.Println(m)
}

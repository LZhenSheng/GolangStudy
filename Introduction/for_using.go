package main

import (
	"fmt"
	"time"
)

func for_api_test1() {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(i)
	}
	i := 0
	for i < 3 {
		time.Sleep(2 * time.Second)
		fmt.Println(i)
		i++
	}
}

// for range字符串、数组、切片、map、channel
func for_api_test() {
	//字符串 字符串的索引 字符串索引的字符拷贝 ，如果不写key返回的是索引
	name := "imooc go"
	for index, value := range name {
		// fmt.Println(index, value)
		fmt.Printf("%d %c\n", index, value)
	}
	for index := range name {
		// fmt.Println(index)
		fmt.Printf("%d\n", index)
	}

}
func main() {
	for_api_test()

}

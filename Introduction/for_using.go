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
// 字符串 字符串的索引 字符串索引的字符拷贝 ，如果不写key返回的是索引
func for_api_test2() {
	//中文可以使用 index,value格式 或者 转为rune
	name := "我爱学习go"
	nameRune := []rune(name)
	for index := range nameRune {
		fmt.Printf("%c", nameRune[index])
	}
	for index := range name {
		// fmt.Println(index, value)
		// fmt.Printf("%d %c\n", index, value)
		fmt.Printf("%c", name[index])

	}
	for index := range name {
		// fmt.Println(index)
		fmt.Printf("%d\n", index)
	}
}

// 数组 数组的索引 数组索引的字符拷贝 ，如果不写key返回的是索引
func for_api_test3() {
	arr := [3]int{1, 3, 4}
	for index, value := range arr {
		// fmt.Println(index, value)
		fmt.Printf("%d %d\n", index, value)
	}
	for index := range arr {
		// fmt.Println(index)
		fmt.Printf("%d\n", index)
	}
}

// 切片 切片的索引 切片索引的字符拷贝 ，如果不写key返回的是索引
func for_api_test4() {
	arr := []int{1, 3, 4}
	for index, value := range arr {
		// fmt.Println(index, value)
		fmt.Printf("%d %d\n", index, value)
	}
	for index := range arr {
		// fmt.Println(index)
		fmt.Printf("%d\n", index)
	}
}

// map map的索引 map索引的字符拷贝 ，如果不写key返回的是map的值
func for_api_test() {
	mapCode := map[string]string{
		"Golang": "my love",
		"Java":   "I can do it",
		"Python": "I am studying",
	}
	for index, value := range mapCode {
		// fmt.Println(index, value)
		fmt.Printf("%s %s\n", index, value)
	}
	for value := range mapCode {
		// fmt.Println(value)
		fmt.Printf("%s\n", value)
	}
}

// channel value返回的是channel接受的数据
func main() {
	for_api_test2()
}

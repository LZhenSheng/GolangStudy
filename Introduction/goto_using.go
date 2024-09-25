package main

import "fmt"

func main() {
	//goto可以实现程序的跳转，主要用于错误跳转，当程序出现错误时，程序跳转到对应的标签处，统一处理
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				goto over
			}
			fmt.Println(i, j)
		}
	}
over:
	{
		fmt.Println("over")
	}
}

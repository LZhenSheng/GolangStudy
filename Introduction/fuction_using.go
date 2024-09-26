package main

import (
	"fmt"
	"sync"
)

func add(a int, b int) (sum int) {
	return a + b
}

//	func add(items ...int) (sum int) {
//		for _, value := range items {
//			sum += value
//		}
//		return sum
//	}
func call(op string, items ...int) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是其他的函数")
		}
	}
	return nil
}
func cal(myfunc func(item ...int) int) int {
	return myfunc()
}

//	func callback(y int, f func(...int) int) {
//		fmt.Println(f(y, 2, 3))
//	}
func callback(y int, f func(int, int) int) {
	fmt.Println(f(y, 2))
}

func auoIncreament() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}
func function_using_test() {
	defer fmt.Println("1")
	defer fmt.Println("2")

	defer fmt.Println("3")

}
func function_using_test1() (ret int) {
	defer func() {
		ret++
	}()
	return 10

}
func function_using_test2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var names map[string]string
	names["go"] = "go"
}
func main() {
	//普通函数、匿名函数、闭包
	//go函数是一等公民、1、函数本身可以当作变量使用2、匿名函数和闭包3、函数可以满足接口
	//go语言都是值传递
	// addFunc := add
	// temp := addFunc(2, 3, 234, 43, 4545, 65, 34, 54)
	// fmt.Println(temp)
	call("+")()
	sum := cal(func(item ...int) int {
		sum := 0
		for _, value := range item {
			sum += value
		}
		return sum
	})
	fmt.Println(sum)
	// callback(1, add)

	callback(1, func(a, b int) int {
		fmt.Println("total is:%d", a+b)
		return a + b
	})
	//匿名函数
	localFunc := func(a, b int) {
		fmt.Println("total is:%d", a+b)
	}
	localFunc(2, 4)

	//闭包
	//有个需求，希望有个函数每次调用一次返回的结果值都是增加一次之后的值
	nextFunc := auoIncreament()
	for i := 0; i < 5; i++ {
		fmt.Println(nextFunc())
	}
	//连接数据库，打开锁、打开文件，希望最后关闭数据库、关闭文件、关闭锁
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()
	//defer先后调用顺序是按栈的方式执行的
	function_using_test()
	//defer有能力修改返回值
	fmt.Println(function_using_test1())

	//error panic recover
	//防御性编程，需要有一个返回值去告诉调用者是否成功，go设计者必须要处理这个err
	//panic这个程序会导致程序退出，不推荐随便使用panic
	//recover捕捉到panic
	function_using_test2()
	//1defer需要放到panic之前定义，recover只有在defer调用的函数中才能生效
	//2recover处理异常后，逻辑不会回到panic的那个点去
	//3多个defer会形成栈，先定义的会后执行
}

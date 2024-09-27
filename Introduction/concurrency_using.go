package main

import (
	"fmt"
	"time"
)

// python,java,php多线程编程、多进程编程，多线程和多进程存在的问题主要是耗费内存
// 内存、线程切换 web2.0 用户级线程，绿程，轻量级线程，协程，asyncio-python php-swoole java-netty
// 内存占用小，切换快 go 语言协程goroutine，非常方便
func asyncPrint() {
	fmt.Println("bobby")
}
func main() {
	// go asyncPrint()
	// time.Sleep(time.Second)

	// go func() {
	// 	for {
	// 		time.Sleep(time.Second)
	// 		fmt.Println("slsdlfja")
	// 	}
	// }()
	// time.Sleep(time.Second * 10)

	//每次for循环的时候，i变量会被重用，当我进行到第二轮的for循环的时候，这个i就变了
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()
	// }
	// fmt.Println("main goroutine")
	// time.Sleep(10 * time.Second)

	// for i := 0; i < 100; i++ {
	// 	tmp := i
	// 	go func() {
	// 		fmt.Println(tmp)
	// 	}()
	// }
	// fmt.Println("main goroutine")
	// time.Sleep(10 * time.Second)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}

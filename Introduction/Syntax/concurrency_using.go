package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// 	"time"
// )

// // python,java,php多线程编程、多进程编程，多线程和多进程存在的问题主要是耗费内存
// // 内存、线程切换 web2.0 用户级线程，绿程，轻量级线程，协程，asyncio-python php-swoole java-netty
// // 内存占用小，切换快 go 语言协程goroutine，非常方便
// func asyncPrint() {
// 	fmt.Println("bobby")
// }

// var total int32
// var wg sync.WaitGroup
// var lock sync.Mutex

// func add1() {
// 	defer wg.Done()
// 	for i := 0; i < 100000; i++ {
// 		atomic.AddInt32(&total, 1)
// 		// lock.Lock()
// 		// total += 1
// 		// lock.Unlock()
// 	}
// }
// func sub1() {
// 	defer wg.Done()
// 	for i := 0; i < 100000; i++ {
// 		// lock.Lock()
// 		// total -= 1
// 		// lock.Unlock()
// 		atomic.AddInt32(&total, -1)
// 	}
// }
// func producer(out chan<- int) {
// 	for i := 0; i < 10; i++ {
// 		out <- i * i
// 	}
// 	close(out)
// }
// func consumer(in <-chan int) {
// 	for num := range in {
// 		fmt.Println(num)
// 	}
// }

// var number, letter = make(chan bool), make(chan bool)

// func printNum() {
// 	i := 1
// 	for {
// 		<-number
// 		//怎么做到，等待另一个goroutine来通知我
// 		fmt.Printf("%d%d", i, i+1)
// 		i += 2
// 		letter <- true
// 	}
// }
// func printLetter() {
// 	i := 0
// 	for {
// 		<-letter
// 		//怎么做到，等待另一个goroutine来通知我
// 		if i > 25 {
// 			return
// 		}
// 		fmt.Printf("%c%c", 'A'+i, 'A'+i+1)
// 		i += 2
// 		number <- true
// 	}
// }
// func main() {
// 	// go asyncPrint()
// 	// time.Sleep(time.Second)

// 	// go func() {
// 	// 	for {
// 	// 		time.Sleep(time.Second)
// 	// 		fmt.Println("slsdlfja")
// 	// 	}
// 	// }()
// 	// time.Sleep(time.Second * 10)

// 	//每次for循环的时候，i变量会被重用，当我进行到第二轮的for循环的时候，这个i就变了
// 	// for i := 0; i < 100; i++ {
// 	// 	go func() {
// 	// 		fmt.Println(i)
// 	// 	}()
// 	// }
// 	// fmt.Println("main goroutine")
// 	// time.Sleep(10 * time.Second)

// 	// for i := 0; i < 100; i++ {
// 	// 	tmp := i
// 	// 	go func() {
// 	// 		fmt.Println(tmp)
// 	// 	}()
// 	// }
// 	// fmt.Println("main goroutine")
// 	// time.Sleep(10 * time.Second)

// 	// for i := 0; i < 100; i++ {
// 	// 	go func(i int) {
// 	// 		fmt.Println(i)
// 	// 	}(i)
// 	// }
// 	// fmt.Println("main goroutine")
// 	// time.Sleep(10 * time.Second)

// 	//go语言生成一些线程，GMP

// 	//子goroutine如何通知到主的goroutine自己结束了，主的goroutine如何知道子的goroutine已经结束了
// 	// var wg sync.WaitGroup
// 	//监控多少个goroutine
// 	// wg.Add(100)
// 	// for i := 0; i < 100; i++ {
// 	// 	go func(i int) {
// 	// 		fmt.Println(i)
// 	// 		wg.Done()
// 	// 	}(i)
// 	// }
// 	// //等待
// 	// wg.Wait()
// 	// fmt.Println("all done")

// 	// wg.Add(100)
// 	// for i := 0; i < 100; i++ {
// 	// 	go func(i int) {
// 	// 		defer wg.Done()
// 	// 		fmt.Println(i)
// 	// 	}(i)
// 	// }
// 	// //等待
// 	// wg.Wait()
// 	// fmt.Println("all done")

// 	//waitgroup主要用于goroutine的执行等到，add和done配套使用

// 	//goroutine锁
// 	//锁-竞争
// 	//lock-灵活 atomic包-性能高
// 	//锁复制后就失去了锁的效果
// 	//锁本质是将并行的东西串形化了，使用lock肯定会影响性能
// 	//即使是设计锁，那么也应该尽量的保证并行
// 	// wg.Add(2)
// 	// go add1()
// 	// go sub1()
// 	// wg.Wait()
// 	// fmt.Println(total)

// 	//我们有两组协程，其中一组负责写数据，另一组负责读数据，web系统绝大部分的场景都是读多写少
// 	//虽然有多个goroutine，读之间并发，读写需要串形，写之间串形
// 	//读写锁
// 	// var num int
// 	// var rwlock sync.RWMutex
// 	// var wg sync.WaitGroup
// 	// wg.Add(6)
// 	// //写的goroutine
// 	// go func() {
// 	// 	time.Sleep(time.Second)
// 	// 	defer wg.Done()
// 	// 	rwlock.Lock() //加写锁，写锁会防止读写锁获取
// 	// 	defer rwlock.Unlock()
// 	// 	num = 12
// 	// 	fmt.Println("get write lock")
// 	// 	time.Sleep(time.Second * 5)
// 	// }()
// 	// //读的goroutine
// 	// for i := 0; i < 5; i++ {
// 	// 	go func() {
// 	// 		defer wg.Done()
// 	// 		for {
// 	// 			rwlock.RLock() //读锁
// 	// 			time.Sleep(time.Millisecond * 500)
// 	// 			fmt.Println(num)
// 	// 			rwlock.RUnlock()
// 	// 		}
// 	// 	}()
// 	// }
// 	// wg.Wait()

// 	//不要通过共享内存来通讯，而要通过通讯来实现内存共享
// 	//php python java多线程编程的时候，两个goroutine之间通讯最常用的方式是一个全局也会提供消息队列的机制，python-queue java 消费者和生产者之间的关系
// 	//channel再加上语法堂使用channel更简单
// 	//默认值nil
// 	//channel=0时阻塞
// 	//有缓冲和无缓冲的channel
// 	//0=无缓冲的channel，只有一种用法，！=0是有缓冲的channel
// 	// var msg chan string
// 	// msg = make(chan string, 0)
// 	// go func(msg chan string) {
// 	// 	data := <-msg
// 	// 	fmt.Println(data)
// 	// }(msg)
// 	// msg <- "bobby"
// 	// time.Sleep(time.Second)
// 	//waitGroup如果少了done调用，容易出现dealock，无缓冲的channel也容易出现
// 	//无缓冲的channel适用于通知，b第一时间知道a是否已经完成
// 	//有缓冲的channel适用于生产者消费者之间的通讯
// 	//channel适用场景：1消息传递消息过滤2信号广播3时间订阅和广播4任务分发5结果汇总6并发控制7同步和异步

// 	// var msg chan int
// 	// msg = make(chan int, 2)
// 	// go func(msg chan int) {
// 	// 	for data := range msg {
// 	// 		fmt.Println(data)
// 	// 	}
// 	// }(msg)
// 	// msg <- 1
// 	// msg <- 2
// 	// close(msg) //与其他编程语言有很大区别
// 	// msg<-3//已经关闭的channel不能放值，但可以取值
// 	// time.Sleep(time.Second)

// 	//----------------------------------
// 	//单向channel
// 	// var ch1 chan int
// 	// var ch2 chan<- float64 //单向channel，只能写入
// 	// var ch3 <-chan int     //单向cahnnel，只能读取

// 	// c := make(chan int, 3)
// 	// var send chan<- int = c //send-only
// 	// var read <-chan int = c //recv-only
// 	// send <- 1
// 	// <-read

// 	// c := make(chan int)
// 	// go producer(c)
// 	// go consumer(c)
// 	// time.Sleep(10 * time.Second)

// 	//------------------------------------------
// 	//交替打印channel
// 	//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728%
// 	// go printNum()
// 	// go printLetter()
// 	// number <- true
// 	// time.Sleep(time.Second * 10)

// 	//----------------------------------------
// 	//select与switch相似，但是select的功能与linux的select、poll、epoll
// 	//select主要作用于多个channel
// 	//两个goroutine，某一个执行完成后，我们立马知道
// 	// var lock sync.Mutex
// 	// g1 := make(chan struct{}) //channel多线程安全
// 	// g2 := make(chan struct{})
// 	// go func(c chan struct{}) {
// 	// 	time.Sleep(time.Second)
// 	// 	g1 <- struct{}{}
// 	// }(g1)
// 	// go func(c chan struct{}) {
// 	// 	time.Sleep(time.Second * 2)
// 	// 	g2 <- struct{}{}
// 	// }(g2)
// 	//1.某一个分支就绪了就执行该分支2两个都就绪了，随机的执行一个，防止饥饿，某一个一直拿锁
// 	// for {
// 	// 	select {
// 	// 	case <-g1:
// 	// 		fmt.Println("g1 done")
// 	// 	case <-g2:
// 	// 		fmt.Println("g2 done")
// 	// 	default:
// 	// 		time.Sleep(10 * time.Millisecond)
// 	// 		fmt.Println("default")
// 	// 	}
// 	// }

// 	// tc := time.NewTicker(5 * time.Second)
// 	// for {
// 	// 	select {
// 	// 	case <-g1:
// 	// 		fmt.Println("g1 done")
// 	// 	case <-g2:
// 	// 		fmt.Println("g2 done")
// 	// 	case <-tc.C:
// 	// 		fmt.Println("timeout")
// 	// 		return
// 	// 	}
// 	// }
// 	// fmt.Println("done")

// 	//----------------------------------------context解决goroutine的信息传递----------------------------------------
// 	//渐进式的方法 有一个goroutine监控cpu的信息
// 	// wg.Add(1)
// 	// go func() {
// 	// 	for {
// 	// 		time.Sleep(2 * time.Second)
// 	// 		fmt.Println("cpu的信息")
// 	// 	}
// 	// }()
// 	// wg.Wait()
// 	// fmt.Println("监控完成")

// 	//主动退出监控程序 共享变量？？
// 	// stop := false
// 	// wg.Add(1)
// 	// go func() {
// 	// 	defer wg.Done()
// 	// 	for !stop {
// 	// 		time.Sleep(2 * time.Second)
// 	// 		fmt.Println("cpu的信息")
// 	// 	}
// 	// }()
// 	// time.Sleep(6 * time.Second)
// 	// stop = true
// 	// wg.Wait()
// 	// fmt.Println("监控完成")

// 	//主动退出监控程序 channel？？不优雅
// 	// stop := make(chan struct{})
// 	// wg.Add(1)
// 	// go func() {
// 	// 	defer wg.Done()
// 	// 	for {
// 	// 		select {
// 	// 		case <-stop:
// 	// 			fmt.Println("退出cpu监控")
// 	// 			return
// 	// 		default:
// 	// 			time.Sleep(2 * time.Second)
// 	// 			fmt.Println("cpu的信息")
// 	// 		}
// 	// 	}
// 	// }()
// 	// time.Sleep(6 * time.Second)
// 	// stop <- struct{}{}
// 	// wg.Wait()
// 	// fmt.Println("监控完成")

// 	//主动退出监控程序 channel？？优雅
// 	// stop := make(chan struct{})
// 	// wg.Add(1)
// 	// go func(c chan struct{}) {
// 	// 	defer wg.Done()
// 	// 	for {
// 	// 		select {
// 	// 		case <-c:
// 	// 			fmt.Println("退出cpu监控")
// 	// 			return
// 	// 		default:
// 	// 			time.Sleep(2 * time.Second)
// 	// 			fmt.Println("cpu的信息")
// 	// 		}
// 	// 	}
// 	// }(stop)
// 	// time.Sleep(6 * time.Second)
// 	// stop <- struct{}{}
// 	// wg.Wait()
// 	// fmt.Println("监控完成")

// 	//主动退出监控程序 context
// 	//context包提供了三种函数WithCancel WithTimeout WithValue
// 	//如果你的goroutine函数中，如果希望被控制、超时和传值，但是不希望影响我原来的接口信息，函数第一个参数就是加上一个context
// 	// wg.Add(1)
// 	// ctx, cancel := context.WithCancel(context.Background())
// 	// go func(ctx context.Context) {
// 	// 	defer wg.Done()
// 	// 	for {
// 	// 		select {
// 	// 		case <-ctx.Done():
// 	// 			fmt.Println("退出cpu监控")
// 	// 			return
// 	// 		default:
// 	// 			time.Sleep(2 * time.Second)
// 	// 			fmt.Println("cpu的信息")
// 	// 		}
// 	// 	}
// 	// }(ctx)
// 	// time.Sleep(6 * time.Second)
// 	// cancel()
// 	// wg.Wait()
// 	// fmt.Println("监控完成")

// 	//1主动发起WithCancel
// 	// wg.Add(1)
// 	// ctx, cancel := context.WithCancel(context.Background())
// 	// ctx2, _ := context.WithCancel(ctx)
// 	// go func(ctx context.Context) {
// 	// 	defer wg.Done()
// 	// 	for {
// 	// 		select {
// 	// 		case <-ctx.Done():
// 	// 			fmt.Println("退出cpu监控")
// 	// 			return
// 	// 		default:
// 	// 			time.Sleep(2 * time.Second)
// 	// 			fmt.Println("cpu的信息")
// 	// 		}
// 	// 	}
// 	// }(ctx2)
// 	// time.Sleep(6 * time.Second)
// 	// cancel()
// 	// wg.Wait()
// 	// fmt.Println("监控完成")

// 	//2timeout主动超时
// 	//3WithDeadline在时间点cancel
// 	//4withValue传值
// 	ctx, canel := context.WithCancel(context.Background())
// 	ctxValue := context.WithValue(ctx, "traceid", "gjw12j")
// 	//主动超时
// 	wg.Add(1)
// 	go func(ctx context.Context) {
// 		fmt.Printf("traceid:%s\r\n", ctx.Value("traceid"))
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("退出cpu监控")
// 				return
// 			default:
// 				time.Sleep(2 * time.Second)
// 				fmt.Println("cpu的信息")
// 			}
// 		}
// 	}(ctxValue)
// 	time.Sleep(6 * time.Second)
// 	canel()
// 	wg.Wait()
// 	fmt.Println("监控完成")
// }

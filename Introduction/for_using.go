package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// time.Sleep(2 * time.Second)
		// fmt.Println(i)
	}
	i := 0
	for i < 3 {
		time.Sleep(2 * time.Second)
		fmt.Println(i)
		i++
	}
}

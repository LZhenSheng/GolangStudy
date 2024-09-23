package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func string_api_test1() {
	s := "abc"
	clone := strings.Clone(s)
	fmt.Println(s == clone)
	fmt.Println(unsafe.StringData(s) == unsafe.StringData(clone))
}
func string_api_test2() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
}
func main() {
	string_api_test2()
}

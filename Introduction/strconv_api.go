package main

import (
	"fmt"
	"strconv"
)

func strconv_api_test() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
}
func main() {

	strconv_api_test()
}

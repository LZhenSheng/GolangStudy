package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	username := "李华"
// 	age := 10
// 	address := "北京"
// 	mobile := "132434234"
// 	// var ages []int = []int{1, 2, 3}
// 	//极难维护
// 	fmt.Println("用户名:"+username, ",年龄:"+strconv.Itoa(age)+",地址:"+address, ",电话:"+mobile)
// 	//非常常用，但是性能不高
// 	fmt.Printf("用户名:%s,年龄:%d,地址:%s,电话:%s\r\n", username, age, address, mobile)
// 	//最好用，性能最差
// 	userMsg := fmt.Sprintf("用户名:%T,年龄:%T,地址:%s,电话:%s", username, age, address, mobile)
// 	fmt.Println(userMsg)

// 	//高性能
// 	var builder strings.Builder
// 	builder.WriteString("用户名:")
// 	builder.WriteString(username)
// 	builder.WriteString(",年龄:")
// 	builder.WriteString(strconv.Itoa(age))
// 	builder.WriteString(",地址:")
// 	builder.WriteString(address)
// 	builder.WriteString(",电话:")
// 	builder.WriteString(mobile)
// 	re := builder.String()
// 	fmt.Println(re)
// }

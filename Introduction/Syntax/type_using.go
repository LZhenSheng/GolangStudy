package main

// import (
// 	"fmt"
// 	"strconv"
// )

// type MyInt2 int

// func (mi MyInt2) string() string {
// 	return strconv.Itoa(int(mi))
// }

// type MyPerson struct {
// 	name    string
// 	age     int
// 	address string
// 	height  float32
// }

// //	type MyStudent struct {
// //		p     MyPerson
// //		score float32
// //	}

// //	type MyStudent struct {
// //		address struct {
// //			city string
// //		}
// //		score float32
// //	}
// //
// // 修改结构体的值或者对象很大（拷贝占用很大的内存空间）
// func (p MyPerson) print() {
// 	fmt.Printf("name:%s,age:%d\n", p.name, p.age)
// }
// func main() {
// 	//type 1定义结构体2定义类型别名3定义接口4类型定义
// 	//别名是为了更好的理解代码
// 	type MyInt = int //类型别名
// 	var i MyInt = 9
// 	var j int = 9
// 	fmt.Println(i + j)
// 	fmt.Printf("%T\n", i)
// 	// type MyInt2 int //类型定义,int不能增加方法，但Myint2可以，扩展基础能力
// 	var ii MyInt2 = 2
// 	fmt.Sprintf("%T\n", ii)
// 	fmt.Println(int(ii) + i)
// 	fmt.Println(ii.string())

// 	var a interface{} = "abc"
// 	//断言
// 	switch a.(type) {
// 	case string:
// 		fmt.Println("string")
// 	}
// 	m := a.(string)
// 	fmt.Println(m)

// 	p1 := MyPerson{
// 		"boddy1",
// 		18,
// 		"慕课网",
// 		1.80,
// 	}
// 	p2 := MyPerson{
// 		name: "boddy1",
// 		age:  18,
// 	}
// 	fmt.Println(p1)
// 	fmt.Println(p2)
// 	person2 := []MyPerson{
// 		{
// 			name: "22323",
// 		},
// 		{
// 			name: "34",
// 		},
// 	}
// 	fmt.Println(person2)

// 	var p MyPerson
// 	p.age = 20
// 	fmt.Println(p.height)
// 	//匿名结构体
// 	addres := struct {
// 		province string
// 		city     string
// 		address  string
// 	}{
// 		province: "北京市",
// 		city:     "通州区",
// 		address:  "xxx",
// 	}
// 	fmt.Println(addres)

// 	// s := MyStudent{
// 	// 	MyPerson{
// 	// 		name: "boddy1",
// 	// 		age:  18,
// 	// 	},
// 	// 	19,
// 	// }
// 	// fmt.Println(s.p.age)

// 	// ss := MyStudent{}
// 	// ss.p.name = "dlkfj"
// 	// ss := MyStudent{}
// 	// ss.address.city = "sdkjflkj"
// 	// fmt.Println(ss)
// 	p2.print()
// }

package main

import (
	"fmt"
)

// 集合类型包括数组、切片、map、list
func collection_type_test1() {
	//course1和course2是不同的类型
	var course1 [3]string
	var course2 [4]string
	fmt.Println("%T\n", course1)
	fmt.Println("%T\n", course2)

	courses3 := [3]string{2: "gin"}
	for _, value := range courses3 {
		fmt.Println(value)
	}

	courses4 := [...]string{"GO", "JAVA", "S"}
	for _, value := range courses4 {
		fmt.Println(value)
	}

	//多维数组
	var courseInfo [3][4]string
	courseInfo[0] = [4]string{"go", "1h", "bobbby", "go课程"}
	courseInfo[1] = [4]string{"grpc", "2h", "bobbby2", "grpc课程"}
	courseInfo[2] = [4]string{"gin", "1.5h", "bobby3", "gin课程"}
	for i := 0; i < len(courseInfo); i++ {
		for j := 0; j < len(courseInfo[i]); j++ {
			fmt.Print(courseInfo[i][j], "\t")
		}
		fmt.Println()
	}

	for _, row := range courseInfo {
		fmt.Println(row)
		// for _, column := range row {
		// 	fmt.Print(column, "\t")
		// }
		// fmt.Println()
	}
	//数组相等需要对应元素相等（位置+元素）

}
func collection_type_test2() {
	//go 弱化数组功能，提供切片实现动态数组功能
	var course []string
	fmt.Printf("%T\n", course)

	//这个方法很特别，第一次用很头疼
	course = append(course, "go")
	fmt.Println(course)

	//切片初始化三种方法:1.数组直接创建,2.使用string{},3.make
	allCourses := []string{"go", "grpc", "gin", "mysql"}
	courseSlice := allCourses[0:1]
	allCourses = append(allCourses, "dlskjf", "dlskjf", "dslkjfsdf", "dlsjfk")
	courseSlice = append(courseSlice, allCourses...)
	fmt.Println(allCourses)
	fmt.Println(courseSlice)

	//make函数
	allCourses2 := make([]string, 3)
	allCourses2[0] = "c"
	fmt.Println(allCourses2)

	//切片删除元素
	courseSlice2 := []string{"go", "grpc", "mysql", "es", "gin"}
	myslice := append(courseSlice2[:2], courseSlice2[3:]...)
	fmt.Println(myslice)

	//复制slice
	// courseSliceCopy := courseSlice
	// courseSliceCopy2 := courseSlice[:]
	// fmt.Println(courseSliceCopy)
	// fmt.Println(courseSliceCopy2)

	var courseSliceCopy = make([]string, len(courseSlice2))
	copy(courseSliceCopy, courseSlice2)
	courseSliceCopy[0] = "java"
	fmt.Println(courseSliceCopy)
	fmt.Println(courseSlice2)

}

//	type slice struct {
//		array unsafe.Pointer //用来存储实际数据的数组指针，指向一块连续的内存
//		len   int            //切片中元素的数量
//		cap   int            //array数组的长度
//	}
func collection_type_test3() {
	// data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// s1 := data[1:6]
	// s2 := data[2:7]
	// s2[0] = 22
	// fmt.Println(s1)
	// fmt.Println(s2)

	// s2 = append(s2, 1, 2, 3, 3, 44, 34, 34, 23, 23, 234, 34, 2, 3, 23, 32, 23, 32)
	// s2[0] = 44

	// fmt.Println(s1)
	// fmt.Println(s2)

	//512之前双倍阔容，512之后就不翻倍阔了
	var data []int
	for i := 0; i < 513; i++ {
		data = append(data, i)
		fmt.Println("len: %d, cap: %d\r\n", len(data), cap(data))
	}
}
func collection_type_test() {
	//map是一个key和value的无序集合，主要是查询方便o(1)
	var courseMap = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
	}

	//取值
	fmt.Println(courseMap["grpc"])
	//放值
	courseMap["mysql"] = "mysql索引"
	fmt.Println(courseMap)

	//map必须初始化之后使用
	var courseM = map[string]string{}
	var courseMM = make(map[string]string, 3)
	fmt.Println(courseM)
	fmt.Println(courseMM)

	var m []string
	if m == nil {
		fmt.Println("yes")
	}
	m = append(m, "dslkjf")
}
func main() {
	collection_type_test()
}

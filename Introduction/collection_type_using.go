package main

import "fmt"

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
func collection_type_test() {

}
func main() {
	collection_type_test()
}

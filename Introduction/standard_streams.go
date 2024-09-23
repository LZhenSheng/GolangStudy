package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// general
func test1() {
	arr := [3]int{1, -23, 3}
	sli := []int{1, -23, 3}
	student := &Student{
		Name: "sdlkjf",
		Age:  23,
	}
	fmt.Printf("%v", arr)
	fmt.Printf("%v", sli)
	fmt.Printf("%v\r\n", student)

	fmt.Printf("%+v", arr)
	fmt.Printf("%+v", sli)
	fmt.Printf("%+v\r\n", student)

	fmt.Printf("%#v", arr)
	fmt.Printf("%#v", sli)
	fmt.Printf("%#v\r\n", student)

	fmt.Printf("%#T", arr)
	fmt.Printf("%#T", sli)
	fmt.Printf("%#T\r\n", student)

	fmt.Printf("%%", arr)
	fmt.Printf("%%", sli)
	fmt.Printf("%%", student)
	fmt.Printf("\n")
}

// bool
func test2() {
	fmt.Printf("%t", true)
	fmt.Printf("%t", false)
}
func test3() {
	var num1 int = 5
	var num2 int8 = 6
	var num3 int16 = 7
	var num4 int32 = 8
	var num5 int64 = -9
	var num6 uint = 10
	var num7 uint8 = 11
	var num8 uint16 = 12
	var num9 uint32 = 13
	var num10 uint64 = 14
	fmt.Printf("%d %d %d %d %d %d %d %d %d %d\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("%+d %+d %+d %+d %+d %+d %+d %+d %+d %+d\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("-%4d %4d %4d %4d %4d %4d %4d %4d %4d %4d-\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("-%-4d %-4d %-4d %-4d %-4d %-4d %-4d %-4d %-4d %-4d-\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("%b %b %b %b %b %b %b %b %b %b\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("%o %o %o %o %o %o %o %o %o %o\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("%O %O %O %O %O %O %O %O %O %O\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("%x %x %x %x %x %x %x %x %x %x\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("%X %X %X %X %X %X %X %X %X %X\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("%U %U %U %U %U %U %U %U %U %U\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	fmt.Printf("%q %q %q %q %q %q %q %q %q %q\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
	// fmt.Printf("%c %c %c %c %c %c %c %c %c\n", num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
}
func test4() {
	var num1 float32 = 2.23
	var num2 float64 = 5.78
	fmt.Printf("%b %b\n", num1, num2)
	fmt.Printf("%e %e\n", num1, num2)
	fmt.Printf("%E %E\n", num1, num2)
	fmt.Printf("%f %f\n", num1, num2)
	fmt.Printf("%F %F\n", num1, num2)
	fmt.Printf("%g %g\n", num1, num2)
	fmt.Printf("%e %e\n", num1, num2)
	fmt.Printf("%G %G\n", num1, num2)
	fmt.Printf("%E %E\n", num1, num2)
	fmt.Printf("%x %x\n", num1, num2)
	fmt.Printf("%X %X\n", num1, num2)

}
func test5() {
	var str1 string = "sdlkfjasdfja"
	var str2 string = "HelloWord"
	fmt.Printf("%s %s\n", str1, str2)
	fmt.Printf("%q %q\n", str1, str2)
	fmt.Printf("%x %x\n", str1, str2)
	fmt.Printf("%X %X\n", str1, str2)

}
func main() {
	//test1()
	//test2()
	// test3()
	// test4()
	test5()
}

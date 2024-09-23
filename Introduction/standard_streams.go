package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

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
	fmt.Printf("%.4f %.4f\n", num1, num2)
	fmt.Printf("-%9f %9f-\n", num1, num2)
	fmt.Printf("-%9.f %9.f-\n", num1, num2)
	fmt.Printf("-%9.4f %9.4f-\n", num1, num2)

	// fmt.Printf("%F %F\n", num1, num2)
	// fmt.Printf("%g %g\n", num1, num2)
	// fmt.Printf("%e %e\n", num1, num2)
	// fmt.Printf("%G %G\n", num1, num2)
	// fmt.Printf("%E %E\n", num1, num2)
	// fmt.Printf("%x %x\n", num1, num2)
	// fmt.Printf("%X %X\n", num1, num2)

}
func test5() {
	var str1 string = "sdlkfjasdfja"
	var str2 string = "HelloWord"
	fmt.Printf("%s %s\n", str1, str2)
	fmt.Printf("%q %q\n", str1, str2)
	fmt.Printf("%x %x\n", str1, str2)
	fmt.Printf("%X %X\n", str1, str2)
	fmt.Printf("%p %p\n", &str1, &str2)
	fmt.Printf("%b %b\n", &str1, &str2)
	fmt.Printf("%d %d\n", &str1, &str2)
	fmt.Printf("%o %o\n", &str1, &str2)
	fmt.Printf("%x %x\n", &str1, &str2)
	fmt.Printf("%X %X\n", &str1, &str2)

}
func test6() {
	fmt.Printf("%[2]d %[1]d\n", 11, 22)
	fmt.Printf("%[3]*.[2]*[1]f\n", 12.0, 2, 6)
	fmt.Printf("%6.2f\n", 12.0)
}
func test7() {
	//类型错误或动词未知：%!verb(type=value)
	fmt.Printf("%d", "hi")
	fmt.Println()
	//%!d(string=hi)

	//参数太多：%!(EXTRA type=value)
	fmt.Printf("hi", "guys")
	fmt.Println()
	//hi%!(EXTRA string=guys)

	//参数太少：%!verb(MISSING)
	fmt.Printf("hi%d\n")
	fmt.Println()
	//hi%!d(MISSING)

	//宽度或精度非整数：%!(BADWIDTH) 或 %!(BADWIDTH)
	fmt.Printf("%*s", 4.5, "hi")
	fmt.Println()
	//%!(BADWIDTH)hi
	fmt.Printf("%.*s", 4.5, "hi")
	fmt.Println()
	//%!(BADPREC)hi

	//参数索引无效或使用不正确：%!(BADINDEX)
	fmt.Printf("%*[2]d", 7)
	fmt.Println()
	//%!d(BADINDEX)
	fmt.Printf("%.[2]d", 7)
	fmt.Println()
	//%!d(BADINDEX)
}
func test8() {
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")

}
func test9() {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
}
func test10() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")

}
func test11() {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
}
func test12() {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscan(r, &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
}
func test13() {
	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}
}
func test14() {
	const name, age = "Kim", 22
	fmt.Print(name, " is ", age, " years old.\n")
}
func main() {
	//test1()
	//test2()
	// test3()
	// test4()
	// test5()
	// test6()
	//test7()
	//test8()
	// test9()
	// test10()
	//test12()
	// test13()
	test14()
}

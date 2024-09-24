package main

import (
	"fmt"
	"unicode"
)

func unicode_api_test1() {
	fmt.Printf("%t\n", unicode.IsDigit('৩'))
	fmt.Printf("%t\n", unicode.IsDigit('A'))
}
func unicode_api_test2() {
	fmt.Printf("%t\n", unicode.IsLetter('A'))
	fmt.Printf("%t\n", unicode.IsLetter('7'))
}
func unicode_api_test3() {
	fmt.Printf("%t\n", unicode.IsLower('a'))
	fmt.Printf("%t\n", unicode.IsLower('A'))
}
func unicode_api_test4() {
	fmt.Printf("%t\n", unicode.IsNumber('Ⅷ'))
	fmt.Printf("%t\n", unicode.IsNumber('A'))
}
func unicode_api_test5() {
	fmt.Printf("%t\n", unicode.IsSpace(' '))
	fmt.Printf("%t\n", unicode.IsSpace('\n'))
	fmt.Printf("%t\n", unicode.IsSpace('\t'))
	fmt.Printf("%t\n", unicode.IsSpace('a'))
}
func unicode_api_test6() {
	fmt.Printf("%t\n", unicode.IsTitle('ǅ'))
	fmt.Printf("%t\n", unicode.IsTitle('a'))
}
func unicode_api_test7() {
	fmt.Printf("%#U\n", unicode.SimpleFold('A'))      // 'a'
	fmt.Printf("%#U\n", unicode.SimpleFold('a'))      // 'A'
	fmt.Printf("%#U\n", unicode.SimpleFold('K'))      // 'k'
	fmt.Printf("%#U\n", unicode.SimpleFold('k'))      // '\u212A' (Kelvin symbol, K)
	fmt.Printf("%#U\n", unicode.SimpleFold('\u212A')) // 'K'
	fmt.Printf("%#U\n", unicode.SimpleFold('1'))      // '1'
}
func unicode_api_test8() {
	const lcG = 'g'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, lcG))

	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, ucG))
}
func unicode_api_test9() {
	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.ToLower(ucG))
}
func unicode_api_test10() {
	const ucG = 'g'
	fmt.Printf("%#U\n", unicode.ToTitle(ucG))
}
func unicode_api_test() {
	const ucG = 'g'
	fmt.Printf("%#U\n", unicode.ToUpper(ucG))
}
func main() {
	unicode_api_test()
}

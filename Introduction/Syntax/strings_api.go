package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// 	"unsafe"
// )

// func string_api_test1() {
// 	s := "abc"
// 	clone := strings.Clone(s)
// 	fmt.Println(s == clone)
// 	fmt.Println(unsafe.StringData(s) == unsafe.StringData(clone))
// }
// func string_api_test2() {
// 	fmt.Println(strings.Compare("a", "b"))
// 	fmt.Println(strings.Compare("a", "a"))
// 	fmt.Println(strings.Compare("b", "a"))
// }
// func string_api_test3() {
// 	fmt.Println(strings.Contains("seafood", "foo"))
// 	fmt.Println(strings.Contains("seafood", "bar"))
// 	fmt.Println(strings.Contains("seafood", ""))
// 	fmt.Println(strings.Contains("", ""))
// }
// func string_api_test4() {
// 	fmt.Println(strings.ContainsAny("team", "i"))
// 	fmt.Println(strings.ContainsAny("fail", "ui"))
// 	fmt.Println(strings.ContainsAny("ure", "ui"))
// 	fmt.Println(strings.ContainsAny("failure", "ui"))
// 	fmt.Println(strings.ContainsAny("foo", ""))
// 	fmt.Println(strings.ContainsAny("", ""))
// }
// func string_api_test5() {
// 	f := func(r rune) bool {
// 		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
// 	}
// 	fmt.Println(strings.ContainsFunc("hello", f))
// 	fmt.Println(strings.ContainsFunc("rhythms", f))
// }
// func string_api_test6() {
// 	// Finds whether a string contains a particular Unicode code point.
// 	// The code point for the lowercase letter "a", for example, is 97.
// 	fmt.Println(strings.ContainsRune("aardvark", 97))
// 	fmt.Println(strings.ContainsRune("timeout", 97))
// }
// func string_api_test7() {
// 	fmt.Println(strings.Count("cheese", "e"))
// 	fmt.Println(strings.Count("five", "")) // before & after each rune
// }
// func string_api_test8() {
// 	show := func(s, sep string) {
// 		before, after, found := strings.Cut(s, sep)
// 		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
// 	}
// 	show("Gopher", "Go")
// 	show("Gopher", "ph")
// 	show("Gopher", "er")
// 	show("Gopher", "Badger")
// }
// func string_api_test9() {
// 	show := func(s, sep string) {
// 		after, found := strings.CutPrefix(s, sep)
// 		fmt.Printf("CutPrefix(%q, %q) = %q, %v\n", s, sep, after, found)
// 	}
// 	show("Gopher", "Go")
// 	show("Gopher", "ph")
// }
// func string_api_test10() {
// 	show := func(s, sep string) {
// 		before, found := strings.CutSuffix(s, sep)
// 		fmt.Printf("CutSuffix(%q, %q) = %q, %v\n", s, sep, before, found)
// 	}
// 	show("Gopher", "Go")
// 	show("Gopher", "er")
// }
// func string_api_test11() {
// 	fmt.Println(strings.EqualFold("Go", "go"))
// 	fmt.Println(strings.EqualFold("AB", "ab")) // true because comparison uses simple case-folding
// 	fmt.Println(strings.EqualFold("ß", "ss"))  // false because comparison does not use full case-folding
// }
// func string_api_test12() {
// 	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

// }
// func string_api_test13() {
// 	f := func(c rune) bool {
// 		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
// 	}
// 	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
// }
// func string_api_test14() {
// 	fmt.Println(strings.HasPrefix("Gopher", "Go"))
// 	fmt.Println(strings.HasPrefix("Gopher", "C"))
// 	fmt.Println(strings.HasPrefix("Gopher", ""))
// }
// func string_api_test15() {
// 	fmt.Println(strings.HasSuffix("Amigo", "go"))
// 	fmt.Println(strings.HasSuffix("Amigo", "O"))
// 	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
// 	fmt.Println(strings.HasSuffix("Amigo", ""))
// }
// func string_api_test16() {
// 	fmt.Println(strings.Index("chicken", "ken"))
// 	fmt.Println(strings.Index("chicken", "dmr"))
// 	fmt.Println(strings.Index("我爱go", "爱"))
// }
// func string_api_test17() {
// 	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
// 	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
// 	fmt.Println(strings.IndexAny("我爱go", "爱"))
// }
// func string_api_test18() {
// 	fmt.Println(strings.IndexByte("golang", 'g'))
// 	fmt.Println(strings.IndexByte("gophers", 'h'))
// 	fmt.Println(strings.IndexByte("golang", 'x'))
// }
// func string_api_test19() {
// 	f := func(c rune) bool {
// 		return unicode.Is(unicode.Han, c)
// 	}
// 	fmt.Println(strings.IndexFunc("Hello, 世界", f))
// 	fmt.Println(strings.IndexFunc("Hello, world", f))
// }
// func string_api_test20() {
// 	fmt.Println(strings.IndexRune("chicken", 'k'))
// 	fmt.Println(strings.IndexRune("chicken", 'd'))
// }
// func string_api_test21() {
// 	s := []string{"foo", "bar", "baz"}
// 	fmt.Println(strings.Join(s, ", "))
// }
// func string_api_test22() {
// 	fmt.Println(strings.Index("go gopher", "go"))
// 	fmt.Println(strings.LastIndex("go gopher", "go"))
// 	fmt.Println(strings.LastIndex("go gopher", "rodent"))
// }
// func string_api_test23() {
// 	fmt.Println(strings.LastIndexAny("go gopher", "go"))
// 	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
// 	fmt.Println(strings.LastIndexAny("go gopher", "fail"))
// }
// func string_api_test24() {
// 	fmt.Println(strings.LastIndexByte("Hello, world", 'l'))
// 	fmt.Println(strings.LastIndexByte("Hello, world", 'o'))
// 	fmt.Println(strings.LastIndexByte("Hello, world", 'x'))
// }
// func string_api_test25() {
// 	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
// 	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
// 	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))
// }
// func string_api_test26() {
// 	rot13 := func(r rune) rune {
// 		switch {
// 		case r >= 'A' && r <= 'Z':
// 			return 'A' + (r-'A'+13)%26
// 		case r >= 'a' && r <= 'z':
// 			return 'a' + (r-'a'+13)%26
// 		}
// 		return r
// 	}
// 	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
// }
// func string_api_test27() {
// 	fmt.Println("ba" + strings.Repeat("na", 2))
// }
// func string_api_test28() {
// 	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
// 	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
// }
// func string_api_test29() {
// 	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
// 	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
// 	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
// 	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
// }
// func string_api_test30() {
// 	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
// }
// func string_api_test31() {
// 	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
// }
// func string_api_test32() {
// 	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
// 	z := strings.SplitN("a,b,c", ",", 0)
// 	fmt.Printf("%q (nil = %v)\n", z, z == nil)
// }
// func string_api_test33() {
// 	fmt.Println(strings.ToLower("Gopher"))
// }
// func string_api_test34() {
// 	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
// }
// func string_api_test35() {
// 	fmt.Println(strings.ToTitle("her royal highness"))
// 	fmt.Println(strings.ToTitle("loud noises"))
// 	fmt.Println(strings.ToTitle("хлеб"))
// }
// func string_api_test36() {
// 	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
// }
// func string_api_test37() {
// 	fmt.Println(strings.ToUpper("Gopher"))
// }
// func string_api_test38() {
// 	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
// }
// func string_api_test39() {
// 	fmt.Printf("%s\n", strings.ToValidUTF8("abc", "\uFFFD"))
// 	fmt.Printf("%s\n", strings.ToValidUTF8("a\xffb\xC0\xAFc\xff", ""))
// 	fmt.Printf("%s\n", strings.ToValidUTF8("\xed\xa0\x80", "abc"))
// }
// func string_api_test40() {
// 	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
// }
// func string_api_test41() {
// 	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
// 		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
// 	}))
// }
// func string_api_test42() {
// 	fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
// }
// func string_api_test43() {
// 	fmt.Print(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
// 		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
// 	}))
// }
// func string_api_test44() {
// 	var s = "¡¡¡Hello, Gophers!!!"
// 	s = strings.TrimPrefix(s, "¡¡¡Hello, ")
// 	s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
// 	fmt.Print(s)
// }
// func string_api_test45() {
// 	fmt.Print(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
// }
// func string_api_test46() {
// 	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
// }
// func string_api_test47() {
// 	var s = "¡¡¡Hello, Gophers!!!"
// 	s = strings.TrimSuffix(s, ", Gophers!!!")
// 	s = strings.TrimSuffix(s, ", Marmots!!!")
// 	fmt.Print(s)
// }
// func string_api_test48() {
// 	var b strings.Builder
// 	for i := 3; i >= 1; i-- {
// 		fmt.Fprintf(&b, "%d...", i)
// 	}
// 	b.WriteString("ignition")
// 	fmt.Println(b.String())
// }
// func string_api_test() {
// 	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
// 	fmt.Println(r.Replace("This is <b>HTML</b>!"))
// }
// func main() {
// 	string_api_test()
// }

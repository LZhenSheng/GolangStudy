package main

import "fmt"

func main() {
	day := "星期二"
	switch day {
	case "星期一":
		fmt.Println("Monday")
	case "星期二":
		fmt.Println("Tuesday")
	default:
		fmt.Println("UnKnown")
	}

	switch {
	case day == "星期一":
		fmt.Println("Monday")
	case day == "星期二":
		fmt.Println("Tuesday")
	default:
		fmt.Println("UnKnown")
	}
}

package main

import "fmt"

type Duck interface {
	//方法的深情
	Gaga()
	Walk()
	Swiming()
}
type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎")
}
func (pd *pskDuck) Walk() {
	fmt.Println("this is pskduck walking")
}
func (pd *pskDuck) Swiming() {
	fmt.Println("swiming")
}
func main() {
	var d Duck = &pskDuck{}
	d.Gaga()
}

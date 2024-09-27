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

type MyWriter interface {
	Write(string) error
}
type MyClose interface {
	Close() error
}
type writerCloser struct{}

func (wc *writerCloser) Write(string) error {
	fmt.Println("write string")
	return nil
}
func (wc *writerCloser) Close() error {
	fmt.Println("close")
	return nil
}
func main() {
	var d Duck = &pskDuck{}
	d.Gaga()

	var my MyWriter = &writerCloser{}
	var mc myCloser = &writerCloser{}

}

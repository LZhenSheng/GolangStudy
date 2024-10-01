package main

// import "fmt"

// type Duck interface {
// 	//方法的深情
// 	Gaga()
// 	Walk()
// 	Swiming()
// }
// type pskDuck struct {
// 	legs int
// }

// func (pd *pskDuck) Gaga() {
// 	fmt.Println("嘎嘎")
// }
// func (pd *pskDuck) Walk() {
// 	fmt.Println("this is pskduck walking")
// }
// func (pd *pskDuck) Swiming() {
// 	fmt.Println("swiming")
// }

// type MyWriter interface {
// 	Write(string) error
// }
// type MyClose interface {
// 	Close() error
// }
// type writerCloser struct {
// 	MyWriter //interface也是一个类型，想放入一个写文件的实现，我想放入一个写数据库的实现
// }
// type fileWriter struct {
// 	filePath string
// }
// type databaseWriter struct {
// 	host string
// 	port string
// }

// func (fw *fileWriter) Write(string) error {
// 	fmt.Println("write string to file")
// 	return nil
// }

// func (dw *databaseWriter) Write(string) error {
// 	fmt.Println("write string to databaset")
// 	return nil
// }

// //	func (wc *writerCloser) Write(string) error {
// //		fmt.Println("write string")
// //		return nil
// //	}

// func interface_add(a, b interface{}) interface{} {
// 	switch a.(type) {
// 	case int:
// 		ai, _ := a.(int)
// 		bi, _ := b.(int)
// 		return ai + bi
// 	case int32:
// 		ai, _ := a.(int32)
// 		bi, _ := b.(int32)
// 		return ai + bi
// 	default:
// 		panic("not support")
// 	}
// }
// func (wc *writerCloser) Close() error {
// 	fmt.Println("close")
// 	return nil
// }

// // 接口的重复使用
// type _MyWriter interface {
// 	Write(string)
// }
// type _MyReader interface {
// 	Read() string
// }
// type MyReadWriter interface {
// 	_MyWriter
// 	_MyReader
// 	MyReadWrite()
// }
// type SreadWriter struct{}

// func (srw *SreadWriter) Read() string {
// 	fmt.Println("read")
// 	return "read"
// }
// func (srw *SreadWriter) Write(str string) {
// 	fmt.Println(str)
// }
// func (srw *SreadWriter) MyReadWrite() {
// 	fmt.Println("MyReadWrite")
// }
// func mPrintln(data ...interface{}) {
// 	for _, value := range data {
// 		fmt.Println(value)
// 	}
// }

// type myinfo struct{}

// func (mi *myinfo) Error() string {
// 	return "我不是error"
// }
// func main() {
// 	var d Duck = &pskDuck{}
// 	d.Gaga()

// 	var my MyClose = &writerCloser{}
// 	var mc MyWriter = &writerCloser{
// 		&fileWriter{},
// 	}
// 	fmt.Println(my.Close())
// 	mc.Write("bobby")
// 	var a int32 = 2
// 	var b int32 = 2
// 	fmt.Println(interface_add(a, b))
// 	re := interface_add(a, b)
// 	ree, _ := re.(int32)
// 	fmt.Println(ree)

// 	var mrw MyReadWriter = &SreadWriter{}
// 	mrw.Read()

// 	//接口遇到了slice
// 	var data = []interface{}{
// 		"bobby",
// 		18,
// 		1.80,
// 	}
// 	mPrintln(data...) //不能使用string类型的slice，但可以将string转换成slice
// 	//实现了Error的就是Error
// 	var err error = &myinfo{}
// 	fmt.Println(err)
// }

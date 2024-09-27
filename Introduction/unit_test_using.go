package main

func main() {
	//单元测试go test
	//go test命令式一个按照一定约定和组织的测试代码驱动程序，在包中，所有以_test.go为后缀的源码文件都会被go test运行到
	//我们写的_test.go源文件不会担心内容过多，因为go build不会讲这些测试文件打包到最后的可执行文件中
	//test文件有四类
	//1Test开头的功能测试 2Benchmark开头的性能测试 3example 4模糊测试

	//go test
	//go test .
	//跳过一些测试
	//以时间短的测试运行
	//go test -short
	// func TestAdd2(t *testing.t){
	// 	if testing.Short(){
	// 		t.skip("short模式下跳过")
	// 	}
	// 	fmt.Println("dlsdkjaflsdfl")
	// }

	//表格驱动测试用例
	// func TestAdd(t *testing.T){
	// 	var dataset=[]struct{
	// 		a int
	// 		b int
	// 		out int
	// 	}{
	// 		{1,1,2},
	// 		{12,12,24},
	// 	}
	// 	for _,value:=range dataset{
	// 		re:=add(value.a,value.b)
	// 		if re!=value.out{
	// 			t.Errorf("expecet:%d,actual:%d",value.out,re)
	// 		}
	// 	}
	// }

	//性能测试测试核心函数
	//go test -bench=".*"
	// func BenchmarAdd(b *testing.B){
	// 	var a,b,c int
	// 	a=123
	// 	b=456
	// 	c=789
	// 	for i:=0;i<bb.N;i++{
	// 		if actual:=add(a,b);actual!=c{
	// 			fmt.Println("%d + %d,expect:%d",a,b,actual)
	// 		}
	// 	}
	// }

	// func BenchmarkStringBuilder(b *testing.B){
	// 	b.ResetTimer()

	// 	for i:=0;i<b.N;i++{
	// 		var str string
	// 		var builder strings.Builder
	// 		for j:=0;j<1000;j++{
	// 			builder.WriteString(strconv.Itoa(j))
	// 		}
	// 		_=builder.String()
	// 	}
	// 	b.ResetTimer()
	// }
}

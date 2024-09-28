package main

import "fmt"

type Person3 struct {
	name string
}

func changeName(p *Person3) {
	p.name = "imooc"
}
func swap(a, b *int) {
	a, b = b, a
}
func main() {
	p := Person3{
		name: "boddy",
	}
	changeName(&p)
	fmt.Println(p)

	//指针的定义
	var po *Person3
	po = &p
	fmt.Println(*po)
	//第一个点出来
	//第二个限制了指针的运算，不能进行运算(不能+1)
	//指针在unsafe包，一般不使用unsafe包，但也可以使用
	fmt.Println(po.name)
	fmt.Println((*po).name)

	//指针需要初始化
	var b *int
	fmt.Println(b)

	//不能对nil的变量访问子变量
	//第一种初始化
	ps := &Person3{}
	fmt.Println(ps.name)
	//第二种初始化
	var emptyPerson Person3
	pi := &emptyPerson
	//第三种初始化
	var pp *Person3 = new(Person3)
	fmt.Println(pi.name)
	fmt.Println(pp)
	//初始化两个关键字,map,channel,slice初始化使用make方法，指针使用new方法
	//map必须初始化

}

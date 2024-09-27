package main

import "GolangStudy/temp"

func main() {
	//package用来组织源码，是多哥go源码的集合，代码复用的基础
	temp.Couse()

	//1通过文件名调用
	//import "GolangStudy/temp"
	//2通过别名调用 避免不同包相同名的问题
	//import u "GolangStudy/temp"
	//3全部导入 尽量少用
	//import . "GolangStudy/temp"
	//4匿名导入 常用语初始化，自动调用init
	//import _ "GolangStudy/temp"

	//设置国内代理
	//go env -w GO111MODULE=on
	//go env -w GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,direct
	//go module 添加以依赖 删除未使用的依赖项
	//查看依赖项
	// go list -m all
	//查看可用版本
	// go list -m -versions github.com/gin-gonic/gin
	//使用特定版本的gin
	// go list -m -versions github.com/gin-gonic/gin@v1.8.0

	//查看帮助
	//go mod help

	//查看依赖关系
	//go mod graph

	//初始化mod文件
	//go mod init

	//整理依赖
	//go mod tidy

	//安装
	//go install

	//升级到最新的次要版本或修订版本
	//go get -u
	//升级到最新的修订版本
	//go get -u =patch

	//替换
	//go mode edit -replace github.com/bobby/A=github.com/bobby/project-A@v1.0.0

	//1.代码规范
	//（1）代码规范不是强制的，单是不同的语言一些细微的规范还是要遵守的
	//（2）代码规范主要是为方便团队内部形成一个统一的代码风格，提高代码的可读性，统一性
	//1.1命名规范
	//1.1.1包名。尽量和目录保持一致 尽量采用有意义的包名，简短，包名不要和标准库重名，包名采用全部小写
	//1.1.2文件名。有多个单词可以采用蛇形命名法
	//1.1.3变量名 蛇形python php 驼峰java c go

}

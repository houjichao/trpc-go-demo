package main

import (
	"fmt"
	"github.com/houjichao/trpc-go-demo/test/pkg/crm/calcutils" //对应文件夹的名字
	"github.com/houjichao/trpc-go-demo/test/pkg/crm/dbutils"
	test "github.com/houjichao/trpc-go-demo/test/pkg/crm/strutils"
)

/*
	这里是对go中包的概念进行学习
	注意细节或者建议：
	1.package 进行包的声明，建议：包的声明和这个包所在的文件夹同名
	2.main包是程序的入口包，一般main函数会放在这个包下
	3.打包语法：
	package 包名
	4.引入包的语法
	import 包路径
	5.如果有多个包，建议一次性导入
	import (
		"fmt"
		"github.com/houjichao/trpc-go-demo/test/pkg/crm/dbutils"
	)
	6.在函数调用的前面要定位到所在的包
	dbutils.GetConn()
	7.函数命名首字母大写，可以被其他包访问
	8.一个目录下不能有重复的函数，go中函数不能被重载
	9.包名和文件夹的名字，可以不一样
	案例：calcutils
	10.一个目录下的同级文件，归属一个包
	同级别的源文件包的声明必须一致
	11.包到底是什么
	（1）在程序层面，所有使用相同package 包名的源文件组成的代码模块
	（2）在源文件层面就是一个文件夹

 */
func main()  {
	fmt.Println("main函数入口进来了")
	dbutils.GetConn()
	//包名和文件夹的名字，可以不一样 -- 建议不这么做
	aaa.AddNum() //对应包名

	//可以对包取别名，但是取了别名后，就不能用原包名进行调用了
	test.SplitStr()
}
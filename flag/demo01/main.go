package main

import (
	"flag"
	"fmt"
)

/*
func String
func String(name string, value string, usage string) *string
String用指定的名称、默认值、使用信息注册一个string类型flag。返回一个保存了该flag的值的指针。
在执行时使用 --name=value进行传递

命令行flag语法：
-flag
-flag=x
-flag x  // 只有非bool类型的flag可以 bool类型使用=赋值
一个-和两个-是一样的

例如：
go run main.go --mode=hello
*/

//第一个参数为name 第二个参数为默认值 第三个参数是描述
//需要注意的是接受的变量是一个指针

var mode = flag.String("mode", "str1", "process mode")

func main() {

	//进行解析
	flag.Parse()

	//输出
	fmt.Println("---" + *mode)

	var str string
	fmt.Println(str)
}

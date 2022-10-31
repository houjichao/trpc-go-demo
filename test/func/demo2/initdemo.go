package main

import (
	"fmt"
	"github.com/houjichao/trpc-go-demo/test/func/demo2/numutils"
	"github.com/houjichao/trpc-go-demo/test/func/demo2/testutils"
)

/*
	init 函数学习

	结论：
	同一个源文件中，执行顺序为 全局变量定义 --> init函数 --> main函数
	多个源文件都有init函数，执行顺序依次为根据import顺序依次执行
	同一个文件中的执行顺序和上相同
*/

var num int = assignment()

func assignment() int {
	fmt.Println("assignment函数被执行了")
	return 10
}

func init() {
	fmt.Println("main中的init函数被执行了")
}
func main() {
	fmt.Println("main函数被执行了")
	fmt.Printf("main中的num为%v\n", num)
	fmt.Printf("Age = %v,Name = %v\n", testutils.Age, testutils.Name)
	fmt.Printf("num = %v\n", numutils.Num)
}

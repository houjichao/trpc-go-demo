package main

import (
	"fmt"
	"time"
)

/*
	go关键字也可以为匿名函数或闭包启动goroutine
	使用匿名函数或闭包创建 goroutine 时，除了将函数定义部分写在 go 的后面之外，还需要加上匿名函数的调用参数，格式如下：

	go func( 参数列表 ){
		函数体
	}( 调用参数列表 )
	其中：
	参数列表：函数体内的参数变量列表。
	函数体：匿名函数的代码。
	调用参数列表：启动 goroutine 时，需要向匿名函数传递的调用参数
*/
func main() {
	go func(times int) {
		for {
			times++
			fmt.Println("tick", times)

			//延迟2s
			time.Sleep(2 * time.Second)

		}
	}(0)

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

	//所有 goroutine 在 main() 函数结束时会一同结束。
}

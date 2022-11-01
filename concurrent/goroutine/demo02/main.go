package main

import (
	"fmt"
	"time"
)

//使用普通函数创建goroutine
//使用go关键字为一个函数创建一个goroutine。一个函数可以创建多个goroutine，一个goroutine只能对应一个函数。
//格式：go 函数名( 参数列表 )

//使用go关键字创建goroutine时，被调用函数的返回值会被忽略。
//如果需要在goroutine中返回数据，则需要根据通道的特性，通过通道把数据从goroutine中作为返回值传出。

func running() {
	var times int
	//无限循环
	for {
		times++
		fmt.Println("tick", times)

		//延迟2s
		time.Sleep(2 * time.Second)
	}
}

func main() {
	//并发执行程序
	go running()

	//
	var input string
	for i := 0; i < 5; i++ {
		scan, err := fmt.Scanln(&input)
		//返回值：它返回成功扫描的项目数
		fmt.Println("-----", scan)
		fmt.Println("&&&&&", input)

		if err != nil {
			return
		}
	}

}

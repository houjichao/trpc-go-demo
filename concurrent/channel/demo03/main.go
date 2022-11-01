package main

import "fmt"

/*
	通道学习
	通道的发送和接收
	把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞。
	Go 程序运行时能智能地发现一些永远无法发送成功的语句并做出提示，代码如下：
	fatal error: all goroutines are asleep - deadlock!

	① 通道的收发操作在不同的两个 goroutine 间进行。
	由于通道的数据在没有接收方处理时，数据发送方会持续阻塞，因此通道的接收必定在另外一个 goroutine 中进行。
	② 接收将持续阻塞直到发送方发送数据。
	如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止。
	③ 每次接收一个元素。
	通道一次只能接收一个数据元素。

*/
func main() {
	// 创建一个空接口通道
	ch := make(chan interface{})
	// 将0放入通道中
	go func() {
		ch <- 0
	}()
	value1 := <-ch
	fmt.Println(value1)
	// 将hello字符串放入通道中
	go func() {
		ch <- "hello"
	}()
	var value2 = <-ch
	fmt.Println(value2)
}

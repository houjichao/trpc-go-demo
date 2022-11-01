package main

import "fmt"

/*
	单向通道
	var 通道实例 chan<- 元素类型    // 只能写入数据的通道
	var 通道实例 <-chan 元素类型    // 只能读取数据的通道
*/


func main() {
	ch := make(chan int)
	// 声明一个只能写入数据的通道类型, 并赋值为ch
	var chSendOnly chan<- int = ch
	fmt.Println(chSendOnly)
	//声明一个只能读取数据的通道类型, 并赋值为ch
	var chRecvOnly <-chan int = ch
	fmt.Println(chRecvOnly)

	//只能读取，不能写入
	//Invalid operation: chRecvOnly <- 0 (send to receive-only type <-chan int)
	//chRecvOnly <- 0

	//Invalid operation: <- chSendOnly (receive from send-only type chan<- int)
	//只能写入，不能读取
	//num := <-chSendOnly

	close(ch)

	x, ok := <-ch
	if !ok {
		println(x)
	}
}

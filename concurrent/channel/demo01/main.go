package main

import "fmt"

/*
	并发编程 -- 通道

题目1:协程之间如何通信？

通过管道之间进行通信
加锁
也可以共享内存

题目2:通道channel如何创建？

通道是引用类型，通过make进行创建
通道实例 := make(chan 数据类型)
//用来通知
ch := make(chan int)

chan类型发送数据和接收值
发送数据：ch <- data
接收值: data <- ch

带缓冲的通道和无缓冲通道的区别？
无缓冲的通道必须保证发送发和接收方的同步，否则会发生阻塞
带缓冲的通道是在无缓冲通道的基础上为通道增加了一个有限大小的存储空间形成缓冲通道。
发送时无需等待接受方接受即可完成发送过程。同理，只要缓冲通道有数据，接受时将不会发生阻塞。

带缓冲的通道len和cap代表的含义是什么？
len是通道中元素的数量
cap表示通道的容量，同时能够容纳多少数据

在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。goroutine 间通过通道就可以通信。
*/

/*
 将
*/
func SumNum(arr []int) int {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var mid = len(arr) / 2

	go func() {
		var sum = 0
		for i := 0; i < mid; i++ {
			sum += arr[i]
		}
		ch1 <- sum
	}()

	go func() {
		var sum = 0
		for i := mid; i < len(arr); i++ {
			sum += arr[i]
		}
		ch2 <- sum
	}()
	return <-ch1 + <-ch2
}

func main() {
	var arr [200]int

	for i := 0; i < 200; i++ {
		arr[i] = i + 1
	}
	slice := arr[:]
	fmt.Println(SumNum(slice))
}

package main

import "fmt"

/*
协程学习

题目1:并行与并发的区别？
并行是指多个程序同时运行
并发是指多个程序轮流使用时间片，看起来好像同时在运行。


题目2:go的并发如何执行？
使用go关键字开一个goroutine
*/

func main() {

	//用来通知
	ch := make(chan int)

	//开启一个协程
	go func() {
		for i := 1; i < 100; i++ {
			fmt.Println("a", i)
		}
		ch <- 0
	}()

	//main线程打印
	for i := 1; i < 100; i++ {
		fmt.Println("b", i)
	}

	//接受通知 未收到通知 会发生阻塞
	<-ch

}

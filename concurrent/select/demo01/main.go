package main

import (
	"fmt"
	"math/rand"
)

/*
	select的作用是什么？
	同时响应多个通道的操作。select的每个case都会对应一个通道的收发过程。当收发完成时，就会触发case中响应的语句。

	select中的case语句需要做怎样的操作？
	操作	示例
	接受任意数据	case <-ch
	接受变量	case d:=<-ch
	发送数据	case ch<-100

	for-select中的break、continue、return分别是什么作用？、
	break 跳出select 不会跳出for循环
	continue 结束本次执行 后续的语句不再执行
	return 跳出select和for

*/
/*
	实现两个 goroutine，其中一个产生500以内的随机数并写入到 channel 中，
	另外一个从 channel 中读取数字并打印到标准输出，最终输出五个随机数，尝试使用for-select实现
*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch := make(chan int, 5)

	//统计次数
	var count = 0
	go func() {
		for {
			select {
			case ch <- rand.Intn(500):
				count++
				if count == 5 {
					ch1 <- 0
					return
				}
			}
		}

	}()

	go func() {
		for {
			select {
			case data := <-ch:
				fmt.Println(data)
			case <-ch1:
				ch2 <- 0
			}
		}
	}()

	<-ch2
}

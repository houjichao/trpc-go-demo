package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，
当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。
反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作。
*/
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan string, 100)
	defer close(ch1)
	defer close(ch2)
	go func() {
		for i := 1; i < 10; i++ {
			ch1 <- i
			fmt.Println("ch1：", i)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			ch2 <- "a" + strconv.Itoa(i)
			fmt.Println("ch2：", i)
			time.Sleep(time.Second)
		}
	}()

	for v := range ch1 {
		fmt.Println(v)
		//这里的条件控制要严格，如果是>9则会panic
		if v == 9 {
			break
		}
	}

	for i := 0; i < len(ch2); i++ {
		var value = <-ch2
		fmt.Println(value)
		if i == 9 {
			break
		}
	}

}

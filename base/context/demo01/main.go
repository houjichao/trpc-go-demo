package main

import (
	"context"
	"fmt"
)

/**
Context 是Go 语言独有功能之一，用于上下文控制，可以在 goroutine 中进行传递。

context 与 select-case 联合，还可以实现上下文的截止时间、信号控制、信息传递等跨 goroutine 的操作，是 Go 语言协程的重要组成部分。
*/

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				//Done：获取一个只读的 channel，类型为结构体。可用于识别当前 channel 是否已经被关闭，其原因可能是到期，也可能是被取消了。
				case <-ctx.Done():
					close(dst)
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	//WithCancel：基于父级 context，创建一个可以取消的新 context。
	//Background：创建一个空的 context，一般常用于作为根的父级 context。

	//不传递 nil context
	//很多时候我们在创建 context 时，还不知道其具体的作用和下一步用途是什么。
	//这种时候大家可能会直接使用 context.Background 方法：
	//但在实际的 context 建议中，我们会建议使用 context.TODO() 方法来创建顶级的 context，直到弄清楚实际 Context 的下一步用途，再进行变更。
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithCancel(context.TODO())

	// defer cancel() // 实际使用中应该在这里调用 cancel
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel() // 这里为了使不熟悉 go 的更能明白在这里调用了 cancel()
			break
		}
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

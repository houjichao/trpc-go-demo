package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	如果两个或者多个 goroutine 在没有相互同步的情况下，访问某个共享的资源，比如同时对该资源进行读写时，就
	会处于相互竞争的状态，这就是并发中的资源竞争
	我们对于同一个资源的读写必须是原子化的，同一时间只能允许一个goroutine对资源进行读写操作。

	Go为我们提供了一个工具帮助我们检查，这个就是go build -race命令。在项目目录下执行这个命令，生成一个可以执行的文件，
	然后运行这个可执行文件，就可以看到打印出的检测信息。
	在 go build 命令中多加了一个 -race 标志，这样生成的可执行程序就自带了检测资源竞争的功能
*/
/*
	锁住共享资源
	原子函数加锁
	原子函数能够以底层的加锁机制来同步访问整型变量和指针
*/
/*
==================
WARNING: DATA RACE
Read at 0x00000122ec3c by goroutine 8:
  main.incCount()
      /Users/houjichao/Work/Go/hjc/trpc-go-demo/concurrent/goroutine/demo05/chain_test.go:39 +0x7c

Previous write at 0x00000122ec3c by goroutine 7:
  main.incCount()
      /Users/houjichao/Work/Go/hjc/trpc-go-demo/concurrent/goroutine/demo05/chain_test.go:42 +0x9b

Goroutine 8 (running) created at:
  main.main()
      /Users/houjichao/Work/Go/hjc/trpc-go-demo/concurrent/goroutine/demo05/chain_test.go:32 +0x7c

Goroutine 7 (finished) created at:
  main.main()
      /Users/houjichao/Work/Go/hjc/trpc-go-demo/concurrent/goroutine/demo05/chain_test.go:31 +0x64
==================
4

*/
var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}
func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}

}

func incCounter() {
	value := count
	count = value
	fmt.Println(count)
}

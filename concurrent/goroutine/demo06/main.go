package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	两个有用的原子函数是 LoadInt64 和 StoreInt64。
	这两个函数提供了一种安全地读和写一个整型值的方式。
	下面是代码就使用了 LoadInt64 和 StoreInt64 函数来创建一个同步标志，这个标志可以向程序里多个 goroutine 通知某个特殊状态。
*/
var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(10 * time.Second)
	fmt.Println("Shutdown Now")
	// main 函数使用 StoreInt64 函数来安全地修改 shutdown 变量的值。如果哪个 doWork goroutine
	//试图在 main 函数调用 StoreInt64 的同时调用 LoadInt64 函数，那么原子函数会将这些调用互相同步，
	//保证这些操作都是安全的，不会进入竞争状态。
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}
func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}

}

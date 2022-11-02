package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, done, wg)
	}
	close(done)
	wg.Wait()
	fmt.Println("all done")

	/*
	看起来好像 goroutine 都执行完了，然而报错：
	-- fatal error: all goroutines are asleep - deadlock!
	实际报错：panic: sync: negative WaitGroup counter
	为什么会发生死锁？goroutine 在退出前调用了 wg.Done() ，程序应该正常退出的。
	原因是 goroutine 得到的 "WaitGroup" 变量是 var wg WaitGroup 的一份拷贝值，即 doIt() 传参只传值。
	所以哪怕在每个 goroutine 中都调用了 wg.Done()， 主程序中的 wg 变量并不会受到影响。

	*/

}

func doIt(workerID int, done <-chan struct{}, wg sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	time.Sleep(3 * time.Second) // 模拟 goroutine 正在执行
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerID)
}

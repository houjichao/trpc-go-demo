package main

import (
	"fmt"
	"sync"
	"time"
)

var  (
	wg sync.WaitGroup
)

func main() {
	done := make(chan struct{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, done, &wg)// wg 传指针，doIt() 内部会改变 wg 的值
	}
	close(done)
	wg.Wait()
	fmt.Println("all done")

}

func doIt(workerID int, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	time.Sleep(3 * time.Second) // 模拟 goroutine 正在执行
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerID)
}

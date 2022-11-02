package main

import (
	"fmt"
	"time"
)
//程序默认不等所有 goroutine 都执行完才退出，这点需要特别注意：
//主程序会直接退出，这显然是不合理的
func main() {
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		go doIt(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("all done!")
}

func doIt(workerID int) {
	fmt.Printf("[%v] is running\n", workerID)
	time.Sleep(3 * time.Second)        // 模拟 goroutine 正在执行
	fmt.Printf("[%v] is done\n", workerID)
}
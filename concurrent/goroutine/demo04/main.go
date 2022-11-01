package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
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

var (
	counter int64
	wg      sync.WaitGroup //用于等待一组线程的结束
)

func main() {
	wg.Add(2) //增加两个线程
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println(counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 10; count++ {
		atomic.AddInt64(&counter, 1) //安全的对counter加一
		runtime.Gosched()            //让当前线程暂停
	}
}

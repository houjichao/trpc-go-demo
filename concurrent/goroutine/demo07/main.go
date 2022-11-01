package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	互斥锁
	同步访问共享资源的方式是使用互斥锁，互斥锁这个名字来自互斥的概念。
	互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界代码。

	Go语言等待组（sync.WaitGroup）
	Go语言中除了可以使用通道（channel）和互斥锁进行两个并发程序间的同步外，还可以使用等待组进行多个任务的同步，等待组可以保证在并发环境中完成指定数量的任务
	在 sync.WaitGroup（等待组）类型中，每个 sync.WaitGroup 值在内部维护着一个计数，此计数的初始默认值为零。

	等待组有下面几个方法可用，如下表所示。
	等待组的方法
	方法名	功能
	(wg * WaitGroup) Add(delta int)	等待组的计数器 +1
	(wg * WaitGroup) Done()	等待组的计数器 -1
	(wg * WaitGroup) Wait()	当等待组计数器不等于 0 时阻塞直到变 0。

	对于一个可寻址的 sync.WaitGroup 值 wg：
	我们可以使用方法调用 wg.Add(delta) 来改变值 wg 维护的计数。
	方法调用 wg.Done() 和 wg.Add(-1) 是完全等价的。
	如果一个 wg.Add(delta) 或者 wg.Done() 调用将 wg 维护的计数更改成一个负数，一个恐慌将产生。
	当一个协程调用了 wg.Wait() 时，如果此时 wg 维护的计数为零，则此 wg.Wait() 此操作为一个空操作（noop）；
	否则（计数为一个正整数），此协程将进入阻塞状态。当以后其它某个协程将此计数更改至 0 时（一般通过调用 wg.Done()），
	此协程将重新进入运行状态（即 wg.Wait() 将返回）。

	等待组内部拥有一个计数器，计数器的值可以通过方法调用实现计数器的增加和减少。
	当我们添加了 N 个并发任务进行工作时，就将等待组的计数器值增加 N。每个任务完成时，这个值减 1。
	同时，在另外一个 goroutine 中等待这个等待组的计数器值为 0 时，表示所有任务已经完成。
*/
var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println(counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 10; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()

	}

}

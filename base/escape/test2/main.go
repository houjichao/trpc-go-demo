package main

/**
2.变量在函数外部有引用
*/

func test() *int {
	a := 10
	return &a
}

func main() {
	_ = test()
}

/*
go build -gcflags '-m -m -l'
# github.com/houjichao/trpc-go-demo/base/escape/test2
./main.go:8:2: a escapes to heap:
./main.go:8:2:   flow: ~r0 = &a:
./main.go:8:2:     from &a (address-of) at ./main.go:9:9
./main.go:8:2:     from return &a (return) at ./main.go:9:2
./main.go:8:2: moved to heap: a
 */

/*
原因分析
变量a在函数外部存在引用。

我们来分析一下执行过程：当函数执行完毕，对应的栈帧就被销毁，但是引用已经被返回到函数之外。
如果这时外部通过引用地址取值，虽然地址还在，但是这块内存已经被释放回收了，这就是非法内存。

为了避免上述非法内存的情况，在这种情况下变量的内存分配必须分配到堆上。

func test() int {
	a := 10
	return a
}
如果将test()改为值类型，则不会有逃逸
 */
package main

/*
3.变量占用内存较大
*/

func test() {
	a := make([]int, 8192, 8192)
	for i := 0; i < 8192; i++ {
		a[i] = i
	}
}

func main() {
	test()
}

/**
# github.com/houjichao/trpc-go-demo/base/escape/test3
./main.go:8:11: make([]int, 8191, 8191) does not escape
 ~/Work/Go/hjc/trpc-go-demo/base/escape/test3   main ●✚  go build -gcflags '-m -m -l'
# github.com/houjichao/trpc-go-demo/base/escape/test3
./main.go:8:11: make([]int, 8192, 8192) escapes to heap:
./main.go:8:11:   flow: {heap} = &{storage for make([]int, 8192, 8192)}:
./main.go:8:11:     from make([]int, 8192, 8192) (too large for stack) at ./main.go:8:11
./main.go:8:11: make([]int, 8192, 8192) escapes to heap
*/

/*
原因分析
我们定义了一个容量为8192的int类型切片，发生了逃逸，内存分配到了堆上（heap）。

注意看：
我们再简单修改一下代码，将切片的容量和长度修改为8191，再次查看逃逸分析的结果，我们发现，没有发生逃逸，内存默认分类到了栈上。
所以，当变量占用内存较大时，会发生逃逸分析，将内存分配到堆上。
 */

/*
可以看demo05
 */

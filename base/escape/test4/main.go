package main

func test() {
	l := 1
	a := make([]int, l, l)
	for i := 0; i < l; i++ {
		a[i] = i
	}
}

func main() {
	test()
}

/*
go build -gcflags '-m -m -l'
# github.com/houjichao/trpc-go-demo/base/escape/test4
./main.go:5:11: make([]int, l, l) escapes to heap:
./main.go:5:11:   flow: {heap} = &{storage for make([]int, l, l)}:
./main.go:5:11:     from make([]int, l, l) (non-constant size) at ./main.go:5:11
./main.go:5:11: make([]int, l, l) escapes to heap
 */

/*
原因分析
我们通过控制台的输出结果可以很明显的看出：发生了逃逸，分配到了heap堆中。

原因是这样的：

我们虽然在代码段中给变量 l 赋值了1，但是编译期间只能识别到初始化int类型切片时，传入的长度和容量是变量l，编译期并不能确定变量l的值，所以发生了逃逸，会把内存分配到堆中
 */
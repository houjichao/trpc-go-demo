package main

import "fmt"

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

//在闭包函数中使用外部变量
func main() {
	in := Increase()
	fmt.Println(in()) // 1
	fmt.Println(in()) // 2

	/*
		go build -gcflags "-m -l" chain_test.go
		# command-line-arguments
		./chain_test.go:6:2: moved to heap: n
		./chain_test.go:7:9: func literal escapes to heap
		./chain_test.go:15:13: ... argument does not escape
		./chain_test.go:15:16: in() escapes to heap
		./chain_test.go:16:13: ... argument does not escape
		./chain_test.go:16:16: in() escapes to heap
	*/
	/*
		根据变量类型是否确定
		在上边例子中，也许你发现了，所有编译输出的最后一行中都是 m escapes to heap 。
		奇怪了，为什么 m 会逃逸到堆上？
		其实就是因为我们调用了 fmt.Println() 函数，它的定义如下
		func Println(a ...interface{}) (n int, err error) {
		    return Fprintln(os.Stdout, a...)
		}
		Go
		可见其接收的参数类型是 interface{} ，对于这种编译期不能确定其参数的具体类型，编译器会将其分配于堆上。
	*/
}

package main

import "fmt"

/**
1.参数是interface类型
*/
func main() {
	a := 666
	fmt.Println(a)
}

/*
go build -gcflags '-m -m -l'
# github.com/houjichao/trpc-go-demo/base/escape/test1
./main.go:10:13: a escapes to heap:
./main.go:10:13:   flow: {storage for ... argument} = &{storage for a}:
./main.go:10:13:     from a (spill) at ./main.go:10:13
./main.go:10:13:     from ... argument (slice-literal-element) at ./main.go:10:13
./main.go:10:13:   flow: {heap} = {storage for ... argument}:
./main.go:10:13:     from ... argument (spill) at ./main.go:10:13
./main.go:10:13:     from fmt.Println(... argument...) (call parameter) at ./main.go:10:13
./main.go:10:13: ... argument does not escape
./main.go:10:13: a escapes to heap
*/

/*
原因分析
因为Println(a ...interface{})的参数是interface{}类型，编译期无法确定其具体的参数类型，所以内存分配到堆中。

参考原则1：
1.不同于JAVA JVM的运行时逃逸分析，Go的逃逸分析是在编译期完成的：编译期无法确定的参数类型必定放到堆中；
*/

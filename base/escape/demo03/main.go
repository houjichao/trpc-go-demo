package main

import "fmt"

//返回任意引用型的变量：Slice 和 Map
func foo() []int {
	a := []int{1,2,3}
	return a
}

func main() {
	b := foo()
	fmt.Println(b)

	/*
	go build -gcflags "-m -l" main.go
	# command-line-arguments
	./main.go:7:12: []int{...} escapes to heap
	./main.go:13:13: ... argument does not escape
	./main.go:13:13: b escapes to heap
	 */
}
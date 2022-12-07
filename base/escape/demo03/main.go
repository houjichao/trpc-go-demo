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
	go build -gcflags "-m -l" chain_test.go
	# command-line-arguments
	./chain_test.go:7:12: []int{...} escapes to heap
	./chain_test.go:13:13: ... argument does not escape
	./chain_test.go:13:13: b escapes to heap
	 */
}
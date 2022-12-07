package main



func test2() int {
	c1 := 20
	return c1
}

func main() {
	c2 := test2()
	println(c2)
	//go build -gcflags '-m -l' chain_test.go
	/*
	# command-line-arguments
	./demo.go:12:13: ... argument does not escape
	./demo.go:12:13: m escapes to heap
	从分析的结果我们并没有看到任何关于 v 变量的逃逸说明，说明其并没有逃逸，它是分配在栈上的。
	而如果该变量还需要在函数范围之外使用，如果还在栈上分配，那么当函数返回的时候，
	该变量指向的内存空间就会被回收，程序势必会报错，因此对于这种变量只能在堆上分配。
	*/
}

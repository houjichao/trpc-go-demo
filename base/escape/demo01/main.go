package main

/*
	什么是逃逸分析
	在编译程序优化理论中，逃逸分析是一种确定指针动态范围的方法——分析在程序的哪些地方可以访问到指针。
	也是就是说逃逸分析是解决指针作用范围的编译优化方法。编程中常见的两种逃逸情景：


	函数中局部对象指针被返回（不确定被谁访问）
	对象指针被多个子程序（如线程 协程）共享使用
	为什么要做逃逸分析
	开始我们提到go语言中对象内存的分配不是由语言运算符或函数决定，而是通过逃逸分析来决定。为什么要这么干呢？
	其实说到底还是为了优化程序。函数中生成一个新对象：
	如果分配到栈上，待函数返回资源就被回收了
	如果分配到堆上，函数返回后交给gc来管理该对象资源
	栈资源的分配及回收速度比堆要快，所以逃逸分析最大的好处应该是减少了GC的压力。

	使用命令go build -gcflags “-m -l” main.go可以分析变量是否逃逸
	go build -gcflags "-m -m -l" chain_test.go 更详细
 */
//指针逃逸
//典型的指针逃逸案例，返回局部变量的指针
func test2() *int {
	c1 := 20
	return &c1
}

func main() {
	c2 := test2()
	println(*c2)

	/*
	go build -gcflags "-m -m -l" chain_test.go
	# command-line-arguments
	./chain_test.go:24:2: c1 escapes to heap:
	./chain_test.go:24:2:   flow: ~r0 = &c1:
	./chain_test.go:24:2:     from &c1 (address-of) at ./chain_test.go:25:9
	./chain_test.go:24:2:     from return &c1 (return) at ./chain_test.go:25:2
	./chain_test.go:24:2: moved to heap: c1
	 */
}

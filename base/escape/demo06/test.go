package demo06

//由于逃逸分析是在编译期就运行的，而不是在运行时运行的。因此避免有一些不定长的变量可能会很大，
//而在栈上分配内存失败，Go 会选择把这些变量统一在堆上申请内存，这是一种可以理解的保险的做法。

var size = 10

func foo() {
	length := 10
	arr := make([]int, 0, length)
}

func bar() {
	arr := make([]int, 0, 10)
}

func test() {
	arr := make([]int, 0, size)
}

/*
go build -gcflags "-m -l" test.go
# command-line-arguments
./test.go:8:13: make([]int, 0, length) escapes to heap
./test.go:12:13: make([]int, 0, 10) does not escape
./test.go:16:13: make([]int, 0, size) escapes to heap
./test.go:8:2: arr declared but not used
./test.go:12:2: arr declared but not used
./test.go:16:2: arr declared but not used
*/

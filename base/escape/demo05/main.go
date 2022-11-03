package main

/*
根据变量的占用大小
最开始的时候，就介绍到，以 64KB 为分界线，我们将内存块分为 小内存块 和 大内存块。

小内存块走常规的 mspan 供应链申请，而大内存块则需要直接向 mheap，在堆区申请。
*/
func foo() {
	nums1 := make([]int, 8191) // < 64KB
	for i := 0; i < 8191; i++ {
		nums1[i] = i
	}
}

func bar() {
	nums2 := make([]int, 8192) // = 64KB
	for i := 0; i < 8192; i++ {
		nums2[i] = i
	}
}

/*
go build -gcflags "-m -l" main.go
# command-line-arguments
./main.go:4:15: make([]int, 8191) does not escape
./main.go:11:15: make([]int, 8192) escapes to heap
*/

/*
那为什么是 64 KB 呢？

我只能说是试出来的 （8191刚好不逃逸，8192刚好逃逸），
网上有很多文章千篇一律的说和 ulimit -a 中的 stack size 有关，但经过了解这个值表示的是系统栈的最大限制是 8192 KB，刚好是 8M。

$ ulimit -a
-t: cpu time (seconds)              unlimited
-f: file size (blocks)              unlimited
-d: data seg size (kbytes)          unlimited
-s: stack size (kbytes)             8192
Bash
我个人实在无法理解这个 8192 （8M） 和 64 KB 是如何对应上的，如果有朋友知道，还请指教一下。
*/

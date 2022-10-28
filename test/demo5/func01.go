package main

import "fmt"

/*
	go支持可变参数的函数
*/

func test(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func main() {
	test()
	fmt.Println("-----------")
	test(3)
	fmt.Println("-----------")
	test(10, 11, 12)
}

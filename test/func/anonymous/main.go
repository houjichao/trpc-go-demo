package main

import "fmt"
//如何让一个匿名函数，可以在整个程序中有效？将匿名函数给一个全局变量就可以了
var Mult = func(num1 int, num2 int) int {
	return num1 * num2
}

/*
	匿名函数学习
*/
func main() {
	//1.定义的同时进行调用，是匿名函数使用最多的场景
	sum := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println("sum is:", sum)

	//将匿名函数赋给一个变量
	//sub变量等价于匿名函数
	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}
	result01 := sub(10, 3)
	fmt.Println("result01 is:", result01)

	result02 := Mult(3, 4)
	fmt.Println("result02 is:", result02)

}

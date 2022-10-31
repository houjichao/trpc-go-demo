package main

import "fmt"

/*
	闭包学习
	闭包就是一个函数和与其相关的引用环境组合的一个整体

	闭包本质：
	是一个匿名函数，只是这个函数引入外借的变量/参数
	匿名函数 + 引用的变量/参数 = 闭包
*/

/*
	函数功能：求和
	函数的名字：getSum 参数为空
	getSum函数返回值为一个函数，这个函数的参数是一个int类型的参数，返回值也是int类型
*/
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}
//闭包：返回的匿名函数+匿名函数以外的变量sum

//闭包应用场景：闭包可以保留上次引用的某个值，我们传入一次就可以反复使用了


func main() {
	f := getSum()
	fmt.Println(f(1))//1
	fmt.Println(f(2))//3
	fmt.Println(f(3))//6
	fmt.Println(f(4))//10
	//匿名函数中引用的变量会一直保存在内存中，可以一直使用
}

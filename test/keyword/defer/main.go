package main

import "fmt"

/*
	defer关键字学习：
	用途：在函数中，程序猿经常需要创建资源，为了在函数执行完毕后，及时的释放资源，go的设计者提供defer关键字

	应用场景：
	比如你想关闭某个使用的资源，在使用的时候直接随手defer，因为defer有延迟执行机制（函数执行完毕再执行defer压入栈中的语句）
	所以用完随手写了关闭，比较省心，省事

*/
func main() {
	fmt.Println("add", add(30, 60))
}

func add(num1 int, num2 int) int {
	//在go中，程序遇到defer关键字，不会立即执行defer后的语句，而是将defer后的语句压入一个栈中，然后继续执行函数后面的语句
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)

	//发现，遇到defer关键字，会将后面的代码语句压入栈中，也会将相同的值拷贝入栈中，不会随着函数后面的变化而变化
	num1 += 90 //num1=120
	num2 += 50 //num2=110
	//栈的特点是：先进后出
	//在函数执行完毕以后，从栈中取出语句开始执行，按照先进后出的规则执行
	var sum = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}

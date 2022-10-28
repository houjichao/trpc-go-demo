package main

import "fmt"

/*
	函数学习
	函数使用驼峰命名
	首字母大写该函数可以被本包文件和其他包文件使用--类似public
	首字母小写只能被本包文件使用，其他包文件不能使用--类似private
*/

func exchangeNumFalse(num1 int, num2 int) {
	//两种方式
	num1, num2 = num2, num1

	/*var t int
	t = num1
	num1 = num2
	num2 = t*/
}

/*
	以值传递方式的数据类型，如果希望在函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量。
	从效果来看类似引用传递
*/
func exchangeNumTrue(ptr1 *int, ptr2 *int) {
	*ptr1, *ptr2 = *ptr2, *ptr1
}

func main() {
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("交换前的两个数：num1 = %v,num2 = %v \n", num1, num2)
	//内存分析
	//基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值
	exchangeNumFalse(num1, num2)
	fmt.Printf("exchangeNumFalse交换后的两个数：num1 = %v,num2 = %v \n", num1, num2)

	//
	exchangeNumTrue(&num1, &num2)
	fmt.Printf("exchangeNumTrue交换后的两个数：num1 = %v,num2 = %v \n", num1, num2)

}

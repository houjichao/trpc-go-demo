package main

import "fmt"

/*
	学习数组定义的几种方式
*/
func main() {
	//第一种
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)
	//长度属于类型的一部分
	fmt.Printf("数组的类型为:%T \n", arr1)
	//第二种
	var arr2 = [3]int{11, 22, 33}
	fmt.Println(arr2)
	//第三种
	var arr3 = [...]int{4, 5, 6, 7}
	fmt.Println(arr3)
	//第四种
	var arr4 = [...]int{2: 11, 0: 9, 1: 10, 3: 12}
	fmt.Println(arr4)
}

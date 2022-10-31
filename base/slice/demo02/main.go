package main

import "fmt"

/*
	定义切片的三种方式
*/
func main() {

	//方式1：定义一个切片，然后让切片去引用一个已经创建好的数组
	var intarr = [6]int{3, 6, 9, 1, 4, 7}
	var slice1 = intarr[1:3]
	fmt.Println(slice1)

	//方式2：通过make内置函数来创建切片，基本语法：var 切片名 = make([]type,len,cap)
	var slice2 = make([]int, 4, 20)
	fmt.Println(slice2)
	slice2[0] = 66
	slice2[1] = 88
	//slice2[19] = 100
	//panic: runtime error: index out of range [19] with length 4
	fmt.Println(slice2)
	fmt.Printf("slice2的数据类型是%T\n", slice2)
	//PS:make底层创建一个数组，对外不可见，所以不可以直接操作这个数组，要通过slice去间接的访问各个元素，不可以直接对数组进行维护/操作

	//方式3：定一个切片，直接就指定具体数组，使用原理类似make的方式
	slice3 := []int{1, 4, 7}
	fmt.Println(slice3)
	fmt.Printf("slice3的数据类型是%T\n", slice3)
}

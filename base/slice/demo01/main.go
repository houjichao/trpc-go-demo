package main

import "fmt"

/*
	切片学习
	切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型，
	这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内。

	Go语言中切片的内部结构包含地址、大小和容量，切片一般用于快速地操作一块数据集合，
	如果将数据集合比作切糕的话，切片就是你要的“那一块”，切的过程包含从哪里开始（切片的起始位置）及切多大（切片的大小），
	容量可以理解为装切片的口袋大小
*/
func main() {

	//切片的学习
	//定义数组
	var intarr = [6]int{3, 6, 9, 1, 4, 7}
	fmt.Println(intarr)

	//切片构建在数组之上
	//[1:3]切片 - 切出一段片段 - 索引：从1开始，到3结束（不包含3）

	var slice1 = intarr[1:3]
	fmt.Println(slice1)

	//元素个数
	fmt.Println("slice的元素个数：", len(slice1))
	//获取切片容量，容量可以动态变化
	fmt.Println("slice的容量：", cap(slice1))

	//切片内存分析
	//切片有3个字段的数据结构：一个是指向底层数组的指针，一个是切片的长度，一个是切片的容量
	fmt.Printf("数组中下标为1位置的地址：%p \n", &intarr[1]) //0xc0000160f8
	fmt.Printf("切片的地址：%p \n", &slice1)           //0xc00000c030
	fmt.Printf("切片中下标为0位置的地址：%p \n", &slice1[0]) //0xc0000160f8

	//通过切片改值
	slice1[1] = 16
	fmt.Println("intarr:", intarr)
	fmt.Println("slice1:", slice1)

	fmt.Println("---------------")
	// 声明字符串切片
	var strList []string

	// 声明整型切片
	var numList []int

	// 声明一个空切片  分配了内存 但是没有元素,所以不为空nil
	var numListEmpty = []int{}

	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)

	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))

	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
}

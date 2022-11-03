package main

import (
	"container/list"
	"fmt"
)

/*
	list学习
*/
func main() {

	//初始化 通过 container/list 包的 New() 函数初始化 list
	//得到的是一个指针类型的数据
	list1 := list.New()
	fmt.Println(list1)

	//var声明
	//通过 var 关键字声明初始化 listvar
	//得到的类型是list.List的
	//列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型
	//，这既带来了便利，也引来一些问题，例如给列表中放入了一个 interface{} 类型的值，
	//取出值后，如果要将 interface{} 转换为其他类型将会发生宕机
	var list2 list.List
	fmt.Println(list2)

	list2.PushBack("111")
	list2.PushFront(65)

	// 尾部添加后保存元素句柄
	element := list2.PushBack("fist")

	// 在fist之后添加high
	list2.InsertAfter("high", element)

	// 在fist之前添加noon
	list2.InsertBefore("noon", element)

	//1.获取列表头节点
	fmt.Println(list2.Front().Value)
	//2.获取列表尾节点
	fmt.Println(list2.Back().Value)
	//3.获取上一个节点
	fmt.Println(element.Prev().Value)
	//4.获取下一个节点
	fmt.Println(element.Next().Value)

	fmt.Println("--------------------------")
	//遍历
	for i := list2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}

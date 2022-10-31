package main

import "fmt"

/*
	数组内存分析
*/
func main() {

	//声明数组
	var arr [3]int16
	//获取数组的长度，声明这里已经开辟了内存空间
	fmt.Println(len(arr))
	//打印数组，这里是数据类型的初始化值
	fmt.Println(arr)
	//证明arr中存储的是地址值：0 -> 1 -> 2 地址每次变化2，和数据类型有关，int16占两个字节
	fmt.Printf("arr的地址为：%p\n", &arr)       //0xc000114000
	fmt.Printf("arr[0]的地址为：%p\n", &arr[0]) //0xc000114000
	fmt.Printf("arr[1]的地址为：%p\n", &arr[1]) //0xc000114002
	fmt.Printf("arr[2]的地址为：%p\n", &arr[2]) //0xc000114004

}

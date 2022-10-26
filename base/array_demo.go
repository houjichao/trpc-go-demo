package base

import (
	"log"
)

/*
	数组学习

	数组定义
	var variable_name [size] variable_type
 */
func ArrayDemo()  {
	//var balance [10] float32

	//如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度
	//var balance1 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	var byteArr = [...]byte{'l', 'i', 'c', 'h', 'u', 'a', 'c', 'h', 'u', 'a'}

	//println(balance)
	arrStr := string(byteArr[:])
	log.Print(arrStr)

}

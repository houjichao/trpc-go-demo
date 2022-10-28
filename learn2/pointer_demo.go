package learn2

import "fmt"

/*
	指针学习

	指针就是内存地址
	&：取内存地址
	*：根据地址取值
*/

func PointerDemo() {

	var age int = 18
	//取地址
	fmt.Println(&age)

	//定义一个指针变量
	//可以理解为指向int型的指针
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr本身这个存储空间的地址为：", &ptr)

	//根据地址取值
	fmt.Println("ptr指向的数值为：", *ptr)

}

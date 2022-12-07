package pointer

import "fmt"

/*
	指针的四个细节
*/
func PointerDemo2() {

	//1.可以通过指针改变指向值
	var num int = 10
	fmt.Println("num的值为:", num)
	var ptr *int = &num
	*ptr = 20
	fmt.Println("num的值为：", num)

	//2.指针变量接收的一定是地址值

	//3.指针变量的地址不可以不匹配
	//var ptr *float32 = &num

	//4.基本数据类型（又叫值类型），都有对应的指针类型，形式为*数据类型
}

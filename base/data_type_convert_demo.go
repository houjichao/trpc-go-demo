package base

import "fmt"

/*
	数据类型转换：
	go里面只有显式转换，并且只有显式转换（强制转换）

	var v2 T2 = T2(n1)
 */
func DataTypeConvertDemo()  {

	var n1 int = 100
	fmt.Println(n1)

	var n2 float32 = float32(n1)
	fmt.Println(n2)
	fmt.Printf("n2的类型是:%T\n",n2)



}

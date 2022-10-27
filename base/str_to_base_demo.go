package base

import (
	"fmt"
	"strconv"
)

/*
	string转基础类型
*/
func StrToBaseDemo() {

	var s1 string = "18"
	var num1 int64
	//strconv.ParseInt的返回值有两个，第二个是错误，可以用_忽略
	num1, _ = strconv.ParseInt(s1, 10, 64)
	fmt.Printf("num1的类型是：%T，num1 = %v \n", num1, num1)
}

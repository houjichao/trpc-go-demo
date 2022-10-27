package base

import (
	"fmt"
	"strconv"
)

/*
	基础数据类型转换为string
*/
func BaseToStr() {
	//strconv函数
	var n1 int = 18
	var s1 string = strconv.FormatInt(int64(n1), 10)
	fmt.Printf("s1对应的类型是: %T, s1 = %q \n", s1, s1)

	var n2 float64 = 20.20
	//
	var s2 string = strconv.FormatFloat(n2, 'f', 8, 64)
	fmt.Printf("s2对应的类型是: %T, s2 = %q \n", s2, s2)

	var f1 bool = true
	var sf1 = fmt.Sprintf("%t", f1)
	fmt.Printf("sf1对应的类型是: %T, sf1 = %v\n", sf1, sf1)

}

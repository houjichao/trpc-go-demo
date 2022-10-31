package main

import (
	"fmt"
	"time"
)

/*
	日期和时间相关函数
*/
func main() {

	now := time.Now()
	fmt.Printf("%v ---对应的类型为: %T \n", now, now)

	fmt.Printf("年：%v\n", now.Year())

	//日期格式化
	//go中必须使用这个指定的字符串，不能用其他
	dateStr := now.Format("2006-01-02 15:04:05")
	fmt.Println(dateStr)

	//任意组合
	dateStr1 := now.Format("2006 15:04")
	fmt.Println(dateStr1)

}

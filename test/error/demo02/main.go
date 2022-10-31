package main

import (
	"errors"
	"fmt"
)

/*
	自定义错误学习
*/
func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误：", err)
		//该函数是为了终止程序执行
		panic(err)
	}
	fmt.Println("上面的除法操作执行成功---")
	fmt.Println("正常执行下面的逻辑---")
}

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		//抛出自定义错误
		return errors.New("除数不能为0！！！")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}

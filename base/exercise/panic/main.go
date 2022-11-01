package main

import (
	"fmt"
	"runtime"
)

/*
实现两个函数foo和bar， 其中foo调用bar， bar内部触发panic， foo中用recover来捕获panic
*/

func bar() {
	//panic(errors.New("is error"))
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func foo() {

	//需要注意的是defer语句应该 定义在bar之前
	defer func() {
		err := recover()

		switch err.(type) {
		case runtime.Error:
			fmt.Println("运行时错误", err)
		default:
			fmt.Println("error----", err)
		}
	}()

	bar()
}

func main() {
	foo()
	fmt.Println("异常被捕获，继续执行")
}

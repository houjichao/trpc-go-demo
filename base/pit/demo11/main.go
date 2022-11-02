package main

import "fmt"

// 错误的 recover 调用示例
/*func main() {
	recover()    // 什么都不会捕捉
	panic("not good")    // 发生 panic，主程序退出
	recover()    // 不会被执行
	println("ok")
}*/

// 正确的 recover 调用示例
func main() {
	defer func() {
		fmt.Println("recovered: ", recover())
	}()
	panic("not good")
	fmt.Println("111111")
}

//错误的调用示例
/*func main() {
	defer func() {
		defer doRecover()
	}()
	panic("not good")
}

func doRecover() {
	fmt.Println("recobered: ", recover())
}*/

package numutils

import "fmt"

var Num int

//init函数中给全局变量赋值
func init()  {
	fmt.Println("numutils中的init函数被执行了")
	Num = 10
}

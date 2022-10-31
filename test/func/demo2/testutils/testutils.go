package testutils

import "fmt"

var Age int
var Name string

//init函数中给全局变量赋值
func init()  {
	fmt.Println("testutils中的init函数被执行了")
	Age = 10
	Name = "jack"
}

package base

import "fmt"

/**
变量学习：
*/

//全局变量
var g int = 10

func VarDemo() {

	/* 声明局部变量 */
	var a, b, c int
	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b
	fmt.Printf("结果： a = %d, b = %d and c = %d\n", a, b, c)

	fmt.Printf("var demo 全局变量g = %d\n", g)

	var g int = 20

	fmt.Printf("var demo 全局变量g = %d\n", g)

}

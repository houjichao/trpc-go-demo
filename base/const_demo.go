package base

import "fmt"

/*
	const学习：常量
	常量是一个简单值的标识符，在程序运行时，不会被修改的量

	常量中的数据类型只可以是布尔型、数字型（整数、浮点和复数）和字符串型
	常量的格式定义：
	const const_name[type] = type
*/

func ConstDemo() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const x, y, z = 1, false, "str" //多重赋值
	area = LENGTH * WIDTH
	fmt.Println("面积为: %d", area)
	println()
	println(x, y, z)

	//iota
	//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
	//const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println("const iota demo: ",a, b, c, d, e, f, g, h, i)

	const (
		j=1<<iota //1
		k=3<<iota //6
		l         //3*(2^2)=12
		m         //3*(2^3)=24
	)

	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
	fmt.Println("m=",m)

	//左移算法：<<n == *(2^n)
}

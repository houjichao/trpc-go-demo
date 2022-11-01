package main

import "fmt"

/*
	int float32等基本数据类型都可以有方法,但是需要通过type重新定义，也就是起别名
 */
type myInt int

func (i myInt) print() {
	i = 30
	fmt.Println("print i = ",i)
}

func (i *myInt) change() {
	*i = 30
	fmt.Println("change i = ",*i)
}

func main()  {
	var i myInt = 20
	i.print()
	fmt.Println("main i = ",i)

	i.change()
	fmt.Println("main i = ",i)
}

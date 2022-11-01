package main

import "fmt"

/*
	方法学习
	1.方法是作用在指定的数据类型上，和指定的数据类型绑定，因此自定义类型，都可以有方法，而不仅仅是struct
	2.方法的声明和调用语法
	type Person struct {
		Name string
	}
	func (p Person) test() {
	}

*/
type Person struct {
	Name string
}

func (p Person) test() {
	p.Name = "露露"
	fmt.Println(p)
}

func (p *Person) test1() {
	(*p).Name = "露露"
	fmt.Println(*p)
}

func main() {

	var p Person
	p.Name = "李莉"
	p.test()
	//调用了test方法后，main中的name并没有被修改，说明是值传递
	fmt.Println(p)

	(&p).test1()
	//调用了test方法后，main中的name并没有被修改，说明是值传递
	fmt.Println(p)

}

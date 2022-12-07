package main

import "fmt"

/**
在方法中使用指针
*/
func main() {
	r := receiver{Name: "zs"}
	fmt.Println(r)
	//使用receiver作为方法参数
	r.methodA()
	fmt.Println(r)
	//使用*receiver作为方法参数
	r.methodB()
	fmt.Println(r)
}

type receiver struct {
	Id   int
	Name string
	Age  int
}

func (receiver receiver) methodA() {
	receiver.Name = "ls"
}

func (receiver *receiver) methodB() {
	receiver.Age = 19
}

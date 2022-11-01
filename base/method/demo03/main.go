package main

import "fmt"

/*
	int float32等基本数据类型都可以有方法,但是需要通过type重新定义，也就是起别名
*/
type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	str := fmt.Sprintf("学生Name =%v. Age=%v \n", s.Name, s.Age)
	return str
}

func main() {
	var stu = new(Student) //这种声明方式是一个指针
	stu.Name = "李莉"
	stu.Age = 20

	//stu已经是一个地址了
	fmt.Println(stu)

	var stu1 Student //这种声明方式是一个指针
	stu1.Name = "李莉1"
	stu1.Age = 21
	//传入地址，如果绑定了String方法就会自动调用
	fmt.Println(&stu1)

}

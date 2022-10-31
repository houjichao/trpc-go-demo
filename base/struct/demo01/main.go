package main

import "fmt"

/*
	结构体学习
*/

func main() {

	//创建老师结构体的对象实例
	//方式1：
	var t1 Teacher
	fmt.Println(t1)

	t1.Name = "jack"
	t1.Age = 30
	t1.School = "清华"
	fmt.Println(t1)

	//方式2
	var t2 Teacher = Teacher{"t2", 10, "家里蹲1"}
	fmt.Println(t2)

	//方式3 new函数
	//t3是指针，指向的就是地址，应该给这个地址的只想的对象的字段赋值
	var t3 *Teacher = new(Teacher)
	(*t3).Name = "t3"
	(*t3).Age = 13 //*的作用，根据地址取值
	t3.School = "家里蹲2"
	fmt.Println(*t3)

	//方式4：返回的结构体指针
	var t4 *Teacher = &Teacher{"t4", 45, "家里蹲3"}
	fmt.Println(*t4)
}

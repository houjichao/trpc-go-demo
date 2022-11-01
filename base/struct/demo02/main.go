package main

import "fmt"

/*
	结构体之间的转换
	1.结构体是用户单独定义的类型，和其他类型进行转换时需要就有完全相同的字段（名字、个数和类型）
	2.结构体进行type重新定义（相当于取别名），go认为是新的数据类型，但是相互间可以强转
*/

type Student struct {
	Age int
}

type Stu Student

type Person struct {
	Age int
}

func main() {
	//1
	var s = Student{18}
	var p = Person{30}
	s = Student(p)

	fmt.Println(s)
	fmt.Println(p)

	//2
	var s1 = Student{20}
	var s2 = Stu{25}
	s2 = Stu(s1)
	fmt.Println(s1)
	fmt.Println(s2)
}
